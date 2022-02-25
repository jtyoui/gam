// Package router
// @Time  : 2022/1/17 下午1:51
// @Author: Jtyoui@qq.com
// @note  : 关于名字更改具体调用的函数
package router

import (
	"github.com/jtyoui/gam/router/name"
	"reflect"
)

// NameFunc 更换路由段枚举
type NameFunc uint

const (
	ApiNameFunc   NameFunc = iota // api
	GroupNameFunc                 // 分组
	AllNameFunc                   // 全部
)

// GetName 根据自定义名称接口来对应的调用函数
func (n NameFunc) GetName(v reflect.Value, r string) (path string, flag bool) {
	obj := v.Interface()
	switch n {
	case ApiNameFunc:
		if m, ok := obj.(name.IApiName); ok {
			path = m.ApiName(r)
			flag = true
		}
	case GroupNameFunc:
		if m, ok := obj.(name.IGroupName); ok {
			path = m.GroupName(r)
			flag = true
		}
	case AllNameFunc:
		if m, ok := obj.(name.IAllName); ok {
			path = m.AllName(r)
			flag = true
		}
	}

	if n != AllNameFunc && !flag {
		return AllNameFunc.GetName(v, r)
	}
	return
}
