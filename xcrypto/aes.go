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

func EncryptCFB(plaintext string, key string) (string, error) {
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

func DecryptCFBToStruct(x any, secretKey string) error {
	vPrt := reflect.ValueOf(x)
	if vPrt.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}
	v := reflect.ValueOf(vPrt.Elem().Interface())
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			k := v.Field(i)
			switch k.Type().Kind() {
			case reflect.String:
				plaintext, err := DecryptCFB(k.String(), secretKey)
				if err != nil {
					return err
				}
				vPrt.Elem().Field(i).SetString(plaintext)
			case reflect.Ptr:
				if err := DecryptCFBToStruct(k.Interface(), secretKey); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
