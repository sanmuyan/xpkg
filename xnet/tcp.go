package xnet

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/sanmuyan/xpkg/xconstant"
)

func Encode(data []byte) ([]byte, error) {
	length := int32(len(data))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) ([]byte, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return nil, err
	}
	lengthBuf := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	if int32(reader.Buffered()) < length+4 {
		return nil, xconstant.BufferedTooSmallError
	}

	pkg := make([]byte, int(4+length))
	_, err = reader.Read(pkg)
	if err != nil {
		return nil, err
	}
	return pkg[4:], nil
}
