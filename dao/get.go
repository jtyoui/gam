// Package dao
// @Time  : 2022/1/18 上午9:49
// @Author: Jtyoui@qq.com
// @note  : 绑定get请求
package dao

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

// GetBind 基本get入参绑定
func GetBind(c *gin.Context, t reflect.Type, varName string) reflect.Value {
	value := c.Query(varName)
	result, err := StringToAny(value, t)
	if err != nil {
		panic(err)
	}
	return reflect.ValueOf(result)
}
