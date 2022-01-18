// Package name
// @Time  : 2021/11/16 下午4:41
// @Author: Jtyoui@qq.com
// @note  : 根据类型自动调用函数
package name

// ChangeName 命名规范转换器
type ChangeName struct {
	NamingConvention NamingConvention // 命名规范类型
	CustomFunc       IApiName         // 自定义命名规范函数
}

/*
	字符串根据命名规则进行转换
	先执行自定义命名
*/
func (c *ChangeName) Change(name string) string {
	if c.CustomFunc != nil {
		return c.CustomFunc.ApiName(name)
	}
	return c.NamingConvention.RName(name)
}
