// Package get
// @Time  : 2022/1/18 上午9:49
// @Author: Jtyoui@qq.com
// @note  : 绑定get请求
package get

import (
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam/dao/base"
	"reflect"
)

// QueryBind 基本get入参绑定
func QueryBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	value := c.Query(varName)
	return base.ParamError(value, t)
}
