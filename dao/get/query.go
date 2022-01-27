// Package get
// @Time  : 2022/1/18 上午9:49
// @Author: Jtyoui@qq.com
// @note  : 绑定get请求
package get

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jtyoui/ginRoute/tool"
)

// QueryBind 基本get入参绑定
func QueryBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	value := c.Query(varName)
	r, err = tool.StringToAny(value, t)
	if err != nil {
		err = errors.New("参数" + varName + "绑定失败:" + err.Error())
		return
	}
	return
}
