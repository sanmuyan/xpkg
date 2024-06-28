package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"reflect"
	"unicode/utf8"
)

// EncryptCFB CFB 字符串加密，加密 key 长度必等于32
func EncryptCFB(plaintext string, key string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("key length invalid")
	}
	plaintextByte := []byte(plaintext)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return plaintext, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintextByte))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return plaintext, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextByte)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptCFB CFB 字符串解密
func DecryptCFB(ciphertext string, key string) (string, error) {
	plaintext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return ciphertext, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ciphertext, err
	}
	iv := plaintext[:aes.BlockSize]
	plaintext = plaintext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, plaintext)
	if !utf8.Valid(plaintext) {
		return ciphertext, errors.New("invalid ciphertext or key")
	}
	return string(plaintext), nil
}

// DecryptCFBToStruct 将结构体中的加密字段转换为明文
func DecryptCFBToStruct(x any, secretKey string) error {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}
	ve := v.Elem()
	if ve.Kind() == reflect.Struct {
		for i := 0; i < ve.NumField(); i++ {
			k := ve.Field(i)
			switch k.Type().Kind() {
			case reflect.String:
				plaintext, err := DecryptCFB(k.String(), secretKey)
				if err != nil {
					return err
				}
				v.Elem().Field(i).SetString(plaintext)
			case reflect.Ptr:
				if err := DecryptCFBToStruct(k.Interface(), secretKey); err != nil {
					return err
				}
			default:
				return errors.New("unsupported type")
			}
		}
	}
	return nil
}
