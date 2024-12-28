package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/bt-smart/util/strutil"
)

// GetPasswordAndSalt 获取密码和盐
func GetPasswordAndSalt(password string) (string, string, error) {
	// 生成32个字符的盐
	salt, err := strutil.GenerateRandomString(32, strutil.LowercaseCharset)
	if err != nil {
		return "", "", err
	}
	psw := Sha256PasswordWithSalt(password, salt)
	return psw, salt, nil
}

// Sha256PasswordWithSalt 对密码进行哈希处理
func Sha256PasswordWithSalt(password string, salt string) string {
	// 将密码和盐拼接起来
	saltedPassword := password + salt
	// 计算哈希值
	hash := sha256.Sum256([]byte(saltedPassword))
	// 返回十六进制表示的哈希值
	return hex.EncodeToString(hash[:])
}

func Sha256(str string) string {
	// 计算哈希值
	hash := sha256.Sum256([]byte(str))
	// 返回十六进制表示的哈希值
	return hex.EncodeToString(hash[:])
}
