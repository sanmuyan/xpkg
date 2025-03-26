package xconstant

import (
	"errors"
)

// 错误定义

var (
	BufferedTooSmallError = errors.New("buffered too small")
	BadParameter          = errors.New("bad parameter")
	PackageTooLengthError = errors.New("package too length")
)

// 常量定义

const (
	// PackageMaxLength 网络数据包允许的最大长度
	PackageMaxLength = 1024 * 1024
)
