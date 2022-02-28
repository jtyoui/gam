// Package base
// @Time  : 2022/2/28 上午11:42
// @Author: Jtyoui@qq.com
// @note  : 将结构体上的binding绑定的错误信息全部转为中文
package base

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()                                  // 实例化验证器
	chinese  = zh.New()                                         // 获取中文翻译器
	uni      = ut.New(chinese, chinese)                         // 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")                          // 获取翻译字典
	_        = zhs.RegisterDefaultTranslations(validate, trans) // 注册翻译器
)

// TranslateError  使用验证器验证结构体
func TranslateError(structs interface{}) (err error) {
	err = validate.Struct(structs)
	if err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			value, _ := json.Marshal(e.Translate(trans))
			err = errors.New(string(value))
		}
	}
	return
}
