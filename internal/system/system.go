package system

import (
	"net"
	"os"
	"runtime"
	"time"
)

var startTime time.Time
var (
	Version   string = ""
	GitCommit string = ""
)

func init() {
	startTime = time.Now()
}

// Uptime returns Uptime for application rounded to seconds
func Uptime() time.Duration {
	return time.Since(startTime).Round(time.Second)
}

type System struct {
	PID             int
	Version         string
	GitCommit       string
	Uptime          string
	RuntimeVersion  string
	OperatingSystem string
	Architecture    string
	NumCPU          int
	Hostname        string
	IP              net.IP
	Environment     []string
}

func NewSystem() System {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = err.Error()
	}
	s := System{
		PID:             os.Getpid(),
		Version:         Version,
		GitCommit:       GitCommit,
		Uptime:          Uptime().String(),
		RuntimeVersion:  runtime.Version(),
		OperatingSystem: runtime.GOOS,
		Architecture:    runtime.GOARCH,
		NumCPU:          runtime.NumCPU(),
		Hostname:        hostname,
		IP:              myIP(),
		Environment:     os.Environ(),
	}
	return s
}

// myIP returns the first found IP of the system or nil of nothing found.
// Loopbacks are ignored.
func myIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}
