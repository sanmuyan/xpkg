package xnet

import (
	"testing"
)

func TestGetDataSpeed(t *testing.T) {
	if GetDataSpeed(1300, 1) != "10.16Kb/s" {
		t.Error("GetDataSpeed() error")
	}

	t.Log(GetDataSpeed(1024*1024, 1))
	t.Log(GetDataSpeed(1024*1021, 1))
}
