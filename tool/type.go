// Package tool
// @Time  : 2022/1/17 下午5:51
// @Author: Jtyoui@qq.com
// @note  : 类型转化
package tool

import (
	"reflect"
	"strconv"
)

// StringToAny 将字符串转为基本类型
func StringToAny(str string, t reflect.Type) (r reflect.Value, err error) {
	type_ := t.Kind()
	switch type_ {
	case reflect.Bool:
		b, e := strconv.ParseBool(str)
		r = reflect.ValueOf(&b)
		err = e
	case reflect.Int:
		i, e := strconv.Atoi(str)
		r = reflect.ValueOf(&i)
		err = e
	case reflect.Int8, reflect.Uint8:
		i8, e := strconv.ParseInt(str, 10, 8)
		r = reflect.ValueOf(&i8)
		err = e
	case reflect.Int16, reflect.Uint16:
		i16, e := strconv.ParseInt(str, 10, 16)
		r = reflect.ValueOf(&i16)
		err = e
	case reflect.Int32, reflect.Uint, reflect.Uint32:
		i32, e := strconv.ParseInt(str, 10, 32)
		r = reflect.ValueOf(&i32)
		err = e
	case reflect.Int64, reflect.Uint64:
		i64, e := strconv.ParseInt(str, 10, 64)
		r = reflect.ValueOf(&i64)
		err = e
	case reflect.Float32:
		f32, e := strconv.ParseFloat(str, 32)
		r = reflect.ValueOf(&f32)
		err = e
	case reflect.Float64:
		f64, e := strconv.ParseFloat(str, 64)
		r = reflect.ValueOf(&f64)
		err = e
	case reflect.String:
		r = reflect.ValueOf(&str)
	default:
		panic("不支持的类型" + type_.String())
	}
	return
}
