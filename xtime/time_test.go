package xtime

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	_t, err := StrUnitToTime("1s")
	if err != nil {
		t.Error(err)
	}
	if _t != 1000*time.Millisecond {
		t.Error(_t)
	}
	_st := TimeToStrUnit(1111 * time.Millisecond)
	if err != nil {
		t.Error(err)
	}
	if _st != "1.111s" {
		t.Error(_st)
	}
	if TimeToStrUnitTrim(1111*time.Millisecond, 0) != "1s" {
		t.Error()
	}
}
