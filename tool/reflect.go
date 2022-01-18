// Package tool
// @Time  : 2022/1/17 上午9:40
// @Author: Jtyoui@qq.com
// @note  : 反射工具类
package tool

import (
	"reflect"
)

// ReflectByValue 根据结构体获取该结构体的值类型，
// 传入的结构体必须是指针类型
func ReflectByValue(router interface{}) reflect.Value {
	t := reflect.TypeOf(router)
	if t.Kind() != reflect.Ptr {
		panic("必须传入指针类型")
	}
	return reflect.ValueOf(router)
}

// IsNil 判断反射类型是不是为nil
func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	}
	return false
}
