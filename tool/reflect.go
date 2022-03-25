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
func ReflectByValue(object interface{}) reflect.Value {
	t := reflect.ValueOf(object)
	if t.Kind() != reflect.Ptr {
		panic("必须传入指针类型")
	}
	return t
}

// RemovePtr 去掉指针
func RemovePtr(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
