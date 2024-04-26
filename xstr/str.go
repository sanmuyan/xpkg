package xstr

import (
	"strings"
)

func BuilderStr(args ...string) string {
	var sb strings.Builder
	for _, str := range args {
		sb.WriteString(str)
	}
	return sb.String()
}
