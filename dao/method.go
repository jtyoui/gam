// Package dao
// @Time  : 2022/1/17 下午2:59
// @Author: Jtyoui@qq.com
// @note  : 根据值去绑定一些方法，比如get、post等
package dao

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam/dao/get"
	"github.com/jtyoui/gam/dao/post"
	"github.com/jtyoui/gam/web"
)

type methodBind struct {
	c    *gin.Context
	h    web.HRM      // 请求协议
	typ  reflect.Type // 获取该方法的数据类型
	name string       // 参数名称
}

func newMethodBind(c *gin.Context, h web.HRM, t reflect.Type, n string) *methodBind {
	return &methodBind{c: c, h: h, typ: t, name: n}
}

func (m *methodBind) get() (r reflect.Value, err error) {
	switch m.typ.Kind() {
	case reflect.Slice:
		r, err = get.ArrayBind(m.c, m.typ, m.name) // 走数组绑定，数组遵守gin的绑定规则
	default:
		r, err = get.QueryBind(m.c, m.typ, m.name) // 走的get绑定
	}
	return
}

func (m *methodBind) post() (r reflect.Value, err error) {
	switch m.typ.Kind() {
	case reflect.Struct, reflect.Slice:
		r, err = post.JsonBind(m.c, m.typ, m.name) // 走的post绑定。现在默认只有json格式
	default:
		r, err = post.FormBind(m.c, m.typ, m.name)
	}
	return
}

// todo 有待优化：此处走的是get操作，但是不一定
func (m *methodBind) delete() (r reflect.Value, err error) {
	return m.get()
}

// todo 有待优化：此处走的是Post操作，但是不一定
func (m *methodBind) put() (r reflect.Value, err error) {
	return m.post()
}

// 根据具体地请求协议去调不同的方法
func (m *methodBind) valueByMethod() (r reflect.Value, err error) {
	switch m.h {
	case web.GET:
		r, err = m.get()
	case web.DELETE:
		r, err = m.delete()
	case web.POST:
		r, err = m.post()
	case web.PUT:
		r, err = m.put()
	}
	return
}
