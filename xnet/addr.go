package xnet

import "net"

// ReplaceAddrPort 替换地址端口
func ReplaceAddrPort(addr string, newPort string) (string, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	return net.JoinHostPort(host, newPort), nil
}

// ReplaceAddrHost 替换地址主机
func ReplaceAddrHost(addr string, newHost string) (string, error) {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	return net.JoinHostPort(newHost, port), nil
}

// GetAddrHostPort 获取地址主机和端口
func GetAddrHostPort(addr string) (string, string, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", "", nil
	}
	return host, port, nil
}
