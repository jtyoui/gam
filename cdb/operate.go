// Package cdb
// @Time  : 2022/3/2 下午3:56
// @Author: Jtyoui@qq.com
// @note  : sql的主函数
package cdb

import (
	"github.com/jtyoui/gam/tool"
	"gorm.io/gorm"
)

/*
	Operate 操作sql的结构体

	Fields 类似于sql中的where条件

	Model 需要提交哪张表，这是一个go对数据库映射的操作

	Data 数据库返回的数据。这个需要自己去断言

	IDs 如果是匹配操作或者需要更加id，需要对 IDS 进行操作

	Total 对表进行统计返回的值
*/
type Operate struct {
	Fields []Fielder   // 过滤属性去拼接db
	Action ActionType  // 要执行的行为类型
	Model  interface{} // 要请求的表属性
	Data   interface{} // 查询到的数据
	IDs    []int       // 数据库主键
	Total  int64       // 数据库表的总数量
}

func (o *Operate) Join(db *gorm.DB) *gorm.DB {
	db = db.Model(tool.ReflectToObject(o.Model))
	for _, f := range o.Fields {
		db = f.Parse(db)
	}

	var data interface{}
	if o.Data == nil {
		data = tool.CreateArray(o.Model, 0, 10)
	} else {
		data = o.Data
	}

	switch o.Action {
	case CREATE:
		db = db.Create(o.Model)
	case DELETE:
		db = db.Delete(&o.Model, o.IDs)
	case FIRST:
		db = db.First(&data)
	case FIND:
		db = db.Find(&data, o.IDs)
	case UPDATES:
		db = db.Updates(o.Model)
	case TOTAL:
		db = db.Count(&o.Total)
	case SAVE:
		db = db.Save(o.Model)
	case UPDATE:
		// todo
	case LAST:
		db = db.Last(&data)
	case TAKE:
		db = db.Take(&data)
	}

	o.Data = data
	return db
}
