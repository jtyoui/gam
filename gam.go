// Package gam
// @Time  : 2022/2/25 上午9:02
// @Author: Jtyoui@qq.com
// @note  : 主控
package gam

import (
	"github.com/jtyoui/gam/cdb"
	"github.com/jtyoui/gam/router"
	"github.com/jtyoui/gam/web"
)

/*
	NewGinRouter 增加路由地址
*/
var (
	NewGinRouter = router.NewGinRouter
)

/*
	NewDist 静态资源打包
	LoadDistFs embed打包静态资源
*/
var (
	NewDist    = web.NewDist
	LoadDistFs = web.LoadDistFs
)

/*
	NewOperate 获取gorm sql 的操作

	model 增、删、修、查 都是必须的，表结构体

	field 一些sql过滤条件
---------------------------------------------
	例子: 查询Api表id为10的数据
		operate := gam.NewOperate(cdb.FIRST, Api{}, gam.DefaultField("id", 10))
		db = operate.Join(db)
		if err := db.Error; err != nil {
			panic(err)
		}
		fmt.Println(operate.Data)

	根据id=10更新api数据
	operate := gam.NewOperate(cdb.UPDATES, Api{IsFill: true}, gam.DefaultField("id", 10))
	db = operate.Join(db)
	if err := db.Error; err != nil {
		panic(err)
	}
*/
func NewOperate(act cdb.ActionType, model interface{}, field ...cdb.Fielder) *cdb.Operate {
	return &cdb.Operate{
		Fields: field,
		Action: act,
		Model:  model,
	}
}

var (
	NewField     = cdb.NewField
	DefaultField = cdb.DefaultField
)
