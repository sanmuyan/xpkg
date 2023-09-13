package xnet

import (
	"fmt"
	"testing"
)

func TestIsPort(t *testing.T) {
	if !IsPort("123,456,1-10") || !IsPort(10) || !IsPort("1-10") || IsPort("-1-10") || IsPort(100000) {
		t.Errorf("test IsPort failed")
	}
}

func TestGeneratePorts(t *testing.T) {
	ports := []int{1, 2, 3, 4, 5, 6}
	if fmt.Sprint(ports) != fmt.Sprint(GeneratePorts("1-5,5,6")) {
		t.Errorf("test GeneratePorts failed")
	}
}

func TestIsAllowPort(t *testing.T) {
	if !IsAllowPort("1-100", "5") || IsAllowPort("1-100", "101") {
		t.Errorf("test IsAllowPort failed")
	}
}
