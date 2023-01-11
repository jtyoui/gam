// Package gam
// @Time  : 2021/10/21 上午10:49
// @Author: Jtyoui@qq.com
// @note  : 自动扫描路由
package gam

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam/dao"
	"github.com/jtyoui/gam/router"
	"github.com/jtyoui/gam/router/name"
	"github.com/jtyoui/gam/tool"
	"github.com/jtyoui/gam/web"
	"reflect"
	"strings"
)

// ginRouter 使用说明
type ginRouter struct {
	rg *gin.RouterGroup
	fs *embed.FS
	name.ChangeName
}

// NewGinFsRouter 初始化,如果是编译成二进制文件，那么需要使用该方法
// fs 需要引入dao层
func NewGinFsRouter(g *gin.RouterGroup, fs *embed.FS) *ginRouter {
	return &ginRouter{
		rg: g,
		fs: fs,
	}
}

// NewGinRouter 初始化
func NewGinRouter(g *gin.RouterGroup) *ginRouter {
	return NewGinFsRouter(g, nil)
}

// findName 获取该结构体的定义名称，如果结构体没有实现IGroupName接口，那么默认
// 路由名称是该结构体名字的小写，否则是定义函数的返回值
func (g *ginRouter) setName(v reflect.Value, path string, f router.NameFunc) (value string) {
	value, ok := f.GetName(v, path)
	if !ok {
		value = g.Change(path)
	}
	return
}

// 设置单一的路由
func (g *ginRouter) setRouter(route interface{}) {
	v := tool.ReflectByValue(route) // 初始化的结构体
	t := v.Type()                   // 转为类型

	// 获取到分组的名字,也就是结构体名字
	r := t.Elem().Name()
	groupName := g.setName(v, r, router.GroupNameFunc)

	// 将结构体名字加入路由分组
	newRouter := g.rg.Group(groupName)

	// 扫描go文件获取变量名称
	scanner := tool.NewGoFileScanner(g.fs)
	filePath := scanner.GetFilePathByReflect(t)
	err := scanner.ParseFile(filePath)
	if err != nil {
		panic(err)
	}
	methods := scanner.GetMethods(r)

	for _, method := range methods {
		// 获取结构体的函数名称
		methodName := method.MethodName

		// 根据正则表达式来判断函数头是否满足,不满足立即跳出
		if prefix := web.IsHttpProtocol(methodName); prefix != web.Nil {
			hrm := prefix.String()            // 获取字符串
			fun := v.MethodByName(methodName) // 根据名称来获取函数对象

			// 回调gin来获取上下午对象
			bind := dao.NewBindHandler(fun, method.Params)

			// 绑定路由参数
			routerFun := bind.BindParams(prefix)

			// 获取新的路由对象：该对象是用结构体对象生成
			value := tool.ReflectByValue(newRouter)

			// prefix满足Post|Get|Delete|Put其中之一。根据prefix来生成对应的router对象
			// 比如：当prefix为Get的时候，生成： gin.Group.GET
			method := value.MethodByName(strings.ToUpper(hrm))

			// 将函数头的名称去掉正则表达式里面的部分：例如函数头GetName去掉之后变成Name，并且转为小写
			// 并保存为路由地址
			realRouter := strings.TrimPrefix(methodName, hrm)

			// 进行改变命名规则
			path := g.setName(v, realRouter, router.ApiNameFunc)

			// 输入：method对象也就是路由函数输入需要两个参数，一个是路由地址，另一个是执行路由的函数
			inputs := []reflect.Value{
				reflect.ValueOf(path),      // 路由地址 比如：[GIN-debug] GET  /user/name
				reflect.ValueOf(routerFun), // 执行的具体函数对象
			}

			// 回调路由对象
			method.Call(inputs)
		}
	}
}

// AutoRouter 自动注册路由
func (g *ginRouter) AutoRouter(routers ...interface{}) {
	// 扫描只要实现了该方法的接口
	for _, r := range routers {
		g.setRouter(r)
	}
}
