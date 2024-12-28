package crypto

import (
	"testing"
)

func TestAES256GCM(t *testing.T) {
	// 测试数据
	tests := []struct {
		name           string
		aesKey         string
		associatedData string
		nonce          string
		plaintext      string
	}{
		{
			name:           "测试用例1",
			aesKey:         "0123456789abcdef0123456789abcdef", // 32字节密钥
			associatedData: "associated-data",
			nonce:          "123456789012", // 12字节nonce
			plaintext:      "Hello, World!",
		},
		{
			name:           "测试用例2-中文",
			aesKey:         "0123456789abcdef0123456789abcdef",
			associatedData: "关联数据",
			nonce:          "123456789012",
			plaintext:      "你好，世界！",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. 首先测试加密
			ciphertext, err := EncryptAES256GCM(tt.aesKey, tt.associatedData, tt.nonce, tt.plaintext)
			if err != nil {
				t.Errorf("加密失败: %v", err)
				return
			}

			// 2. 然后测试解密
			decrypted, err := DecryptAES256GCM(tt.aesKey, tt.associatedData, tt.nonce, ciphertext)
			if err != nil {
				t.Errorf("解密失败: %v", err)
				return
			}

			// 3. 验证解密后的结果是否与原文相同
			if decrypted != tt.plaintext {
				t.Errorf("解密结果与原文不匹配\n期望: %s\n实际: %s", tt.plaintext, decrypted)
			}
		})
	}
}

func TestAES256GCM_InvalidInput(t *testing.T) {
	// 测试错误情况
	t.Run("无效的密钥长度", func(t *testing.T) {
		_, err := EncryptAES256GCM("短密钥", "data", "123456789012", "plaintext")
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})

	t.Run("无效的nonce长度", func(t *testing.T) {
		_, err := EncryptAES256GCM(
			"0123456789abcdef0123456789abcdef",
			"data",
			"短nonce",
			"plaintext",
		)
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})

	t.Run("无效的密文", func(t *testing.T) {
		_, err := DecryptAES256GCM(
			"0123456789abcdef0123456789abcdef",
			"data",
			"123456789012",
			"invalid-base64",
		)
		if err == nil {
			t.Error("期望获得错误，但没有")
		}
	})
}
