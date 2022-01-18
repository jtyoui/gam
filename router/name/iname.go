// Package name
// @Time  : 2022/1/17 上午9:24
// @Author: Jtyoui@qq.com
// @note  : 定义路由名称的接口文件
package name

/*
	更改路由分组名称
	比如有一些这样的路由地址：
		/student/add
		/student/id
		/student/deleteById
	想把里面的分组student改成user
	/student/add -> /user/add
*/
type IGroupName interface {
	GroupName(name string) string
}

/*
	更改api名称规范
	比如有一些这样的路由地址：
		/student/add
		/student/id
		/student/deleteById
	想把路由里面的字符a全部转化为A,只需要实现改接口
	/student/add --> /student/Add

	注意：只会更改最后一层的路由
	比如：  /a/add --> /a/Add
*/
type IApiName interface {
	ApiName(name string) string
}

/*
	更改路由所有字符名称规范
	比如有一些这样的路由地址：
		/a/add
	更改里面所有的a为大写的A
		/a/add --> /A/Add
*/
type IAllName interface {
	AllName(name string) string
}
