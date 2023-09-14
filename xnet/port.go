package xnet

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func isOnePort(port string) bool {
	p, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if p >= 0 && p <= 65535 {
		return true
	}
	return false
}

func isPortRange(port string) bool {
	if strings.Contains(port, "-") {
		portRange := strings.Split(port, "-")
		if len(portRange) == 2 {
			if isOnePort(portRange[0]) && isOnePort(portRange[1]) {
				minPort, _ := strconv.Atoi(portRange[0])
				maxPort, _ := strconv.Atoi(portRange[1])
				if minPort <= maxPort {
					return true
				}
			}
		}
	}
	return false
}

func isMultiplePorts(port string) bool {
	if strings.Contains(port, ",") {
		for _, v := range strings.Split(port, ",") {
			if !isOnePort(v) && !isPortRange(v) {
				return false
			}
		}
		return true
	}
	return false

}

// IsPort 判断是否是网络端口，格式支持 80 1-100 22,23
func IsPort[T ~int | ~string](port T) bool {
	p := fmt.Sprint(port)
	if isOnePort(p) || isMultiplePorts(p) || isPortRange(p) {
		return true
	}
	return false
}

func generatePortRange(port string) (portList []int) {
	portRange := strings.Split(port, "-")
	minPort, _ := strconv.Atoi(portRange[0])
	maxPort, _ := strconv.Atoi(portRange[1])
	for i := minPort; i <= maxPort; i++ {
		portList = append(portList, i)
	}
	return portList
}

func generateMultiplePorts(port string) (portList []int) {
	for _, v := range strings.Split(port, ",") {
		if isPortRange(v) {
			portList = append(portList, generatePortRange(v)...)
		} else {
			p, _ := strconv.Atoi(v)
			portList = append(portList, p)
		}
	}
	return portList
}

// GeneratePorts 生成端口范围的所有端口
func GeneratePorts[T ~int | ~string](port T) (portList []int) {
	p := fmt.Sprint(port)
	if IsPort(p) {
		if isOnePort(p) {
			_p, _ := strconv.Atoi(p)
			portList = append(portList, _p)
		}
		if isPortRange(p) {
			portList = append(portList, generatePortRange(p)...)
		}
		if isMultiplePorts(p) {
			portList = append(portList, generateMultiplePorts(p)...)
		}
	}
	for i := 0; i < len(portList); i++ {
		for j := i + 1; j < len(portList); j++ {
			if portList[i] == portList[j] {
				portList = append(portList[:j], portList[j+1:]...)
			}
		}
	}
	sort.Ints(portList)
	return portList
}

// IsAllowPort 判断是端口是否在某个端口范围
func IsAllowPort[T ~int | ~string](allowPorts, port T) bool {
	_port := fmt.Sprint(port)
	if !isOnePort(_port) {
		return false
	}
	portInt, _ := strconv.Atoi(_port)
	ports := GeneratePorts(allowPorts)
	if len(ports) == 0 {
		return false
	}
	if portInt >= ports[0] && portInt <= ports[len(ports)-1] {
		return true
	}
	return false
}
