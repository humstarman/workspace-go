package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const vagrantIP = `10.0.2.15`

func IsDocker0(ipnet *net.IPNet) bool {
	ips, _ := GetIpByInterface("docker0")
	str := fmt.Sprintf("%v", ips)
	fmt.Println(str)
	ip := ipnet.IP.String()
	return strings.Contains(str, ip)
}

func IsVagrant(ipnet *net.IPNet) bool {
	return ipnet.IP.String() == string(vagrantIP)
}

func IsMaskEmpty(ipnet *net.IPNet) bool {
	return fmt.Sprintf("%v", ipnet.Mask) == "ffffffff"
}

func GetIpByInterface(name string) ([]string, error) {
	var ips []string
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}
	return ips, nil
}

func GetIntranetIP() []*net.IPNet {
	var ipnets []*net.IPNet
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && !IsDocker0(ipnet) && !IsVagrant(ipnet) && !IsMaskEmpty(ipnet) {
			fmt.Println("IP: ", ipnet.IP.String())
			fmt.Println("Mask: ", addr.(*net.IPNet).Mask)
			ipnets = append(ipnets, ipnet)

		}
		fmt.Println("---")
	}
	return ipnets
}

func main() {
	x := GetIntranetIP()
	fmt.Println(len(x))
}
