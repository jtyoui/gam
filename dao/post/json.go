// Package post
// @Time  : 2022/1/18 上午9:47
// @Author: Jtyoui@qq.com
// @note  : 判断json格式
package post

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
)

// JsonBind 基本json入参绑定
func JsonBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	param := reflect.New(t).Interface() // 获取类型参数，并不是参数名称
	err = c.ShouldBindJSON(param)
	if err != nil {
		if err == io.EOF {
			err = errors.New("绑定的结构体不存在，请检查变量：" + varName)
		}
		return
	}
	r = reflect.ValueOf(param)
	return
}
