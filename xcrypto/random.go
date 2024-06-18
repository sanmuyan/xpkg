package xcrypto

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int, containsNumbers, containsUppercase, containsLowercase, containsSpecial bool) string {
	const digits = "0123456789"
	const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowercase = "abcdefghijklmnopqrstuvwxyz"
	const special = "!@#$%^&*"
	var charset strings.Builder
	var letter strings.Builder
	letter.WriteString(uppercase)
	letter.WriteString(lowercase)
	letterSet := letter.String()
	if containsNumbers {
		charset.WriteString(digits)
	}
	if containsUppercase {
		charset.WriteString(uppercase)
	}
	if containsLowercase {
		charset.WriteString(lowercase)
	}
	if containsSpecial {
		charset.WriteString(special)
	}
	if charset.Len() == 0 {
		charset.WriteString(digits)
		charset.WriteString(uppercase)
		charset.WriteString(lowercase)
		charset.WriteString(special)
	}
	characterSet := charset.String()
	result := make([]byte, length)
	for i := range result {
		if i == 0 {
			charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(letter.Len())))
			result[i] = letterSet[charIndex.Int64()]
			continue
		}
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(charset.Len())))
		s := characterSet[charIndex.Int64()]
		result[i] = s
	}
	return string(result)
}
