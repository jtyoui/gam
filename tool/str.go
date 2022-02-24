// Package tool
// @Time  : 2022/1/17 上午9:43
// @Author: Jtyoui@qq.com
// @note  : 关于字符串的处理工具类
package tool

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// 对方项目必须是go mod形式。读取mod文件，返回自定义库
func (s *GoFileScanner) readGoMod() (mod string) {
	var rf io.Reader
	var err error
	if s.fs == nil {
		rf, err = os.Open("go.mod")
	} else {
		rf, err = s.fs.Open("go.mod")
	}

	if err != nil {
		panic("项目必须是go mod形式,没有找到mod文件")
	}
	br := bufio.NewReader(rf)
	for {
		a, _, err := br.ReadLine()
		line := strings.TrimSpace(string(a))
		if line != "" {
			number := strings.Split(line, " ")
			return number[1]
		}
		if err == io.EOF {
			break
		}
	}
	return
}

// GetFilePathByReflect 根据反射来获取文件地址路径
func (s *GoFileScanner) GetFilePathByReflect(t reflect.Type) string {
	tf := t.Elem()
	fileName := tf.Name()          // 获取文件名称
	fileName = Underline(fileName) // 将文件名称改成蛇形命名
	pkgPath := tf.PkgPath()
	if pkgPath == "main" {
		return fileName + ".go"
	}
	goMod := s.readGoMod()
	if goMod == "" {
		panic("请填写正确的go.mod文件")
	}
	suffix := strings.TrimPrefix(pkgPath, goMod)

	if s.fs == nil {
		address, err := filepath.Abs("./" + suffix)
		if err != nil {
			panic(err)
		}
		return filepath.Join(address, fileName+".go")
	}

	return filepath.Join(".", suffix, fileName+".go")
}

// ReplaceSepByFS 判断是否是非Linux系统，全部的路径符号需要将\转为/
// 在Fs中，所有的sep全是/
func ReplaceSepByFS(path string) string {
	root := strings.ReplaceAll(path, "\\", "/")
	return root
}
