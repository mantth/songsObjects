package conf

import (
	"net"
	"strings"
)

const (
	CanNotGetIP = "get ip failed..."
)

func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return CanNotGetIP, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.IP.String(), ":")[0], nil
}
