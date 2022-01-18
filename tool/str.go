// Package tool
// @Time  : 2022/1/17 上午9:43
// @Author: Jtyoui@qq.com
// @note  : 关于字符串的处理工具类
package tool

import (
	"path/filepath"
	"reflect"
	"strings"
)

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

// GetFilePathByReflect 根据反射来获取文件地址路径
func GetFilePathByReflect(t reflect.Type) string {
	tf := t.Elem()
	pkgPath := tf.PkgPath()
	path, err := filepath.Abs("./" + pkgPath)
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(path + ".go")
	return filePath
}
