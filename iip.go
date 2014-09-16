package main

import (
	"fmt"
	"net"
)

type CIDRAddress struct {
	Address net.IP
	Network *net.IPNet
}

func main() {

	internalAddresses := [...]string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	cidrAddresses := make([]*CIDRAddress, 0)
	for _, addr := range internalAddresses {
		a := &CIDRAddress{}
		a.Address, a.Network, _ = net.ParseCIDR(addr)
		cidrAddresses = append(cidrAddresses, a)
	}

	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, address := range interfaces {
		ip, _, err := net.ParseCIDR(address.String())
		if err != nil {
			panic(err)
		}

		// No default mask? Not IPv4
		if ip.DefaultMask() == nil {
			continue
		}

		for _, check := range cidrAddresses {
			if check.Network.Contains(ip) {
				fmt.Println(ip)
				return
			}
		}
	}

}
