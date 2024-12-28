package crypto

import "testing"

func TestGetPasswordAndSalt(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "普通密码",
			password: "password123",
		},
		{
			name:     "中文密码",
			password: "测试密码123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试是否能正常获取密码和盐
			hashedPwd, salt, err := GetPasswordAndSalt(tt.password)
			if err != nil {
				t.Errorf("GetPasswordAndSalt() error = %v", err)
				return
			}

			// 验证返回值不为空
			if hashedPwd == "" || salt == "" {
				t.Error("GetPasswordAndSalt() 返回空值")
			}
		})
	}
}

func TestSha256PasswordWithSalt(t *testing.T) {
	password := "testpassword"
	salt := "testsalt"

	// 验证相同输入得到相同输出
	hash1 := Sha256PasswordWithSalt(password, salt)
	hash2 := Sha256PasswordWithSalt(password, salt)

	if hash1 != hash2 {
		t.Error("相同输入得到不同的哈希值")
	}
}

func TestSha256(t *testing.T) {
	input := "test string"

	// 验证相同输入得到相同输出
	hash1 := Sha256(input)
	hash2 := Sha256(input)

	if hash1 != hash2 {
		t.Error("相同输入得到不同的哈希值")
	}
}
