package crypto

import (
	"testing"
)

// TestMd5 测试Md5
func TestMd5(t *testing.T) {
	md5 := Md5([]byte("hello"))
	if md5 != "5d41402abc4b2a76b9719d911017c592" {
		t.Errorf("Md5 failed, expected %s, got %s", "5d41402abc4b2a76b9719d911017c592", md5)
	}
}
