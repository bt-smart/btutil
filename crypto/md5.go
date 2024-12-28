package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 计算md5值
func Md5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// CalculateFileMD5 计算文件的MD5值
func CalculateFileMD5(reader io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
