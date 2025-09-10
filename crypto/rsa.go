package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// ParseRSAPrivateKeyFromPEM 解析 PEM 格式的 RSA 私钥
func ParseRSAPrivateKeyFromPEM(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid RSA private key format")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
	}
	return privateKey, nil
}

// ParseRSAPublicKeyFromPEM 解析 PEM 格式的 RSA 公钥
func ParseRSAPublicKeyFromPEM(pemStr string) (*rsa.PublicKey, error) {
	blockPub, _ := pem.Decode([]byte(pemStr))
	if blockPub == nil || blockPub.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}
	pubKey, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("parsed key is not an RSA public key")
	}
	return rsaPubKey, nil
}
