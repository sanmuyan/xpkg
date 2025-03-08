package xnet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

// Encode 用于 TCP 包编码
func Encode(data []byte) ([]byte, error) {
	length := uint32(len(data))
	pkg := new(bytes.Buffer)

	if err := binary.Write(pkg, binary.LittleEndian, length); err != nil {
		return nil, err
	}
	pkg.Write(data)
	return pkg.Bytes(), nil
}

// Decode 用于 TCP 包解码
func Decode(reader *bufio.Reader) ([]byte, error) {
	lengthByte := make([]byte, 4)
	_, err := io.ReadFull(reader, lengthByte)
	if err != nil {
		return nil, err
	}

	length := binary.LittleEndian.Uint32(lengthByte)

	pkg := make([]byte, length)
	_, err = io.ReadFull(reader, pkg)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}
