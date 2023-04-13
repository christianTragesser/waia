package system

import (
	"log"
	"net"
	"os"
)

type System struct {
	Hostname string
	Address  string
	Version  string
	BgColor  string
	TxtColor string
}

type SystemAPI interface {
	getIP() string
	getColors() (string, string)
}

func (s *System) GetInfo() (System, error) {
	var err error

	s.Hostname, err = os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	s.Address = s.getIP()

	s.BgColor, s.TxtColor = s.getColors()

	return *s, err
}

func (s *System) getIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ipAddr := conn.LocalAddr().(*net.UDPAddr)

	return ipAddr.IP.String()
}

func (s *System) getColors() (string, string) {
	txt := "black"

	bg, bgSet := os.LookupEnv("BG_COLOR")
	if !bgSet || bg == "" {
		bg = "white"
	}

	if bg != "white" {
		txt = "white"
	}

	return bg, txt
}
