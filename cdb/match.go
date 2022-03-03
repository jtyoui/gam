// Package cdb
// @Time  : 2022/3/2 下午3:27
// @Author: Jtyoui@qq.com
// @note  : sql的匹配规则
package cdb

type MatchType uint // 匹配类型

const (
	AccurateMatching MatchType = iota // 精准匹配
	FuzzyMatching                     // 模糊匹配
	RegexpMatching                    // 正则匹配
	ArrayMatching                     // in匹配
	Desc                              // 降序
	Asc                               // 升序
)

func (m MatchType) String() (s string) {
	switch m {
	case FuzzyMatching:
		s = "LIKE ?"
	case AccurateMatching:
		s = "= ?"
	case RegexpMatching:
		s = "REGEXP ?"
	case ArrayMatching:
		s = "IN ?"
	case Desc:
		s = "DESC"
	case Asc:
		s = "ASC"
	}
	return
}
