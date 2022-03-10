// Package cdb
// @Time  : 2022/3/2 下午3:27
// @Author: Jtyoui@qq.com
// @note  : sql的匹配规则
package cdb

type MatchType uint // 匹配类型

const (
	NIL  MatchType = iota // 不进行任何匹配操作
	ACC                   // 精准匹配
	FUZ                   // 模糊匹配
	REG                   // 正则匹配
	IN                    // in匹配
	DESC                  // 降序
	ASC                   // 升序
)

func (m MatchType) String() (s string) {
	switch m {
	case FUZ:
		s = "LIKE ?"
	case ACC:
		s = "= ?"
	case REG:
		s = "REGEXP ?"
	case IN:
		s = "IN ?"
	case DESC:
		s = "DESC"
	case ASC:
		s = "ASC"
	case NIL:
		s = ""
	}
	return
}
