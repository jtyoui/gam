// Package dao
// @Time  : 2022/1/18 上午9:49
// @Author: Jtyoui@qq.com
// @note  : 绑定get请求
package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
)

// GetBind 基本get入参绑定
func GetBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	value := c.Query(varName)
	result, err := StringToAny(value, t)
	if err != nil {
		err = errors.New("参数" + varName + "绑定失败")
		return
	}
	r = reflect.ValueOf(result)
	return
}
