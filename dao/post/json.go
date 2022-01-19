// Package post
// @Time  : 2022/1/18 上午9:47
// @Author: Jtyoui@qq.com
// @note  : 判断json格式
package post

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

// JsonBind 基本json入参绑定
func JsonBind(c *gin.Context, t reflect.Type) (r reflect.Value, err error) {
	param := reflect.New(t).Interface()
	err = c.ShouldBindJSON(param)
	if err != nil {
		return
	}
	r = reflect.ValueOf(param)
	return
}
