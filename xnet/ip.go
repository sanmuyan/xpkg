package xnet

import "net"

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// ParseIPList 解析传入的字符串是单个 IP 地址还是网段，如果是网段会解析出该网段的所有IP地址
func ParseIPList(s string) []string {
	var ips []string
	if ip := net.ParseIP(s); ip != nil {
		ips = append(ips, ip.String())
		return ips
	}
	_, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil
	}
	minIP := ipNet.IP.Mask(ipNet.Mask)
	maxIP := ipNet.IP.Mask(ipNet.Mask)
	for i := range minIP {
		maxIP[i] |= ^ipNet.Mask[i]
	}
	for ip := minIP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
		if ip.Equal(maxIP) {
			break
		}
	}
	return ips
}
