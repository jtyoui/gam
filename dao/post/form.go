// Package post
// @Time  : 2022/2/12 下午2:53
// @Author: Jtyoui@qq.com
// @note  : post 表单格式
package post

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jtyoui/gam/dao/base"
	"reflect"
)

// FormBind 绑定表单格式
func FormBind(c *gin.Context, t reflect.Type, varName string) (r reflect.Value, err error) {
	value, ok := c.GetPostForm(varName)
	if !ok {
		err = errors.New("绑定的参数不存在，请检查变量：：" + varName)
		return
	}
	return base.ParamError(value, t)
}
