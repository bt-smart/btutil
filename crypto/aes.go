package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// DecryptAES256GCM
// 解密
// 使用 AEAD_AES_256_GCM 算法进行解密
func DecryptAES256GCM(aesKey, associatedData, nonce, ciphertext string) (plaintext string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	dataBytes, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return "", err
	}
	return string(dataBytes), nil
}

// EncryptAES256GCM
// 加密
// 使用 AEAD_AES_256_GCM 算法进行加密
func EncryptAES256GCM(aesKey, associatedData, nonce, plaintext string) (ciphertext string, err error) {
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	// 将 nonce 转换为字节切片
	nonceBytes := []byte(nonce)

	// 检查 nonce 的长度
	if len(nonceBytes) != gcm.NonceSize() {
		return "", fmt.Errorf("nonce 长度必须为 %d 字节", gcm.NonceSize())
	}

	// 加密
	ciphertextBytes := gcm.Seal(nil, nonceBytes, []byte(plaintext), []byte(associatedData))

	// 返回 Base64 编码的密文
	return base64.StdEncoding.EncodeToString(ciphertextBytes), nil
}
