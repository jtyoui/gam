// Package web
// @Time  : 2022/1/18 上午10:47
// @Author: Jtyoui@qq.com
// @note  : HTTP request methods
package web

import "strings"

type HRM uint // HTTP request methods
const (
	GET HRM = iota
	POST
	DELETE
	PUT
	Nil
)

func GetHRM() map[string]HRM {
	return map[string]HRM{
		"Get":    GET,
		"Delete": DELETE,
		"Post":   POST,
		"Put":    PUT,
		"Nil":    Nil,
	}
}

func (h HRM) String() string {
	for key, value := range GetHRM() {
		if value == h {
			return key
		}
	}
	return ""
}

// IsHttpProtocol 判断是否是http请求协议
func IsHttpProtocol(str string) HRM {
	for key, value := range GetHRM() {
		if strings.HasPrefix(str, key) {
			return value
		}
	}
	return Nil
}
