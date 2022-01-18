// Package name
// @Time  : 2021/11/16 下午3:37
// @Author: Jtyoui@qq.com
// @note  : 命名规范类型
package name

import "github.com/jtyoui/ginRoute/tool"

// NamingConvention 路由命名规范
type NamingConvention uint

// 命名规范
const (
	CamelCase NamingConvention = iota // 骆驼命名法   = 小骆驼命名法  https://baike.baidu.com/item/%E9%AA%86%E9%A9%BC%E5%91%BD%E5%90%8D%E6%B3%95
	Upper                             // 全大写
	Title                             // 按标题
	Lower                             // 全下写
	Pascal                            // 帕斯卡命名法 = 大骆驼命名法  https://baike.baidu.com/item/%E5%B8%95%E6%96%AF%E5%8D%A1%E5%91%BD%E5%90%8D%E6%B3%95/9464494
	Underline                         // 下划线
)

// RName 根据类型执行不同的函数
func (n NamingConvention) RName(name string) string {
	t := tool.NewNameTo(name)
	switch n {
	case Lower:
		return t.Lower()
	case Upper:
		return t.Upper()
	case Title:
		return t.Title()
	case CamelCase:
		return t.CamelCase()
	case Pascal:
		return t.Pascal()
	case Underline:
		return t.Underline()
	default:
		return name
	}
}
