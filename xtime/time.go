package xtime

import (
	"github.com/sanmuyan/xpkg/xconstant"
	"strconv"
	"time"
)

// TimeUnitConv 支持把带单位的字符串转换为毫秒
// 比如把 1s 转换为 1000 毫秒
// Example: 1s 3w 2d
func TimeUnitConv(timeStr string) (time.Duration, error) {
	timeInt, err := strconv.ParseInt(timeStr, 10, 64)
	if err == nil {
		return time.Duration(timeInt) * time.Millisecond, nil
	}
	units := make(map[string]int64)
	units["ms"] = 0
	units["s"] = 1000
	units["m"] = units["s"] * 60
	units["h"] = units["m"] * 60
	units["d"] = units["h"] * 24
	units["w"] = units["d"] * 7
	// 一年按365.25天
	// 一个月按30.4375天
	units["M"] = units["d"]*30 + (units["s"] * 37638)
	units["y"] = units["d"]*365 + (units["h"] * 6)
	if len(timeStr) != 3 && len(timeStr) != 2 {
		return 0, xconstant.BadParameter
	}
	unitNumber := timeStr[0]
	unitNumberInt, err := strconv.ParseInt(string(unitNumber), 10, 64)
	if err != nil {
		return 0, xconstant.BadParameter
	}
	unit := timeStr[1:]
	if _, ok := units[unit]; !ok {
		return 0, xconstant.BadParameter
	}
	return time.Duration(unitNumberInt*units[unit]) * time.Millisecond, nil
}
