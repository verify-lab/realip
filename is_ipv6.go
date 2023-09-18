package realip

import (
	"net"
	"strings"
)

func IsIPv6(ip net.IP) bool {
	return IsStringIPv6(ip.String())
}

func IsStringIPv6(ip string) bool {
	return strings.Contains(ip, ":")
}
