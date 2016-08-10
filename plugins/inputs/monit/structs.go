package monit

type Manager struct {
	Url      string
	User     string
	Password string
}

type Monit struct {
	Manager
	Server      Instance            `xml:"server"`
	Platform    System              `xml:"platform"`
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

type Instance struct {
	ID            string `xml:"id"`
	Incarnation   string `xml:"incarnation"`
	Version       string `xml:"version"`
	Uptime        string `xml:"uptime"`
	Poll          string `xml:"poll"`
	StartDelay    string `xml:"startdelay"`
	LocalHostName string `xml:"localhostname"`
	ControlFile   string `xml:"controlfile"`
	Httpd         Httpd  `xml:"httpd"`
}

type Httpd struct {
	Address string `xml:"address"`
	Port    string `xml:"port"`
	SSL     int    `xml:"ssl"`
}

type System struct {
	Name    string `xml:"name"`
	Release string `xml:"release"`
	Version string `xml:"version"`
	Machine string `xml:"machine"`
	CPU     int    `xml:"cpu"`
	Memory  int    `xml:"memory"`
	Swap    int    `xml:"swap"`
}

// ServiceFile - monit's <service type="2">
// <name>monitrc</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>309506</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <mode>600</mode>
// <uid>0</uid>
// <gid>0</gid>
// <timestamp>1468351670</timestamp>
// <size>12375</size>
// </service>
type ServiceFile struct {
	Service
}

// ServiceDirectory - monit's <service type="1">
// <name>bin</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>308827</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <mode>755</mode>
// <uid>0</uid>
// <gid>0</gid>
// <timestamp>1465722239</timestamp>
// </service>
type ServiceDirectory struct {
	Service
}

// ServiceFifo - monit's <service type="6">
// <name>testFifo</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>309509</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <mode>664</mode>
// <uid>1000</uid>
// <gid>1000</gid>
// <timestamp>1468351653</timestamp>
// </service>
type ServiceFifo struct {
	Service
}

// ServiceFilesystem - monit's <service type="0">
// <name>datafs</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>309040</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <mode>660</mode>
// <uid>0</uid>
// <gid>6</gid>
// <flags>4096</flags>
// <block>
// <percent>4.3</percent>
// <usage>8470.9</usage>
// <total>196889.0</total>
// </block>
// <inode>
// <percent>0.6</percent>
// <usage>73691</usage>
// <total>12812288</total>
// </inode>
// </service>
type ServiceFilesystem struct {
	Service
}

// ServiceNet - monit's <service type="8">
// <name>public</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>309479</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
//  <link>
//   <state>1</state>
//   <speed>-1</speed>
//   <duplex>-1</duplex>
//     <download>
//       <packets>
//         <now>0</now>
//         <total>444601</total>
//       </packets>
//       <bytes>
//          <now>0</now>
//          <total>380250398</total>
//       </bytes>
//       <errors>
//         <now>0</now>
//         <total>0</total>
//       </errors>
//     </download>
//     <upload>
//        <packets>
//           <now>0</now>
//           <total>340995</total>
//        </packets>
//        <bytes>
//          <now>0</now>
//          <total>55739459</total>
//        </bytes>
// 		  <errors>
//           <now>0</now>
//           <total>0</total>
//        </errors>
//      </upload>
//  </link>
// </service>
type ServiceNet struct {
	Service
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
	Service
}

// ServiceHost - monit's <service type="4">
// <name>myserver</name>
// <collected_sec>1468351672</collected_sec>
// <collected_usec>309194</collected_usec>
// <status>0</status>
// <status_hint>0</status_hint>
// <monitor>1</monitor>
// <monitormode>0</monitormode>
// <pendingaction>0</pendingaction>
// <icmp>
// <type>Ping</type>
// <responsetime>0.000063</responsetime>
// </icmp>
// </service>
type ServiceHost struct {
	Service
}

// ServiceSystem - monit's <service type="5">
type ServiceSystem struct {
	Service
	System struct {
		Load struct {
			Average01 float32 `xml:"avg01"`
			Average05 float32 `xml:"avg05"`
			Average15 float32 `xml:"avg15"`
		} `xml:"load"`
		CPU struct {
			User   float32 `xml:"user"`
			System float32 `xml:"system"`
			Wait   float32 `xml:"wait"`
		} `xml:"cpu"`
		Memory struct {
			Percent  float32 `xml:"percent"`
			KiloByte int     `xml:"kilobyte"`
		} `xml:"memory"`
		Swap struct {
			Percent  float32 `xml:"percent"`
			KiloByte int     `xml:"kilobyte"`
		} `xml:"swap"`
	} `xml:"system"`
}

// ServiceProcess - monit's <service type="3">
type ServiceProcess struct {
	Service
	PID      int           `xml:"pid"`
	PPID     int           `xml:"ppid"`
	UID      int           `xml:"uid"`
	EUID     int           `xml:"euid"`
	GID      int           `xml:"gid"`
	Uptime   int           `xml:"uptime"`
	Threads  int           `xml:"threads"`
	Children int           `xml:"children"`
	Memory   ProcessMemory `xml:"memory"`
	CPU      ProcessCPU    `xml:"cpu"`
}

type ProcessMemory struct {
	Percent      float32 `xml:"percent"`
	PercentTotal float32 `xml:"percenttotal"`
	// TODO maybe int64 will be needed
	KiloByte      int `xml:"kilobyte"`
	KiloByteTotal int `xml:"kilobytetotal"`
}

type ProcessCPU struct {
	Percent      float32 `xml:"percent"`
	PercentTotal float32 `xml:"percenttotal"`
}

// Service - common service structure for all subservices
type Service struct {
	ServiceType string `xml:"type,attr"`
	Name        string `xml:"name"`
	// TODO maybe int64 will be needed
	CollectedSec  int `xml:"collected_sec"`
	CollectedUSec int `xml:"collected_usec"`
	Status        int `xml:"status"`
	StatusHint    int `xml:"status_hint"`
	Monitor       int `xml:"monitor"`
	MonitorMode   int `xml:"monitormode"`
	PendingAction int `xml:"pendingaction"`
}
