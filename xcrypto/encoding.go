package xcrypto

import "encoding/base64"

func Base64Encode(src []byte) []byte {
	return base64.StdEncoding.AppendEncode(nil, src)
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.AppendDecode(nil, src)
}
