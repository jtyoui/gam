// Package cdb
// @Time  : 2022/3/2 下午3:59
// @Author: Jtyoui@qq.com
// @note  : sql行为，表示要执行的动作
package cdb

type ActionType uint // 要执行的行为类型

const (
	CREATE ActionType = iota
	FIRST
	LAST
	FIND
	DELETE
	UPDATE
	UPDATES // 更新整个表
	TOTAL   // 获取表的总数量
	SAVE
	TAKE
)
