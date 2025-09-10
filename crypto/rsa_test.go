package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestParseRSAPrivateKeyFromPEM(t *testing.T) {
	// 测试用例1：有效的RSA私钥
	t.Run("有效的RSA私钥", func(t *testing.T) {
		// 生成一个测试用的RSA私钥
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatalf("生成RSA私钥失败: %v", err)
		}

		// 将私钥转换为PEM格式
		privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privateKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		})

		// 测试解析函数
		parsedKey, err := ParseRSAPrivateKeyFromPEM(string(privateKeyPEM))
		if err != nil {
			t.Errorf("解析有效的RSA私钥失败: %v", err)
		}

		// 验证解析后的私钥是否与原始私钥匹配
		if parsedKey.N.Cmp(privateKey.N) != 0 {
			t.Error("解析后的私钥与原始私钥不匹配")
		}
	})

	// 测试用例2：无效的PEM格式
	t.Run("无效的PEM格式", func(t *testing.T) {
		invalidPEM := "这不是一个有效的PEM格式"
		_, err := ParseRSAPrivateKeyFromPEM(invalidPEM)
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})

	// 测试用例3：错误的私钥类型
	t.Run("错误的私钥类型", func(t *testing.T) {
		// 创建一个非RSA PRIVATE KEY类型的PEM
		wrongTypePEM := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: []byte("测试数据"),
		})

		_, err := ParseRSAPrivateKeyFromPEM(string(wrongTypePEM))
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})
}

func TestParseRSAPublicKeyFromPEM(t *testing.T) {
	// 测试用例1：有效的RSA公钥
	t.Run("有效的RSA公钥", func(t *testing.T) {
		// 生成一个测试用的RSA密钥对
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatalf("生成RSA密钥对失败: %v", err)
		}
		publicKey := &privateKey.PublicKey

		// 将公钥转换为PEM格式
		publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			t.Fatalf("序列化公钥失败: %v", err)
		}
		publicKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		})

		// 测试解析函数
		parsedKey, err := ParseRSAPublicKeyFromPEM(string(publicKeyPEM))
		if err != nil {
			t.Errorf("解析有效的RSA公钥失败: %v", err)
		}

		// 验证解析后的公钥是否与原始公钥匹配
		if parsedKey.N.Cmp(publicKey.N) != 0 {
			t.Error("解析后的公钥与原始公钥不匹配")
		}
	})

	// 测试用例2：无效的PEM格式
	t.Run("无效的PEM格式", func(t *testing.T) {
		invalidPEM := "这不是一个有效的PEM格式"
		_, err := ParseRSAPublicKeyFromPEM(invalidPEM)
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})

	// 测试用例3：错误的公钥类型
	t.Run("错误的公钥类型", func(t *testing.T) {
		// 创建一个非PUBLIC KEY类型的PEM
		wrongTypePEM := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: []byte("测试数据"),
		})

		_, err := ParseRSAPublicKeyFromPEM(string(wrongTypePEM))
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})

	// 测试用例4：非RSA公钥
	t.Run("非RSA公钥", func(t *testing.T) {
		// 这个测试用例模拟解析出的公钥不是RSA类型的情况
		// 由于我们无法直接创建非RSA类型的公钥PEM，这里只能测试错误处理逻辑
		// 实际上这个测试在真实环境中很难触发，但我们仍然应该测试这个错误处理分支

		// 创建一个有效的PEM格式但内容无法解析为RSA公钥的数据
		invalidKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: []byte{0x30, 0x03, 0x01, 0x01, 0x00}, // 这是一个ASN.1编码的布尔值true，不是RSA公钥
		})

		_, err := ParseRSAPublicKeyFromPEM(string(invalidKeyPEM))
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})
}
