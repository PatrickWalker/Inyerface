package main

import (
	"fmt"
	"net"
)

//This prints out all your network interfaces
//Helps you to pick one to use in place of eth0
func main() {
	intfs, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found the following interfaces")
	for _, intf := range intfs {
		fmt.Printf(`
		Interface Name: %v
		Interface Index: %v
		Inteface Hardware Addr: %v
		Interface Flags : %v
		`, intf.Name, intf.Index, intf.HardwareAddr,
			intf.Flags)
	}
}
