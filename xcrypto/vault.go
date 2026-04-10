package xcrypto

import (
	"encoding/base64"
)

// 密码库加密逻辑
// 1. 使用 PBKDF2 生成 Master Key，通常使用用户密码和邮箱作为输入
// 2. 使用 HKDF 从 Master Key 扩展出专用的 Vault Key
// 3. 生成随机的 DEK，作为数据 AES-GCM 的加密密钥
// 4. 使用 DEK 加密数据
// 5. 使用 Vault Key 加密 DEK

type Vault struct {
	Salt   string `json:"salt"`
	EncDEK string `json:"enc_dek"`
	Data   string `json:"data"`
}

func genVaultKey(masterKey, salt []byte) ([]byte, error) {
	return GenExtendedKey(masterKey, salt, []byte("vault")) //
}

// CreateVault 创建 Vault
func CreateVault(masterKey, plaintext []byte) (vault *Vault, err error) {
	salt := GenSalt()
	vaultKey, err := genVaultKey(masterKey, salt)
	if err != nil {
		return
	}
	dek := GenDEK()
	encData, err := EncryptGCM(plaintext, dek)
	if err != nil {
		return
	}
	encDEK, err := EncryptGCM(dek, vaultKey)
	if err != nil {
		return
	}

	vault = &Vault{
		Salt:   base64.StdEncoding.EncodeToString(salt),
		EncDEK: base64.StdEncoding.EncodeToString(encDEK),
		Data:   base64.StdEncoding.EncodeToString(encData),
	}
	return
}

// DecryptVault 解密 Vault
func DecryptVault(v *Vault, masterKey []byte) (data []byte, err error) {
	salt, err := base64.StdEncoding.DecodeString(v.Salt)
	if err != nil {
		return
	}
	vaultKey, err := genVaultKey(masterKey, salt)
	if err != nil {
		return
	}
	encDEK, err := base64.StdEncoding.DecodeString(v.EncDEK)
	if err != nil {
		return
	}
	dek, err := DecryptGCM(encDEK, vaultKey)
	if err != nil {
		return
	}
	encData, err := base64.StdEncoding.DecodeString(v.Data)
	if err != nil {
		return
	}
	data, err = DecryptGCM(encData, dek)
	if err != nil {
		return
	}

	return
}

// UpdateVaultKey 更新 Vault Key
func UpdateVaultKey(v *Vault, oldMasterKey, newMasterKey []byte, isUpdateDEK bool) (err error) {
	salt, err := base64.StdEncoding.DecodeString(v.Salt)
	if err != nil {
		return
	}
	oldVaultKey, err := genVaultKey(oldMasterKey, salt)
	if err != nil {
		return
	}
	encDEK, err := base64.StdEncoding.DecodeString(v.EncDEK)
	if err != nil {
		return
	}
	dek, err := DecryptGCM(encDEK, oldVaultKey)
	if err != nil {
		return
	}
	if isUpdateDEK {
		data, err := base64.StdEncoding.DecodeString(v.Data)
		plaintext, err := DecryptGCM(data, dek)
		if err != nil {
			return err
		}
		dek = GenDEK()
		newEncData, _ := EncryptGCM(plaintext, dek)
		v.Data = base64.StdEncoding.EncodeToString(newEncData)
	}
	newVaultKey, err := genVaultKey(newMasterKey, salt)
	if err != nil {
		return err
	}
	newEncDEK, err := EncryptGCM(dek, newVaultKey)
	if err != nil {
		return err
	}
	v.EncDEK = base64.StdEncoding.EncodeToString(newEncDEK)

	return nil
}
