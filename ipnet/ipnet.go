package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
)

const vagrantIP = `10.0.2.15`

const getInterfaceSc = `IP={{.ip}} 
INTERFACE=$(ip addr | grep $IP)
INTERFACE=${INTERFACE##*" "}
echo ${INTERFACE}
`

func IsVirtual(ipnet *net.IPNet) bool {
	return interfaceContainsString(ipnet, "veth") || interfaceContainsString(ipnet, "virbr")
}

func IsBridge(ipnet *net.IPNet) bool {
	return interfaceContainsString(ipnet, "br-")
}

func IsFlannel(ipnet *net.IPNet) bool {
	return interfaceContainsString(ipnet, "flannel")
}

func IsCalico(ipnet *net.IPNet) bool {
	return interfaceContainsString(ipnet, "cali")
}

func IsTunl(ipnet *net.IPNet) bool {
	return interfaceContainsString(ipnet, "tunl")
}

func interfaceContainsString(ipnet *net.IPNet, str string) bool {
	ip := ipnet.IP.String()
	iface := getInterfaceByIP(ip)
	fmt.Println("ip: ", ip)
	fmt.Println("interface: ", iface)
	return strings.Contains(iface, str)
}

func getInterfaceByIP(ip string) string {
	cmd := string(getInterfaceSc)
	cmd = strings.Replace(cmd, "{{.ip}}", ip, -1)
	ret, _ := execCmd(cmd)
	return ret
}

func execCmd(strCommand string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return "", err
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return "", err
	}
	str := string(out_bytes)
	return str, nil
}

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

func GetIntranetIP() []string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			fmt.Println("IP: ", ipnet.IP.String())
			fmt.Println("Mask: ", addr.(*net.IPNet).Mask)
			ips = append(ips, ipnet.IP.String())
		}
		fmt.Println("---")
	}
	return ips
}

func getIPNet(ip string) *net.IPNet {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.String() == ip {
			return ipnet
		}
	}
	return nil
}

func main() {
	ips := GetIntranetIP()
	fmt.Println(ips)
	fmt.Println("===")
	ip := "192.168.100.167"
	ipnet := getIPNet(ip)
	fmt.Println(*ipnet)
}
