// Package tool
// @Time  : 2022/1/17 上午9:43
// @Author: Jtyoui@qq.com
// @note  : 关于字符串的处理工具类
package tool

import (
	"bufio"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// 对方项目必须是go mod形式。读取mod文件，返回自定义库
func readGoMod() (mod string) {
	rf, err := os.Open("go.mod")
	if err != nil {
		panic("项目必须是go mod形式,没有找到mod文件")
	}
	defer func(rf *os.File) { _ = rf.Close() }(rf)

	br := bufio.NewReader(rf)
	for {
		a, _, _ := br.ReadLine()
		line := strings.TrimSpace(string(a))
		if line != "" {
			number := strings.Split(line, " ")
			return number[1]
		}
	}
	return
}

// GetFilePathByReflect 根据反射来获取文件地址路径
func GetFilePathByReflect(t reflect.Type) string {
	tf := t.Elem()
	pkgPath := tf.PkgPath()
	if pkgPath == "main" {
		return tf.Name() + ".go"
	}
	goMod := readGoMod()
	if goMod == "" {
		panic("请填写正确的go.mod文件")
	}
	suffix := strings.TrimPrefix(pkgPath, goMod)
	address, err := filepath.Abs("./" + suffix)
	if err != nil {
		panic(err)
	}
	p := filepath.Join(address, tf.Name()+".go")
	return p
}
