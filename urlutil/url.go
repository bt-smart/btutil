package urlutil

import (
	"errors"
	"strings"
)

// MatchesPattern 检查 URL 是否与模式匹配，支持单层通配符 "+" 和多层通配符 "#"
func MatchesPattern(url, pattern string) bool {
	// 去除前缀和后缀的斜杠
	url = strings.Trim(url, "/")
	pattern = strings.Trim(pattern, "/")

	urlParts := strings.Split(url, "/")
	patternParts := strings.Split(pattern, "/")

	for i, part := range patternParts {
		if part == "#" {
			return true
		}
		if part == "+" {
			if i >= len(urlParts) {
				return false
			}
			continue
		}
		if i >= len(urlParts) || part != urlParts[i] {
			return false
		}
	}

	return len(urlParts) == len(patternParts)
}

func ValidatePatterns(patterns []string) error {
	for _, p := range patterns {
		err := validatePattern(p)
		if err != nil {
			return err
		}
	}
	return nil
}

// 检查 pattern 是否符合规则
// 规则:
// 1. 开头必须是 /
// 2. 结尾不能有 /
// 3. 不能存在连续的 /
// 4. 单层通配符 + 必须占据整个层级
// 5. 多层通配符 # 必须占据整个层级且是最后一个字符
func validatePattern(pattern string) error {
	// 1. 开头必须是/
	if !strings.HasPrefix(pattern, "/") {
		return errors.New("pattern must start with '/'")
	}

	// 2. 结尾不能有/
	if strings.HasSuffix(pattern, "/") {
		return errors.New("pattern must not end with '/'")
	}

	// 3. 不能存在连续的/
	if strings.Contains(pattern, "//") {
		return errors.New("pattern must not contain '//'")
	}

	// 分割 pattern 成多个部分
	parts := strings.Split(pattern, "/")

	// 4. 校验单层通配符和多层通配符
	for i, part := range parts {
		// 单层通配符检查
		if part == "+" {
			continue
		}
		if strings.Contains(part, "+") {
			return errors.New("single-level wildcard '+' must occupy the entire level")
		}

		// 多层通配符检查
		if part == "#" {
			if i != len(parts)-1 {
				return errors.New("multi-level wildcard '#' must be at the end")
			}
			continue
		}
		if strings.Contains(part, "#") {
			return errors.New("multi-level wildcard '#' must occupy the entire level and be at the end")
		}
	}

	return nil
}

// IsValidHTTPSURL 校验字符串是否是合法的 HTTPS 链接，并且不携带参数
func IsValidHTTPSURL(url string) bool {
	if url == "" {
		return false
	}
	// 确保 URL 以 "https://" 开头
	if !strings.HasPrefix(url, "https://") {
		return false
	}

	// 检查是否包含参数（'?'）或锚点（'#'）
	if strings.Contains(url, "?") || strings.Contains(url, "#") {
		return false
	}

	// 确保 URL 不为空
	if len(url) <= len("https://") {
		return false
	}

	return true
}
