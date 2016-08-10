package monit

import (
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"

	"github.com/fatih/structs"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type MonitConfigs struct {
	Urls     []string
	Instance Monit
}

var sampleConfig = `
  ## An array of Apache status URI to gather stats.
  ## Default is "http://admin:monit@localhost:2812/_status?format=xml".
  urls = ["http://admin:monit@localhost:2812/_status?format=xml"]
`

func (n *MonitConfigs) SampleConfig() string {
	return sampleConfig
}

func (n *MonitConfigs) Description() string {
	return "Read monit stats from xml web page"
}

func (n *MonitConfigs) Gather(acc telegraf.Accumulator) error {
	if len(n.Urls) == 0 {
		n.Urls = []string{"http://admin:monit@localhost:2812/_status?format=xml"}
	}

	var outerr error
	var errch = make(chan error)

	for _, u := range n.Urls {
		addr, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("Unable to parse address '%s': %s", u, err)
		}

		go func(addr *url.URL) {
			errch <- n.gatherUrl(addr, acc)
		}(addr)
	}

	// Drain channel, waiting for all requests to finish and save last error.
	for range n.Urls {
		if err := <-errch; err != nil {
			outerr = err
		}
	}

	return outerr
}

var tr = &http.Transport{
	ResponseHeaderTimeout: time.Duration(3 * time.Second),
}

var client = &http.Client{
	Transport: tr,
	Timeout:   time.Duration(4 * time.Second),
}

func (m *MonitConfigs) gatherUrl(addr *url.URL, acc telegraf.Accumulator) error {
	resp, err := client.Get(addr.String())
	if err != nil {
		return fmt.Errorf("error making HTTP request to %s: %s", addr.String(), err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned HTTP status %s", addr.String(), resp.Status)
	}
	// tags := getTags(addr)
	m.Instance.Parse(resp.Body)
	m.Instance.structurizeFieldsForTelegraf(acc)

	return nil
}

func (s Service) createFieldsForTelegraf() map[string]interface{} {
	return structs.Map(s)
}

func (s ServiceSystem) createFieldsForTelegraf() map[string]interface{} {
	fields := s.Service.createFieldsForTelegraf()
	fields["System.Load.Average01"] = s.System.Load.Average01
	fields["System.Load.Average05"] = s.System.Load.Average05
	fields["System.Load.Average15"] = s.System.Load.Average15
	fields["System.CPU.System"] = s.System.CPU.System
	fields["System.CPU.User"] = s.System.CPU.User
	fields["System.CPU.Wait"] = s.System.CPU.Wait
	fields["System.Memory.KiloByte"] = s.System.Memory.KiloByte
	fields["System.Memory.Percent"] = s.System.Memory.Percent
	fields["System.Swap.KiloByte"] = s.System.Swap.KiloByte
	fields["System.Swap.Percent"] = s.System.Swap.Percent
	return fields
}

func (m *Monit) structurizeFieldsForTelegraf(acc telegraf.Accumulator) {
	acc.AddFields("monit-server", structs.Map(m.Server), nil)
	acc.AddFields("monit-platform", structs.Map(m.Platform), nil)
	for _, filesystem := range m.Filesystems {
		acc.AddFields("monit-filesystem", structs.Map(filesystem), nil)
	}
	for _, directory := range m.Directories {
		acc.AddFields("monit-directory", structs.Map(directory), nil)
	}
	for _, process := range m.Processes {
		acc.AddFields("monit-process", structs.Map(process), nil)
	}
	for _, host := range m.Hosts {
		acc.AddFields("monit-host", structs.Map(host), nil)
	}
	for _, system := range m.Systems {
		acc.AddFields("monit-system", system.createFieldsForTelegraf(), nil)
	}
	for _, fifo := range m.Fifos {
		acc.AddFields("monit-fifo", structs.Map(fifo), nil)
	}
	for _, programm := range m.Programms {
		acc.AddFields("monit-programm", structs.Map(programm), nil)
	}
	for _, network := range m.Networks {
		acc.AddFields("monit-network", structs.Map(network), nil)
	}
}

// Get tag(s) for the monit plugin
func getTags(addr *url.URL) map[string]string {
	h := addr.Host
	host, port, err := net.SplitHostPort(h)
	if err != nil {
		host = addr.Host
		if addr.Scheme == "http" {
			port = "80"
		} else if addr.Scheme == "https" {
			port = "443"
		} else {
			port = ""
		}
	}
	return map[string]string{"server": host, "port": port}
}

func (m *Monit) Parse(data io.ReadCloser) {

	defer data.Close()

	decoder := xml.NewDecoder(data)
	decoder.CharsetReader = charset.NewReaderLabel

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			switch {
			case se.Name.Local == "server":
				decoder.DecodeElement(&m.Server, &se)
				fmt.Printf("%+v\n", m.Server)
			case se.Name.Local == "platform":
				decoder.DecodeElement(&m.Platform, &se)
				fmt.Printf("%+v\n", m.Platform)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "0":
				decoder.DecodeElement(&m.Filesystems, &se)
				fmt.Printf("%+v\n", m.Filesystems)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "1":
				decoder.DecodeElement(&m.Directories, &se)
				fmt.Printf("%+v\n", m.Directories)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "2":
				decoder.DecodeElement(&m.Files, &se)
				fmt.Printf("%+v\n", m.Files)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "3":
				decoder.DecodeElement(&m.Processes, &se)
				fmt.Printf("%+v\n", m.Processes)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "4":
				decoder.DecodeElement(&m.Hosts, &se)
				fmt.Printf("%+v\n", m.Hosts)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "5":
				decoder.DecodeElement(&m.Systems, &se)
				fmt.Printf("%+v\n", m.Systems)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "6":
				decoder.DecodeElement(&m.Fifos, &se)
				fmt.Printf("%+v\n", m.Fifos)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "7":
				decoder.DecodeElement(&m.Programms, &se)
				fmt.Printf("%+v\n", m.Programms)
			case se.Name.Local == "service" && se.Attr[0].Name.Local == "type" && se.Attr[0].Value == "8":
				decoder.DecodeElement(&m.Networks, &se)
				fmt.Printf("%+v\n", m.Networks)
			}
		}
	}
}

func init() {
	inputs.Add("monit", func() telegraf.Input {
		return &MonitConfigs{}
	})
}
