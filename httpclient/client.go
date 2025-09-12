package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Response 封装响应
type Response struct {
	StatusCode int
	Body       []byte
}

// Client HTTP 客户端
type Client struct {
	httpClient *http.Client
}

// New 创建一个带超时的 http 客户端
func New(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
	}
}

// Get 发送 GET 请求
func (c *Client) Get(reqURL string, headers map[string]string) (*Response, error) {
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

// PostJSON 发送 POST JSON 请求
func (c *Client) PostJSON(reqURL string, data any, headers map[string]string) (*Response, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return handleResponse(resp)
}

// 统一处理响应
func handleResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
	}, nil
}
