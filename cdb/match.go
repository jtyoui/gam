// Package cdb
// @Time  : 2022/3/2 下午3:27
// @Author: Jtyoui@qq.com
// @note  : sql的匹配规则
package cdb

type MatchType uint // 匹配类型

const (
	NULLMatch     MatchType = iota
	AccurateMatch           // 精准匹配
	FuzzyMatch              // 模糊匹配
	RegexpMatch             // 正则匹配
	ArrayMatch              // in匹配
	DescMatch               // 降序
	AscMatch                // 升序
)

func (m MatchType) String() (s string) {
	switch m {
	case FuzzyMatch:
		s = "LIKE ?"
	case AccurateMatch:
		s = "= ?"
	case RegexpMatch:
		s = "REGEXP ?"
	case ArrayMatch:
		s = "IN ?"
	case DescMatch:
		s = "DESC"
	case AscMatch:
		s = "ASC"
	}
	return
}
