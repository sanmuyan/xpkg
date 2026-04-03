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

// ValidAesKey 验证 AES 密钥长度是否合法，合法的密钥长度为16(aes-128) 24(aes-192) 32(aes-256)
func ValidAesKey(key string) error {
	switch len(key) {
	case 16, 24, 32:
		return nil
	default:
		return errors.New("key length invalid")
	}
}

// EncryptCFB CFB 字符串加密
func EncryptCFB(plaintext string, key string) (string, error) {
	if len(plaintext) == 0 {
		return plaintext, nil
	}
	err := ValidAesKey(key)
	if err != nil {
		return plaintext, err
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
	if len(ciphertext) == 0 {
		return ciphertext, nil
	}
	err := ValidAesKey(key)
	if err != nil {
		return ciphertext, err
	}
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

// EncryptGCM GCM 字符串加密
func EncryptGCM(plaintext string, key string) (string, error) {
	if len(plaintext) == 0 {
		return plaintext, nil
	}
	err := ValidAesKey(key)
	if err != nil {
		return plaintext, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	result := append(nonce, ciphertext...)

	return hex.EncodeToString(result), nil
}

// DecryptGCM GCM 字符串解密
func DecryptGCM(ciphertext string, key string) (string, error) {
	gcmCiphertext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(gcmCiphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, nonceCiphertext := gcmCiphertext[:nonceSize], gcmCiphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, nonceCiphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// DecryptToStruct 将结构体中的加密字段转换为明文
func DecryptToStruct(x any, secretKey string, decryptFunc func(ciphertext string, key string) (string, error)) error {
	xv := reflect.ValueOf(x)
	if xv.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}
	ve := xv.Elem()
	if ve.Kind() == reflect.Struct {
		for i := 0; i < ve.NumField(); i++ {
			fv := ve.Field(i)
			switch fv.Kind() {
			case reflect.String:
				plaintext, err := decryptFunc(fv.String(), secretKey)
				if err != nil {
					return err
				}
				fv.SetString(plaintext)
			case reflect.Ptr:
				if fv.IsNil() {
					continue
				}
				switch fv.Elem().Kind() {
				case reflect.String:
					plaintext, err := decryptFunc(fv.Elem().String(), secretKey)
					if err != nil {
						return err
					}
					fv.Set(reflect.ValueOf(&plaintext))
				case reflect.Struct:
					if err := DecryptToStruct(fv.Interface(), secretKey, decryptFunc); err != nil {
						return err
					}
				default:
				}
			default:
			}
		}
	}
	return nil
}

// DecryptCFBToStruct 将结构体中的 CFB 加密字段转换为明文
func DecryptCFBToStruct(x any, secretKey string) error {
	return DecryptToStruct(x, secretKey, DecryptCFB)
}

// DecryptGCMToStruct 将结构体中的 GCM 加密字段转换为明文
func DecryptGCMToStruct(x any, secretKey string) error {
	return DecryptToStruct(x, secretKey, DecryptGCM)
}
