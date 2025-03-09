package xnet

import (
	"fmt"
)

func GetDataSpeed(b, s int) string {
	bit := float64(b * 8 / s)
	if bit >= 1024 && bit < 1024*1024 {
		return fmt.Sprintf("%.2fKb/s", bit/1024)
	}
	if bit >= 1024*1024 && bit < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMb/s", bit/(1024*1024))
	}
	if bit >= 1024*1024*1024 {
		return fmt.Sprintf("%.0fGb/s", bit/(1024*1024*1024))
	}
	return fmt.Sprintf("%.0fb/s", bit)
}
