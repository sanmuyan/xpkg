package xcrypto

import (
	"encoding/base64"
	"encoding/hex"
)

func Base64Encode(src []byte) []byte {
	return base64.StdEncoding.AppendEncode(nil, src)
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.AppendDecode(nil, src)
}

func HexEncode(src []byte) []byte {
	return hex.AppendEncode(nil, src)
}

func HexDecode(src []byte) ([]byte, error) {
	return hex.AppendDecode(nil, src)
}
