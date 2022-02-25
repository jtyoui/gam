// Package base
// @Time  : 2022/2/12 下午3:20
// @Author: Jtyoui@qq.com
// @note  : 处理一些通用的异常
package base

import (
	"errors"
	"github.com/jtyoui/gam/tool"
	"reflect"
)

// ParamError 根据参数判断是否异常
func ParamError(param string, typ reflect.Type) (r reflect.Value, err error) {
	r, err = tool.StringToAny(param, typ)
	if err != nil {
		err = errors.New("参数" + param + "绑定失败:" + err.Error())
		return
	}
	return
}
