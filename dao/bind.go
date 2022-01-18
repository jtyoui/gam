// Package dao
// @Time  : 2022/1/17 下午2:59
// @Author: Jtyoui@qq.com
// @note  : 绑定web http请求的参数到gin
package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/ginRoute/dao/post"
	"github.com/jtyoui/ginRoute/tool"
	"reflect"
)

type BindHandler struct {
	F      reflect.Value
	Type   reflect.Type
	Params []string
}

// 去掉指针
func removePtr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

/*
	绑定参数,get支持URL传入，post支持json传入
	req 表示请求类型
	暂且不支持其它类型
*/
func (b *BindHandler) BindParams(hrm tool.HRM) func(context *gin.Context) {
	num := b.Type.NumIn() // 获取函数的参数个数
	params := make([]reflect.Value, num)
	f := func(c *gin.Context) {
		for i := 0; i < num; i++ { // 遍历每一个具体的参数
			p := b.Type.In(i)  // 获取具体的参数信息
			rp := removePtr(p) // 去掉类型指针

			// 判断参数是不是*gin.Context
			if p.Kind() == reflect.Ptr && rp.Name() == "Context" {
				params[i] = reflect.ValueOf(c)
				continue
			}

			switch hrm {
			case tool.GET, tool.DELETE:
				params[i] = GetBind(c, rp, b.Params[i])
			case tool.POST, tool.PUT:
				params[i] = post.JsonBind(c, rp)
			}
		}
		b.F.Call(params)
	}
	return f
}
