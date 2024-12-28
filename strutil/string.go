package strutil

import (
	"crypto/rand"
	"math/big"
	"unicode"
)

const AllLettersAndDigits = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 大小写字母 + 数字
const LowercaseCharset = "abcdefghijklmnopqrstuvwxyz0123456789"                              // 小写字母 + 数字

// ValidateString 校验字符串只能包含小写字母或数字，并且长度在[minLength, maxLength]范围内
func ValidateString(s string, minLength, maxLength int) bool {
	length := len(s)
	if length < minLength || length > maxLength {
		return false
	}
	for _, ch := range s {
		if !unicode.IsLower(ch) && !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int, charset string) (string, error) {
	result := make([]byte, length)
	maxIdx := big.NewInt(int64(len(charset)))
	for i := range result {
		idx, err := rand.Int(rand.Reader, maxIdx)
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}
	return string(result), nil
}
