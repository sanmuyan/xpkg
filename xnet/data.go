package xnet

import (
	"fmt"
	"strings"
)

func GetDataSpeed(b, s int) string {
	bit := float64(b * 8 / s)
	if bit >= 1024 && bit < 1024*1024 {
		f := fmt.Sprintf("%.2fKb/s", bit/1024)
		sl := strings.Split(f, ".")
		if sl[1] == "00Kb/s" {
			return fmt.Sprintf("%.0fKb/s", bit/1024)
		}
		return f
	}
	if bit >= 1024*1024 && bit < 1024*1024*1024 {
		f := fmt.Sprintf("%.2fMb/s", bit/(1024*1024))
		sl := strings.Split(f, ".")
		if sl[1] == "00Mb/s" {
			return fmt.Sprintf("%.0fMb/s", bit/(1024*1024))
		}
		return f
	}
	if bit >= 1024*1024*1024 {
		f := fmt.Sprintf("%.2fMb/s", bit/(1024*1024*1024))
		sl := strings.Split(f, ".")
		if sl[1] == "00Mb/s" {
			return fmt.Sprintf("%.0fGb/s", bit/(1024*1024*1024))
		}
		return f
	}
	return fmt.Sprintf("%.0fb/s", bit)
}
