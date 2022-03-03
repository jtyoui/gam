// Package cdb
// @Time  : 2022/3/2 下午3:23
// @Author: Jtyoui@qq.com
// @note  : sql的属性
package cdb

import (
	"fmt"
	"gorm.io/gorm"
)

type Fielder interface {
	Parse(db *gorm.DB) *gorm.DB
}

// 根据过滤属性去拼接db
type field struct {
	k, v  interface{}
	match MatchType
	ship  ShipType
}

func (f *field) Parse(db *gorm.DB) *gorm.DB {
	if f.match == FuzzyMatching {
		f.v = fmt.Sprintf("%%%s%%", f.v)
	}

	switch f.ship {
	case AND:
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
	}
	return db
}

// NewField 初始化数据库操作属性
func NewField(key, value interface{}, match MatchType, ship ShipType) Fielder {
	return &field{
		k:     key,
		v:     value,
		match: match,
		ship:  ship,
	}
}

func DefaultField(key, value interface{}) Fielder {
	return NewField(key, value, AccurateMatching, AND)
}
