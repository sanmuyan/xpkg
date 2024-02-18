package xconstant

import (
	"errors"
)

// 错误定义

var (
	BufferedTooSmallError = errors.New("buffered too small")
	BadParameter          = errors.New("bad parameter")
)
