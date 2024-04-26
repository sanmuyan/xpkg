package xnet

import (
	"testing"
)

func TestGetDataSpeed(t *testing.T) {
	if GetDataSpeed(1300, 1) != "10.16Kbps/s" {
		t.Error("GetDataSpeed() error")
	}
}
