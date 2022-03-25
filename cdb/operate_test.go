// Package cdb_test
// @Time  : 2022/3/18 下午3:41
// @Author: Jtyoui@qq.com
// @note  : 测试
package cdb_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/jtyoui/gam/cdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

var DB, _ = gorm.Open(sqlite.Open("gam.db"), &gorm.Config{})

type User struct {
	gorm.Model
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func createTable() {
	if DB.Migrator().AutoMigrate(User{}) != nil {
		panic("创建表失败")
	}
}

func TestMain(m *testing.M) {
	createTable()
	code := m.Run()
	os.Remove("gam.db")
	os.Exit(code)
}

func newDBOperate01(t *testing.T) {
	user := &[]User{{Name: "张伟", Age: 18}, {Name: "刘露", Age: 19}}
	g := cdb.NewDBOperate[User](cdb.CREATE, user)
	if db := g.Join(DB); db.Error != nil {
		t.Fatal(db.Error)
	}
}

func newDBOperate02(t *testing.T) {
	params := map[string]string{"name": "刘露"}
	g := cdb.NewDBOperate[User](cdb.FIND, params)
	if db := g.Join(DB); db.Error != nil {
		t.Fatal(db.Error)
	}
	assert.Equal(t, g.Data[0].Name, "刘露")
}

func newDBOperate03(t *testing.T) {
	params := User{Name: "刘露"}
	g := cdb.NewDBOperate[User](cdb.FIND, params)
	if db := g.Join(DB); db.Error != nil {
		t.Fatal(db.Error)
	}
	assert.Equal(t, g.Data[0].Name, "刘露")
}

func TestNewDBOperate(t *testing.T) {
	newDBOperate01(t)
	newDBOperate02(t)
	newDBOperate03(t)
}
