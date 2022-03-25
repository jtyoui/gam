// Package cdb
// @Time  : 2022/3/25 上午11:11
// @Author: Jtyoui@qq.com
// @note  : 例子
package cdb_test

import (
	"fmt"
	"github.com/jtyoui/gam/cdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ExampleNewDBOperate() {
	var DB, _ = gorm.Open(sqlite.Open("gam.db"), &gorm.Config{})
	type User struct {
		Name string `gorm:"column:name"`
		Age  int    `gorm:"column:age"`
	}
	user := &[]User{{Name: "张伟", Age: 18}, {Name: "刘露", Age: 19}}
	g := cdb.NewDBOperate[User](cdb.CREATE, user)
	if db := g.Join(DB); db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(g.Total)
	// ------------
	params := User{Name: "刘露", Age: 19}
	g = cdb.NewDBOperate[User](cdb.FIND, params)
	if db := g.Join(DB); db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(g.Data[0])
	// Output:
	// 2
	// {刘露 19}
}
