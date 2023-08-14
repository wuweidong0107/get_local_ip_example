package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(interfaces)

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPAddr:
				ip = v.IP
			case *net.IPNet:
				ip = v.IP
			default:
				continue
			}
			if ip.To4() != nil && !ip.IsLoopback() {
				fmt.Println(i.Name + ": " + ip.String())
			}
		}
	}
}
