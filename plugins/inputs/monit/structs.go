package monit

type Manager struct {
	Url      string
	User     string
	Password string
}

type Monit struct {
	Manager
	Server      `xml:"server"`
	Platform    `xml:"platform"`
	Filesystems []ServiceFilesystem // type=0
	Directories []ServiceDirectory  // type=1
	Files       []ServiceFile       // type=2
	Processes   []ServiceProcess    // type=3
	Hosts       []ServiceHost       // type=4
	Systems     []ServiceSystem     // type=5
	Fifos       []ServiceFifo       // type=6
	Programms   []ServiceProgramm   // type=7
	Networks    []ServiceNet        // type=8
}

type Server struct {
	ID            string `xml:"id"`
	Incarnation   string `xml:"incarnation"`
	Version       string `xml:"version"`
	Uptime        string `xml:"uptime"`
	Poll          string `xml:"poll"`
	StartDelay    string `xml:"startdelay"`
	LocalHostName string `xml:"localhostname"`
	ControlFile   string `xml:"controlfile"`
	Httpd         struct {
		Address string `xml:"address"`
		Port    string `xml:"port"`
		SSL     int    `xml:"ssl"`
	} `xml:"httpd" structs:",dotflatten"`
}

type Platform struct {
	Name    string `xml:"name"`
	Release string `xml:"release"`
	Version string `xml:"version"`
	Machine string `xml:"machine"`
	CPU     int    `xml:"cpu"`
	Memory  int    `xml:"memory"`
	Swap    int    `xml:"swap"`
}

// ServiceFile - monit's <service type="2">
type ServiceFile struct {
	Service        `structs:",flatten"`
	FileAttributes `structs:",flatten"`
}

// ServiceDirectory - monit's <service type="1">
type ServiceDirectory struct {
	Service        `structs:",flatten"`
	FileAttributes `structs:",flatten"`
}

type FileAttributes struct {
	Mode      int    `xml:"mode"`
	UID       int    `xml:"uid"`
	GID       int    `xml:"gid"`
	Timestamp uint64 `xml:"timestamp"`
	Size      uint64 `xml:"size"`
}

// ServiceFifo - monit's <service type="6">
type ServiceFifo struct {
	Service        `structs:",flatten"`
	FileAttributes `structs:",flatten"`
}

// ServiceFilesystem - monit's <service type="0">
type ServiceFilesystem struct {
	Service `structs:",flatten"`
	Mode    int `xml:"mode"`
	UID     int `xml:"uid"`
	GID     int `xml:"gid"`
	Flags   int `xml:"flags"`
	Block   struct {
		Percent float64 `xml:"percent"`
		Usage   float64 `xml:"usage"`
		Total   float64 `xml:"total"`
	} `xml:"block" structs:",dotflatten"`
	Inode struct {
		Percent float64 `xml:"percent"`
		Usage   float64 `xml:"usage"`
		Total   float64 `xml:"total"`
	} `xml:"inode" structs:",dotflatten"`
}

// ServiceNet - monit's <service type="8">
type ServiceNet struct {
	Service `structs:",flatten"`
	Link    struct {
		State    int     `xml:"state"`
		Speed    int     `xml:"speed"`
		Duplex   int     `xml:"duplex"`
		Download Network `xml:"download" structs:",dotflatten"`
		Upload   Network `xml:"upload" structs:",dotflatten"`
	} `xml:"link" structs:",flatten"`
}

type Network struct {
	Packets NetworkStats `xml:"packets" structs:",dotflatten"`
	Bytes   NetworkStats `xml:"bytes" structs:",dotflatten"`
	Errors  NetworkStats `xml:"errors" structs:",dotflatten"`
}

type NetworkStats struct {
	Now   int `xml:"now"`
	Total int `xml:"total"`
}

// ServiceProgramm - monit's <service type="7">
// <name>who</name>
// <collected_sec>1468352460</collected_sec>
// <collected_usec>313722</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <program>
// <started>1468352460</started>
// <status>0</status>
// <output>
// <![CDATA[ root ]]>
// </output>
// </program>
// </service>
type ServiceProgramm struct {
	Service `structs:",flatten"`
	Program struct {
		Started uint64 `xml:"started"`
		Status  int    `xml:"status"`
		Output  string `xml:"output"`
	} `xml:"program" structs:",flatten"`
}

// ServiceHost - monit's <service type="4">
type ServiceHost struct {
	Service `structs:",flatten"`
	ICMP    SimpleCheck `xml:"icmp,omitempty" structs:",dotflatten"`
	Port    []Check     `xml:"port" structs:",dotflatten"`
}

type SimpleCheck struct {
	Type         string  `xml:"type"`
	ResponseTime float64 `xml:"responsetime"`
}

type Check struct {
	Hostname   string `xml:"hostname,omitempty"`
	Number     string `xml:"portnumber,omitempty"`
	UnixSocket string `xml:"unixsocket"`
	Request    string `xml:"request,omitempty"`
	Protocol   string `xml:"protocol,omitempty"`
	SimpleCheck
}

// ServiceSystem - monit's <service type="5">
type ServiceSystem struct {
	Service `structs:",flatten"`
	System  struct {
		Load struct {
			Average01 float32 `xml:"avg01"`
			Average05 float32 `xml:"avg05"`
			Average15 float32 `xml:"avg15"`
		} `xml:"load" structs:",dotflatten"`
		CPU struct {
			User   float32 `xml:"user"`
			System float32 `xml:"system"`
			Wait   float32 `xml:"wait"`
		} `xml:"cpu" structs:",dotflatten"`
		Memory struct {
			Percent  float32 `xml:"percent"`
			KiloByte int     `xml:"kilobyte"`
		} `xml:"memory" structs:",dotflatten"`
		Swap struct {
			Percent  float32 `xml:"percent"`
			KiloByte int     `xml:"kilobyte"`
		} `xml:"swap" structs:",dotflatten"`
	} `xml:"system" structs:",flatten"`
}

// ServiceProcess - monit's <service type="3">
type ServiceProcess struct {
	Service  `structs:",flatten"`
	PID      int `xml:"pid"`
	PPID     int `xml:"ppid"`
	UID      int `xml:"uid"`
	EUID     int `xml:"euid"`
	GID      int `xml:"gid"`
	Uptime   int `xml:"uptime"`
	Threads  int `xml:"threads"`
	Children int `xml:"children"`
	Memory   struct {
		Percent      float32 `xml:"percent"`
		PercentTotal float32 `xml:"percenttotal"`
		// TODO maybe int64 will be needed
		KiloByte      int `xml:"kilobyte"`
		KiloByteTotal int `xml:"kilobytetotal"`
	} `xml:"memory" structs:",dotflatten"`
	CPU struct {
		Percent      float32 `xml:"percent"`
		PercentTotal float32 `xml:"percenttotal"`
	} `xml:"cpu" structs:",dotflatten"`
	UNIX struct {
		Path         string `xml:"path"`
		Protocol     string `xml:"protocol"`
		ResponseTime string `xml:"responsetime"`
	} `xml:"unix,omitempty" structs:",dotflatten"`
}

// Service - common service structure for all subservices
type Service struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
	// TODO maybe int64 will be needed
	CollectedSec  int `xml:"collected_sec"`
	CollectedUSec int `xml:"collected_usec"`
	Status        int `xml:"status"`
	StatusHint    int `xml:"status_hint"`
	Monitor       int `xml:"monitor"`
	MonitorMode   int `xml:"monitormode"`
	PendingAction int `xml:"pendingaction"`
}
