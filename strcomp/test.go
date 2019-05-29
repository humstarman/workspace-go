package main

import (
	"fmt"
)

func main() {
	ip1 := "10.1.0.2"
	ip2 := "10.1.0.4"
	ip3 := "192.168.100.167"
	fmt.Println(ip1 < ip2)
	fmt.Println(ip1 > ip2)
	fmt.Println(ip1 == ip2)
	fmt.Println(ip1 < ip3)
	fmt.Println(ip1 > ip3)
	fmt.Println(ip1 == ip3)
}
