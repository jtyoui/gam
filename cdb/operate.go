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
		/*
			注意一下 GORM 只会更新非零值的字段，如果需要更新零值字段需要和Select联合使用 https://gorm.io/zh_CN/docs/update.html#%E6%9B%B4%E6%96%B0%E5%A4%9A%E5%88%97

			f1 := gam.NewField("IsFill", "", cdb.NULLMatch, cdb.SELECT) // 选中可能的非零值，如果要全部更新，可以使用*
			f2 := gam.NewField("id", 10, cdb.AccurateMatch, cdb.AND)    // 更新的筛选条件
			operate := gam.NewOperate(cdb.UPDATES, Api{IsFill: false}, f1, f2)
			if err := operate.Join(db).Error; err != nil {
				panic(err)
			}
		*/
		db = db.Updates(o.Model)
	case TOTAL:
		db = db.Count(&o.Total)
	case SAVE:
		db = db.Save(o.Model)
	case LAST:
		db = db.Last(&data)
	case TAKE:
		db = db.Take(&data)
	}
	o.Total = db.RowsAffected
	o.Data = data
	return db
}
