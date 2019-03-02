package main

import (
	"fmt"
	"net"

	hierr "github.com/reconquest/hierr-go"
)

func getIP(netdev string, IPv4Flag bool) (string, error) {

	ifaces, err := net.InterfaceByName(netdev)
	if err != nil {
		return "", hierr.Errorf(err, "can't get network interfaces")
	}
	addrs, err := ifaces.Addrs()
	if err != nil {
		return "", hierr.Errorf(err, "can't get addresses for interface")
	}

	var IP string

	for _, a := range addrs {
		ips, _, err := net.ParseCIDR(a.String())
		if err != nil {
			return "", hierr.Errorf(err, "can't parse cidr")
		}

		switch {
		case IPv4Flag:
			if ips.To4() != nil {
				IP = ips.To4().String()
				break
			}
		case !IPv4Flag:
			if ips.To4() == nil {
				IP = ips.To16().String()
				break
			}
		}
	}

	if len(IP) == 0 {
		return "", fmt.Errorf("can't found IP address on interface")
	}

	return IP, nil
}
