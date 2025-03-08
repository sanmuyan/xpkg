package xnet

import (
	"fmt"
	"strconv"
	"strings"
)

func GetDataSpeed(b, s int) string {
	bit := b * 8 / s
	if bit >= 1024 && bit < 1024*1024 {
		f := float64(bit) / float64(1024)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		i := strings.Split(s, ".")
		if i[1] == "00" {
			return fmt.Sprintf("%sKbps/s", i[0])
		}
		return fmt.Sprintf("%sKbps/s", s)
	}
	if bit >= 1024*1024 && bit < 1024*1024*1024 {
		f := float64(bit) / float64(1024*1024)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		i := strings.Split(s, ".")
		if i[1] == "00" {
			return fmt.Sprintf("%sKbps/s", i[0])
		}
		return fmt.Sprintf("%sMbps/s", s)
	}
	if bit >= 1024*1024*1024 {
		f := float64(bit) / float64(1024*1024*1024)
		s := strconv.FormatFloat(f, 'f', 2, 64)
		i := strings.Split(s, ".")
		if i[1] == "00" {
			return fmt.Sprintf("%sKbps/s", i[0])
		}
		return fmt.Sprintf("%sGbps/s", s)
	}
	return fmt.Sprintf("%dbps/s", bit)
}
