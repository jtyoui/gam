// Package gam
// @Time  : 2022/2/25 上午9:02
// @Author: Jtyoui@qq.com
// @note  : 主控
package gam

import (
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
