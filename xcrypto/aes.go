package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"reflect"
)

// ValidAesKey 验证 AES 密钥长度是否合法，合法的密钥长度为16(aes-128) 24(aes-192) 32(aes-256)
func ValidAesKey(key []byte) error {
	switch len(key) {
	case 16, 24, 32:
		return nil
	default:
		return errors.New("key length invalid")
	}
}

// EncryptCFB CFB 加密
func EncryptCFB(plaintext, key []byte, opts ...EncryptOption) (ciphertext []byte, err error) {
	err = ValidAesKey(key)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	c := applyEncryptOption(opts...)
	ciphertext = encryptEncode(ciphertext, c)
	return
}

// DecryptCFB CFB 解密
func DecryptCFB(ciphertext, key []byte, opts ...EncryptOption) (plaintext []byte, err error) {
	c := applyEncryptOption(opts...)
	ciphertext, err = decryptDecode(ciphertext, c)
	if err != nil {
		return
	}
	err = ValidAesKey(key)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	plaintext = append(plaintext, ciphertext...)
	iv := plaintext[:aes.BlockSize]
	plaintext = plaintext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, plaintext)
	return
}

// EncryptGCM GCM 字符串加密
func EncryptGCM(plaintext, key []byte, opts ...EncryptOption) (ciphertext []byte, err error) {
	if len(plaintext) == 0 {
		return nil, nil
	}
	err = ValidAesKey(key)
	if err != nil {
		return
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return
	}

	ciphertextPayload := gcm.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(nonce, ciphertextPayload...)
	c := applyEncryptOption(opts...)
	ciphertext = encryptEncode(ciphertext, c)
	return
}

// DecryptGCM GCM 字符串解密
func DecryptGCM(ciphertext, key []byte, opts ...EncryptOption) (plaintext []byte, err error) {
	c := applyEncryptOption(opts...)
	ciphertext, err = decryptDecode(ciphertext, c)
	if err != nil {
		return
	}
	err = ValidAesKey(key)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertextPayload := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err = gcm.Open(nil, nonce, ciphertextPayload, nil)
	if err != nil {
		return
	}
	return
}

type DecryptFunc func(ciphertext, key []byte, opts ...EncryptOption) (plaintext []byte, err error)

// DecryptToStruct 将结构体中的加密字段转换为明文
func DecryptToStruct(x any, key []byte, decryptFunc DecryptFunc, inDecoder EncryptOption) error {
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
				plaintext, err := decryptFunc([]byte(fv.String()), key, inDecoder)
				if err != nil {
					return err
				}
				fv.SetString(string(plaintext))
			case reflect.Ptr:
				if fv.IsNil() {
					continue
				}
				switch fv.Elem().Kind() {
				case reflect.String:
					plaintext, err := decryptFunc([]byte(fv.Elem().String()), key, inDecoder)
					if err != nil {
						return err
					}
					fv.Set(reflect.ValueOf(new(string(plaintext))))
				case reflect.Struct:
					if err := DecryptToStruct(fv.Interface(), key, decryptFunc, inDecoder); err != nil {
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
