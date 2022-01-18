// Package web
// @Time  : 2022/1/18 下午3:29
// @Author: Jtyoui@qq.com
// @note  : 错误处理
package web

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// NewError 返回绑定异常
func NewError(err error) Error {
	return Error{
		Code: 6500,
		Msg:  err.Error(),
	}
}
