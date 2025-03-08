package xnet

import (
	"testing"
)

func TestGetDataSpeed(t *testing.T) {
	if GetDataSpeed(1300, 1) != "10.16Kb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024, 1) != "8.00Kb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024*1024, 1) != "8.00Mb/s" {
		t.Error("GetDataSpeed() error")
	}
	if GetDataSpeed(1024*1024*1024, 1) != "8.00Gb/s" {
		t.Error("GetDataSpeed() error")
	}
}
