// Package get
// @Time  : 2022/1/27 上午10:25
// @Author: Jtyoui@qq.com
// @note  : 获取get里面的数组类型
package get

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

// ArrayBind 基本get入参数组绑定
func ArrayBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	value := c.QueryArray(varName) // 获取数组参数的内容
	typ := t.String()              // 获取函数头的名字,根据名字来判断切片类型
	switch typ {
	case "[]int":
		data, e := intBind(value)
		if err = e; e != nil {
			return
		}
		r = reflect.ValueOf(&data)
	case "[]float64":
		data, e := float64Bind(value)
		if err = e; e != nil {
			return
		}
		r = reflect.ValueOf(&data)
	case "[]bool":
		data, e := boolBind(value)
		if err = e; e != nil {
			return
		}
		r = reflect.ValueOf(&data)
	case "[]string":
		r = reflect.ValueOf(&value)
	default: // todo 待优化：关于其它的切片类型
		panic("GET参数类型暂时不支持：" + typ)
	}
	return
}

// 绑定int切片
func intBind(values []string) (data []int, err error) {
	data = make([]int, len(values))
	for index, item := range values {
		value, e := strconv.Atoi(item)
		if err = e; e != nil {
			return
		}
		data[index] = value
	}
	return
}

// 绑定float64切片
func float64Bind(values []string) (data []float64, err error) {
	data = make([]float64, len(values))
	for index, item := range values {
		f64, e := strconv.ParseFloat(item, 64)
		if err = e; e != nil {
			return
		}
		data[index] = f64
	}
	return
}

// 绑定bool切片
func boolBind(values []string) (data []bool, err error) {
	data = make([]bool, len(values))
	for index, item := range values {
		value, e := strconv.ParseBool(item)
		if err = e; e != nil {
			return
		}
		data[index] = value
	}
	return
}
