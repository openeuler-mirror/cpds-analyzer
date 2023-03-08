package net

import (
	"net"
)

func IsValidIPAdress(IPAddress string) bool {
	return net.ParseIP(IPAddress) != nil
}

// 0 is considered as a non valid port
func IsValidPort(port int) bool {
	return port > 0 && port < 65535
}
