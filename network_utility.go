package golib

import (
	"net"
)

func LocalIp() (string, error) {
	ip := "127.0.0.1"

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ip, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()

				break
			}
		}
	}

	return ip, nil
}
