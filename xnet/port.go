package xnet

import (
	"fmt"
	"strconv"
	"strings"
)

func isOnePort(port string) bool {
	_port, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	if _port >= 0 && _port <= 65535 {
		return true
	}
	return false
}

func isPortRange(port string) bool {
	if strings.Contains(port, "-") {
		portRange := strings.Split(port, "-")
		if len(portRange) == 2 {
			if isOnePort(portRange[0]) && isOnePort(portRange[1]) {
				if portRange[0] <= portRange[1] {
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

func IsPort[T ~int | ~string](port T) bool {
	_port := fmt.Sprint(port)
	if isOnePort(_port) || isMultiplePorts(_port) || isPortRange(_port) {
		return true
	}
	return false
}

func generatePortRange(port string) (res []int) {
	portRange := strings.Split(port, "-")
	minPort, _ := strconv.Atoi(portRange[0])
	maxPort, _ := strconv.Atoi(portRange[1])
	for i := minPort; i <= maxPort; i++ {
		res = append(res, i)
	}
	return res
}

func generateMultiplePorts(port string) (res []int) {
	for _, v := range strings.Split(port, ",") {
		if isPortRange(v) {
			res = append(res, generatePortRange(v)...)
		} else {
			_port, _ := strconv.Atoi(v)
			res = append(res, _port)
		}
	}
	return res
}

func GeneratePorts[T ~int | ~string](port T) (res []int) {
	p := fmt.Sprint(port)
	if IsPort(p) {
		if isOnePort(p) {
			_p, _ := strconv.Atoi(p)
			res = append(res, _p)
		}
		if isPortRange(p) {
			res = append(res, generatePortRange(p)...)
		}
		if isMultiplePorts(p) {
			res = append(res, generateMultiplePorts(p)...)
		}
	}
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] == res[j] {
				res = append(res[:j], res[j+1:]...)
			}
		}
	}
	return res
}

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
	for _, v := range ports {
		if v == portInt {
			return true
		}
	}
	return false
}
