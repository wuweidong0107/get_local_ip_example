package main

import (
	"fmt"
	"net"
)

func main() {
	// get list of available addresses
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet) // Type Assertion, 判断 addr 是否为类型 *net.IPNet
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
