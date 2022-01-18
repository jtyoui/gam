// Package dao
// @Time  : 2022/1/17 下午5:51
// @Author: Jtyoui@qq.com
// @note  : 类型转化
package dao

import (
	"reflect"
	"strconv"
)

// StringToAny 将字符串转为基本类型
func StringToAny(str string, t reflect.Type) (v interface{}, err error) {
	type_ := t.Kind()
	switch type_ {
	case reflect.Bool:
		v, err = strconv.ParseBool(str)
	case reflect.Int:
		v, err = strconv.Atoi(str)
	case reflect.Int8:
		v, err = strconv.ParseInt(str, 10, 8)
	case reflect.Int16:
		v, err = strconv.ParseInt(str, 10, 8)
	case reflect.Int32:
		v, err = strconv.ParseInt(str, 10, 32)
	case reflect.Int64:
		v, err = strconv.ParseInt(str, 10, 64)
	case reflect.Uint:
		v, err = strconv.ParseUint(str, 10, 32)
	case reflect.Uint8:
		v, err = strconv.ParseInt(str, 10, 8)
	case reflect.Uint16:
		v, err = strconv.ParseInt(str, 10, 16)
	case reflect.Uint32:
		v, err = strconv.ParseInt(str, 10, 32)
	case reflect.Uint64:
		v, err = strconv.ParseInt(str, 10, 64)
	case reflect.Float32:
		v, err = strconv.ParseFloat(str, 32)
	case reflect.Float64:
		v, err = strconv.ParseFloat(str, 64)
	case reflect.String:
		return str, nil
	default:
		panic("不支持的类型" + type_.String())
	}
	return
}
