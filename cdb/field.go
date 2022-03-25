// Package cdb
// @Time  : 2022/3/2 下午3:23
// @Author: Jtyoui@qq.com
// @note  : sql的属性
package cdb

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

/*
	Fielder 对 gorm.DB 进行封装

	该接口实际上只改变DB的拼接条件，对DB应该保持赋值操作

	列如： DB.Where(?).Or(?)...等等
*/
type Fielder interface {
	Parse(db *gorm.DB) *gorm.DB
}

/*
	field 表示一个命令的最小条件单元，根据条件去拼接db

	其中 k v 一般表示要拼接内容的字符串，变量为k,值为v. match 表示内容的拼接方式，是等于还是like还是不等于。 ship 表示执行条件类型

	例如： Where(name = "joke")为一个最小的条件单元，把它拆开: k 为name, v 为joke, match 为=, ship 为 where
*/
type field struct {
	k, v  interface{}
	match MatchType
	ship  ShipType
}

func (f *field) Parse(db *gorm.DB) *gorm.DB {
	if f.match == FUZ {
		value := f.v.(string)
		if !strings.HasPrefix(value, "%") {
			value = "%" + value
		}
		if !strings.HasSuffix(value, "%") {
			value += "%"
		}
		f.v = value // 模糊操作等价于%?%
	}

	switch f.ship {
	case AND, WHERE:
		factor := fmt.Sprintf("%s %s", f.k, f.match)
		db = db.Where(factor, f.v)
	case OR:
		factor := fmt.Sprintf("%s %s", f.k, f.match)
		db = db.Or(factor, f.v)
	case LIMIT:
		db = db.Offset(f.k.(int)).Limit(f.v.(int))
	case SELECT:
		db = db.Select(f.k)
	case NOT:
		factor := fmt.Sprintf("%s %s", f.k, f.match)
		db = db.Not(factor, f.v)
	case PRELOAD:
		db = db.Preload(f.k.(string))
	case ORDER:
		factor := fmt.Sprintf("%v %s", f.k, f.match)
		db = db.Order(factor)
	case OMIT:
		db = db.Omit(f.k.(string))
	case NULL: // 不操作
	}
	return db
}

// NewField 初始化数据库操作属性
func NewField(key, value any, match MatchType, ship ShipType) Fielder {
	return &field{
		k:     key,
		v:     value,
		match: match,
		ship:  ship,
	}
}

func DefaultField(key, value any) Fielder {
	return NewField(key, value, ACC, AND)
}
