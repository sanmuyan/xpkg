package xnet

import (
	"bufio"
	"encoding/binary"
	"github.com/sanmuyan/xpkg/xconstant"
	"io"
)

// Encode 用于 TCP 包编码
func Encode(data []byte) ([]byte, error) {
	length := uint32(len(data))

	pkg := make([]byte, 4+len(data))

	binary.LittleEndian.PutUint32(pkg[:4], length)

	copy(pkg[4:], data)

	return pkg, nil
}

// Decode 用于 TCP 包解码
func Decode(reader *bufio.Reader) ([]byte, error) {
	lengthByte := make([]byte, 4)
	if _, err := io.ReadFull(reader, lengthByte); err != nil {
		return nil, err
	}

	length := binary.LittleEndian.Uint32(lengthByte)

	if length > xconstant.PackageMaxLength {
		return nil, xconstant.PackageTooLengthError
	}

	pkg := make([]byte, length)
	if _, err := io.ReadFull(reader, pkg); err != nil {
		return nil, err
	}

	return pkg, nil
}
