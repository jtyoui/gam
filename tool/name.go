// Package tool
// @Time  : 2022/1/17 上午9:18
// @Author: Jtyoui@qq.com
// @note  : 名字规范工具类
package tool

import (
	"fmt"
	"strings"
)

type NameTo struct {
	Name string
}

func NewNameTo(name string) *NameTo {
	return &NameTo{Name: name}
}

// Lower 全下写
func (n *NameTo) Lower() string {
	return Lower(n.Name)
}

// Lower 全下写
func Lower(name string) string {
	return strings.ToLower(name)
}

// Upper 全大写
func (n *NameTo) Upper() string {
	return Upper(n.Name)
}

// Upper 全大写
func Upper(name string) string {
	return strings.ToUpper(name)
}

// Title 按标题
func (n *NameTo) Title() string {
	return Title(n.Name)
}

// Title 按标题
func Title(name string) string {
	return strings.ToTitle(name)
}

// CamelCase 骆驼命名法
func (n *NameTo) CamelCase() string {
	return CamelCase(n.Name)
}

// CamelCase 骆驼命名法
func CamelCase(name string) string {
	if name[0] >= 65 && name[0] <= 90 {
		n := name[0] + 32
		name = fmt.Sprintf("%c%s", n, name[1:])
	}
	return name
}

// Pascal 帕斯卡命名法
func (n *NameTo) Pascal() string {
	return Pascal(n.Name)
}

// Pascal 帕斯卡命名法
func Pascal(name string) string {
	return name
}

// Underline 下划线
func (n *NameTo) Underline() string {
	return Underline(n.Name)
}

// Underline 下划线
func Underline(name string) string {
	value := make([]int32, 0, len(name))
	for _, i := range name {
		if 65 <= i && i <= 90 {
			value = append(value, 95, i+32) // 95的ascii是下划线
		} else {
			value = append(value, i)
		}
	}
	return strings.TrimPrefix(string(value), "_")
}
