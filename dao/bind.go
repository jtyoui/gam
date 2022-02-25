// Package dao
// @Time  : 2022/1/17 下午2:59
// @Author: Jtyoui@qq.com
// @note  : 绑定web http请求的参数到gin
package dao

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam/tool"
	"github.com/jtyoui/gam/web"
)

type bindHandler struct {
	F      reflect.Value
	Type   reflect.Type
	Params []string
}

// NewBindHandler 将参数绑定到反射中
func NewBindHandler(value reflect.Value, params []string) *bindHandler {
	bind := &bindHandler{
		F:      value,
		Type:   value.Type(),
		Params: params,
	}
	return bind
}

/*
	绑定参数,get支持URL传入，post支持json传入
	req 表示请求类型
	暂且不支持其它类型
*/
func (b *bindHandler) BindParams(hrm web.HRM) func(context *gin.Context) {
	num := b.Type.NumIn() // 获取函数的参数个数
	params := make([]reflect.Value, num)
	f := func(c *gin.Context) {
		for i := 0; i < num; i++ { // 遍历每一个具体的参数
			p := b.Type.In(i)       // 获取具体的参数信息
			rp := tool.RemovePtr(p) // 去掉类型指针，默认统一为非指针类型
			query := b.Params[i]    // 形参

			// 判断参数是不是*gin.Context
			if p.Kind() == reflect.Ptr && rp.Name() == "Context" {
				params[i] = reflect.ValueOf(c) // 如果是gin的上下文直接返回
				continue
			}
			method := newMethodBind(c, hrm, rp, query) // 获取方法的结构体
			value, err := method.valueByMethod()       // 根据具体的参数来获取值
			if err != nil {
				response(c, web.NewError(err)) // 绑定失败
				return
			}
			if p.Kind() != reflect.Ptr { // 判断绑定是不是指针类型
				value = value.Elem() // 如果是指针需要解引用，这点很重要，因为我默认所有的参数都是非指针类型
			}
			params[i] = value
		}
		values := b.F.Call(params) // 回调注册函数。获取返回值
		if values != nil {
			resultByData(c, values)
		}
	}
	return f
}

// 返回响应
func response(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// 更加函数返回的数据来返回结果
func resultByData(c *gin.Context, values []reflect.Value) {
	data := values[0].Interface() // 获取第一个参数
	length := len(values)         // 获取长度

	if length == 1 {
		response(c, data)
	} else if length == 2 {
		isErr := values[1].Interface()
		if isErr == nil { // 如果第二个参数是nil。返回第一个参数内容
			response(c, data)
		} else {
			if err, ok := isErr.(error); ok {
				response(c, web.NewError(err))
			} else {
				err := errors.New("第二个出参必须是error类型")
				response(c, web.NewError(err))
			}
		}
	} else {
		err := errors.New("出参的长度个数不能大于2")
		response(c, web.NewError(err))
	}
}
