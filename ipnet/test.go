package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "192.168.100.167/24"
	x, y, z := net.ParseCIDR(ip)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
