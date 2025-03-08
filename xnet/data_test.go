package xnet

import (
	"testing"
)

func TestGetDataSpeed(t *testing.T) {
	if GetDataSpeed(1300, 1) != "10.16Kb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024, 1) != "8Kb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024*1024, 1) != "8Mb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024*1300, 1) != "10.16Mb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024*1024*1024, 1) != "8Gb/s" {
		t.Error("GetDataSpeed() error")
	}
}
