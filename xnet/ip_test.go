package xnet

import (
	"slices"
	"testing"
)

func TestParseIPList(t *testing.T) {
	ipList := []string{"192.168.1.1", "192.168.1.2"}
	if slices.Compare(ParseIPList("192.168.1.0/30"), ipList) != 0 {
		t.Errorf("test ParseIPList failed")
	}
	if slices.Compare(ParseIPList("192.168.1.1-192.168.1.2"), ipList) != 0 {
		t.Errorf("test ParseIPList failed")
	}
	ipCIDR := ParseIPCIDR("192.168.1.1/28")
	if ipCIDR.IPList[0] != "192.168.1.1" || ipCIDR.IPList[len(ipCIDR.IPList)-1] != "192.168.1.14" || ipCIDR.Subnet != "192.168.1.0" || ipCIDR.BroadcastIP != "192.168.1.15" {
		t.Errorf("test ParseIPCIDR failed %+v", ipCIDR)
	}
}
