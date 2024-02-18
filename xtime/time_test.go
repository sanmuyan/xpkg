package xtime

import (
	"testing"
	"time"
)

func TestTimeUnitConversion(t *testing.T) {
	_t, err := TimeUnitConv("1s")
	if err != nil {
		t.Error(err)
	}
	if _t != 1000*time.Millisecond {
		t.Error(_t)
	}
}
