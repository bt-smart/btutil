package urlutil

import (
	"testing"
)

// 测试 matchesPattern 函数
func TestMatchesPattern(t *testing.T) {
	tests := []struct {
		url     string
		pattern string
		match   bool
	}{
		{"/user/1/add", "/user/+/add", true},
		{"/user/1/add", "/user/#", true},
		{"/user/1/add", "/user/1/#", true},
		{"/user/1/add", "/user/1/+", true},
		{"/user/1/add", "/user/1/2", false},
		{"/user/1/add", "/user/+", false},
		{"/user/1/add", "/user/+++/add", false},
		{"/user/1/add", "/user//add", false},
		{"/user/1/add/", "/user/1/add", true},
	}

	for _, test := range tests {
		if result := MatchesPattern(test.url, test.pattern); result != test.match {
			t.Errorf("matchesPattern(%s, %s) = %v; want %v", test.url, test.pattern, result, test.match)
		}
	}
}

// 测试 validatePattern 函数
func TestValidatePattern(t *testing.T) {
	tests := []struct {
		pattern string
		valid   bool
	}{
		{"/sensor/+", true},
		{"/sensor/#", true},
		{"/sensor/bedroom#", false},
		{"/sensor/#/temperature", false},
		{"/sensor/+/temperature/", false},
		{"/sensor/+/temperature", true},
		{"/sensor//temperature", false},
		{"/sensor/bedroom/+/temperature", true},
		{"/sensor/bedroom/#", true},
		{"/sensor/bedroom/+/temperature/", false},
		{"/sensor/bedroom/++/temperature", false},
	}

	for _, test := range tests {
		err := validatePattern(test.pattern)
		if (err == nil) != test.valid {
			t.Errorf("validatePattern(%s) = %v; want %v", test.pattern, err == nil, test.valid)
		}
	}
}
