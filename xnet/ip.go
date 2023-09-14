package xnet

import (
	"net"
	"strconv"
	"strings"
)

func NextIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

type CIDR struct {
	IPList      []string
	Subnet      string
	BroadcastIP string
}

// ParseIPCIDR 解析网段的所有IP地址
// Example: 192.168.1.1/28
func ParseIPCIDR(s string) (cidr CIDR) {
	if ip := net.ParseIP(s); ip != nil {
		cidr.IPList = append(cidr.IPList, ip.String())
		return cidr
	}
	ipAddr, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return cidr
	}
	startIP := ipNet.IP.Mask(ipNet.Mask)
	endIP := ipNet.IP.Mask(ipNet.Mask)
	for i := range startIP {
		endIP[i] |= ^ipNet.Mask[i]
	}
	for ip := startIP.Mask(ipNet.Mask); ipNet.Contains(ip); NextIP(ip) {
		// 去掉网络地址
		if ip.Equal(ipAddr.Mask(ipNet.Mask)) {
			cidr.Subnet = ip.String()
			continue
		}
		// 去掉广播地址
		if ip.Equal(endIP) {
			cidr.BroadcastIP = ip.String()
			break
		}
		cidr.IPList = append(cidr.IPList, ip.String())
	}
	return cidr
}

// IsIP 判断是否是 IP
func IsIP(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	}
	return false
}

// IsCIDR 判断是否是网段
func IsCIDR(curd string) bool {
	_, _, err := net.ParseCIDR(curd)
	if err != nil {
		return false
	}
	return true
}

// IsIPStartLtEnd 判断左边 IP 是否小于右边 IP
func IsIPStartLtEnd(start, end net.IP) bool {
	for i := 0; i < len(start); i++ {
		if start[i] < end[i] {
			return true
		} else if start[i] > end[i] {
			return false
		}
	}
	return false

}

// IsIPRange 判断是否是 IP 范围
func IsIPRange(ipRange string) bool {
	ipSplit := strings.Split(ipRange, "-")
	if len(ipSplit) != 2 {
		return false
	}
	startIP := net.ParseIP(ipSplit[0])
	endIP := net.ParseIP(ipSplit[1])
	endIPInt, _ := strconv.Atoi(ipSplit[1])
	if startIP == nil {
		return false
	}
	if endIP == nil {
		if endIPInt > 0 && endIPInt <= 255 {
			return true
		}
		return false
	}
	if !IsIPStartLtEnd(startIP, endIP) {
		return false
	}
	return true
}

// ParseIPRange 解析 IP 范围的所有 IP
// Example: 192.168.1.1-192.168.1.10 192.168.1.1-10
func ParseIPRange(ipRange string) (ipList []string) {
	if !IsIPRange(ipRange) {
		return ipList
	}
	ipSplit := strings.Split(ipRange, "-")

	startIP := net.ParseIP(ipSplit[0])
	endIP := net.ParseIP(ipSplit[1])
	endIPInt, _ := strconv.Atoi(ipSplit[1])
	if endIPInt > 0 && endIPInt <= 255 {
		endIP = net.ParseIP(startIP.String())
		endIP[len(endIP)-1] = byte(endIPInt)
	}
	for ip := startIP; ip.String() != endIP.String(); NextIP(ip) {
		ipList = append(ipList, ip.String())
	}
	ipList = append(ipList, endIP.String())
	return ipList
}

// ParseIPList 生成 IP 列表，可以传入 IP/IP Range/CIDR
func ParseIPList(ip string) []string {
	if IsCIDR(ip) || IsIP(ip) {
		return ParseIPCIDR(ip).IPList
	}
	if IsIPRange(ip) {
		return ParseIPRange(ip)
	}
	return []string{}
}
