package result

import (
	"encoding/json"
	"testing"
)

// 定义一个测试用结构体，模拟业务结构
type TestUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func TestDataWithPageData(t *testing.T) {
	// 模拟分页数据
	users := []TestUser{
		{ID: 1, Username: "alice"},
		{ID: 2, Username: "bob"},
	}
	page := PageData[TestUser]{
		List:  users,
		Total: 2,
	}

	// 封装响应
	resp := Data(page)

	// 转成 JSON
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	t.Logf("Response JSON: %s", jsonBytes)

	// 断言结果
	expected := `{"code":0,"msg":"","data":{"list":[{"id":1,"username":"alice"},{"id":2,"username":"bob"}],"total":2}}`
	if string(jsonBytes) != expected {
		t.Errorf("Unexpected JSON.\nGot:  %s\nWant: %s", jsonBytes, expected)
	}
}
