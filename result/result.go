package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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

type PageData[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

// --------- 构造器 ---------

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

func Page[T any](list []T, total int64) *Response {
	return &Response{Code: 0, Msg: "", Data: PageData[T]{List: list, Total: total}}
}

func Bool(ok bool) *Response {
	if ok {
		return Ok()
	}
	return Fail()
}

func BoolWithMsg(ok bool, msg string) *Response {
	if ok {
		return Ok()
	}
	return FailWithMsg(msg)
}

// --------- Gin 简化方法 (GinXxx) ---------

func GinOk(c *gin.Context) {
	c.JSON(http.StatusOK, Ok())
}

func GinFail(c *gin.Context) {
	c.JSON(http.StatusOK, Fail())
}

func GinFailMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, FailWithMsg(msg))
}

func GinData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Data(data))
}

func GinPage[T any](c *gin.Context, list []T, total int64) {
	c.JSON(http.StatusOK, Page(list, total))
}

func GinBool(c *gin.Context, ok bool) {
	c.JSON(http.StatusOK, Bool(ok))
}

func GinBoolMsg(c *gin.Context, ok bool, msg string) {
	c.JSON(http.StatusOK, BoolWithMsg(ok, msg))
}
