package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ip := "192.168.100.167"
	mask := 24
	cidr := fmt.Sprintf("%v/%v", ip, mask)
	n := getCidrHostNum(mask)
	fmt.Printf("hosts num: %v\n", n)
	fmt.Printf("cidr: %v\n", cidr)
	looper1(cidr)
	minIp, maxIp := getCidrIpRange(cidr)
	fmt.Println("CIDR最小IP：", minIp, " CIDR最大IP：", maxIp)
}

func looper1(cidr string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)
	fmt.Printf("seg2 from %v to %v\n", seg2MinIp, seg2MaxIp)
	fmt.Printf("seg3 from %v to %v\n", seg3MinIp, seg3MaxIp)
	fmt.Printf("seg4 from %v to %v\n", seg4MinIp, seg4MaxIp)
	for i := seg2MinIp; i <= seg2MaxIp; i++ {
		for j := seg3MinIp; j <= seg3MaxIp; j++ {
			for k := seg4MinIp; k <= seg4MaxIp; k++ {
				tmp := fmt.Sprintf("%v.%v.%v.%v", ipSegs[0], i, j, k)
				fmt.Println(tmp)
			}
		}
	}
	fmt.Printf("seg2 from %v to %v\n", seg2MinIp, seg2MaxIp)
	fmt.Printf("seg3 from %v to %v\n", seg3MinIp, seg3MaxIp)
	fmt.Printf("seg4 from %v to %v\n", seg4MinIp, seg4MaxIp)
}

func looper(cidr string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)
	fmt.Printf("seg2 from %v to %v\n", seg2MinIp, seg2MaxIp)
	fmt.Printf("seg3 from %v to %v\n", seg3MinIp, seg3MaxIp)
	fmt.Printf("seg4 from %v to %v\n", seg4MinIp, seg4MaxIp)
	if seg2MinIp != seg2MaxIp {
		for i := seg2MinIp; i <= seg2MaxIp; i++ {
			for j := seg3MinIp; j <= seg3MaxIp; j++ {
				for k := seg4MinIp; k <= seg4MaxIp; k++ {
					tmp := fmt.Sprintf("%v.%v.%v.%v", ipSegs[0], i, j, k)
					fmt.Println(tmp)
				}
			}
		}
	} else if seg3MinIp != seg3MaxIp {
		for j := seg3MinIp; j <= seg3MaxIp; j++ {
			for k := seg4MinIp; k <= seg4MaxIp; k++ {
				tmp := fmt.Sprintf("%v.%v.%v.%v", ipSegs[0], ipSegs[1], j, k)
				fmt.Println(tmp)
			}
		}
	} else {
		for k := seg4MinIp; k <= seg4MaxIp; k++ {
			tmp := fmt.Sprintf("%v.%v.%v.%v", ipSegs[0], ipSegs[1], ipSegs[2], k)
			fmt.Println(tmp)
		}
	}
	fmt.Printf("seg2 from %v to %v\n", seg2MinIp, seg2MaxIp)
	fmt.Printf("seg3 from %v to %v\n", seg3MinIp, seg3MaxIp)
	fmt.Printf("seg4 from %v to %v\n", seg4MinIp, seg4MaxIp)
}

func getCidrIpRange(cidr string) (string, string) {
	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, _ := strconv.Atoi(strings.Split(cidr, "/")[1])
	seg2MinIp, seg2MaxIp := getIpSeg2Range(ipSegs, maskLen)
	seg3MinIp, seg3MaxIp := getIpSeg3Range(ipSegs, maskLen)
	seg4MinIp, seg4MaxIp := getIpSeg4Range(ipSegs, maskLen)
	//ipPrefix := ipSegs[0] + "." + ipSegs[1]
	ipPrefix := ipSegs[0]
	return fmt.Sprintf("%v.%v.%v.%v", ipPrefix, seg2MinIp, seg3MinIp, seg4MinIp), fmt.Sprintf("%v.%v.%v.%v", ipPrefix, seg2MaxIp, seg3MaxIp, seg4MaxIp)
}

func getCidrHostNum(maskLen int) uint {
	cidrIpNum := uint(0)
	i := uint(32 - maskLen - 1)
	for ; i >= 1; i-- {
		cidrIpNum += 1 << i
	}
	return cidrIpNum
}

func getCidrIpMask(maskLen int) string {
	cidrMask := ^uint32(0) << uint(32-maskLen)
	fmt.Println(fmt.Sprintf("%b \n", cidrMask))
	cidrMaskSeg1 := uint8(cidrMask >> 24)
	cidrMaskSeg2 := uint8(cidrMask >> 16)
	cidrMaskSeg3 := uint8(cidrMask >> 8)
	cidrMaskSeg4 := uint8(cidrMask & uint32(255))
	return fmt.Sprintf("%v.%v.%v.%v", cidrMaskSeg1, cidrMaskSeg2, cidrMaskSeg3, cidrMaskSeg4)
}

func getIpSeg2Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 16 {
		segIp, _ := strconv.Atoi(ipSegs[1])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[1])
	return getIpSegRange(uint8(ipSeg), uint8(16-maskLen))
}

func getIpSeg3Range(ipSegs []string, maskLen int) (int, int) {
	if maskLen > 24 {
		segIp, _ := strconv.Atoi(ipSegs[2])
		return segIp, segIp
	}
	ipSeg, _ := strconv.Atoi(ipSegs[2])
	return getIpSegRange(uint8(ipSeg), uint8(24-maskLen))
}

func getIpSeg4Range(ipSegs []string, maskLen int) (int, int) {
	ipSeg, _ := strconv.Atoi(ipSegs[3])
	segMinIp, segMaxIp := getIpSegRange(uint8(ipSeg), uint8(32-maskLen))
	return segMinIp + 1, segMaxIp
}

func getIpSegRange(userSegIp, offset uint8) (int, int) {
	var ipSegMax uint8 = 255
	netSegIp := ipSegMax << offset
	segMinIp := netSegIp & userSegIp
	segMaxIp := userSegIp&(255<<offset) | ^(255 << offset)
	return int(segMinIp), int(segMaxIp)
}
