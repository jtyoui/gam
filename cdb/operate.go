// Package cdb
// @Time  : 2022/3/2 下午3:56
// @Author: Jtyoui@qq.com
// @note  : sql的主函数
package cdb

import (
	"gorm.io/gorm"
)

/*
	Operate 表示一条完整的sql命令，包括需要执行的操作，数据库表，返回数据等基本行为。

	Fields 无数的sql条件单元

	Model 需要提交哪张表，这是一个go对数据库映射的操作

	Data 数据库返回的数据。这个需要自己去断言

	IDs 如果是匹配操作或者需要更加id，需要对 IDS 进行操作

	Total 对表进行统计返回的值
*/
type DBOperate[T any] struct {
	Fields []Fielder  // 过滤属性去拼接db
	Action ActionType // 要执行的行为类型
	Model  any        // 要请求的表属性
	Data   []T        // 查询到的数据
	IDs    []int      // 数据库主键
	Total  int64      // 数据库表的总数量
}

/*
	NewOperate 获取gorm sql 的操作

	model 增、删、修、查 都是必须的，表结构体

	field 一些sql过滤条件
*/
func NewDBOperate[T any](act ActionType, model any, field ...Fielder) *DBOperate[T] {
	return &DBOperate[T]{
		Fields: field,
		Action: act,
		Model:  model,
	}
}

func (o *DBOperate[T]) Join(db *gorm.DB) *gorm.DB {
	db = db.Model(new(T))

	for _, f := range o.Fields {
		db = f.Parse(db)
	}

	switch o.Action {
	case CREATE:
		db = db.Create(o.Model)
	case DELETE:
		db = db.Delete(&o.Model, o.IDs)
	case FIRST:
		db = db.Where(o.Model).First(&o.Data)
	case FIND:
		db = db.Where(o.Model).Find(&o.Data, o.IDs)
	case UPDATES:
		/*
			注意一下 GORM 只会更新非零值的字段，如果需要更新零值字段需要和Select联合使用 https://gorm.io/zh_CN/docs/update.html#%E6%9B%B4%E6%96%B0%E5%A4%9A%E5%88%97

			f1 := NewField("IsFill", "", cdb.NIL, cdb.SELECT) // 选中可能的非零值，如果要全部更新，可以使用*
			f2 := NewField("id", 10, cdb.ACC, cdb.AND)       // 更新的筛选条件
			operate := NewOperate[Api](cdb.UPDATES, Api{IsFill: false}, f1, f2)
			if err := operate.Join(db).Error; err != nil {
				panic(err)
			}
		*/
		db = db.Updates(o.Model)
	case TOTAL:
		db = db.Where(o.Model).Count(&o.Total)
	case SAVE:
		db = db.Save(o.Model)
	case LAST:
		db = db.Where(o.Model).Last(&o.Data)
	case TAKE:
		db = db.Where(o.Model).Take(&o.Data)
	}
	o.Total = db.RowsAffected
	return db
}
