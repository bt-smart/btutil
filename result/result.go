package result

// Response 通用响应结构
type Response struct {
	// 错误码
	// 非0为有错误
	Code int `json:"code"`
	// 错误信息
	Msg string `json:"msg"`
	// 响应数据
	Data any `json:"data"`
}

func Ok() *Response {
	return &Response{Code: 0, Msg: "", Data: nil}
}

func Fail() *Response {
	return &Response{Code: 1, Msg: "", Data: nil}
}

func FailWithMsg(msg string) *Response {
	return &Response{Code: 1, Msg: msg, Data: nil}
}

func FailWithCodeAndMsg(code int, msg string) *Response {
	return &Response{Code: code, Msg: msg, Data: nil}
}

func Data(data any) *Response {
	return &Response{Code: 0, Msg: "", Data: data}
}
