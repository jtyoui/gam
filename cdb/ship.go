// Package cdb
// @Time  : 2022/3/2 下午3:22
// @Author: Jtyoui@qq.com
// @note  : sql拼接的关系
package cdb

type ShipType uint // 拼接的关系

const (
	NULL    ShipType = iota
	WHERE            // 并且 等价于 AND
	AND              // 并且
	OR               // 或者
	LIMIT            // 限制
	SELECT           // 选择
	NOT              // 非
	PRELOAD          // 预加载
	ORDER            // 排序
	OMIT             // 省略，一般用在更新上
)
