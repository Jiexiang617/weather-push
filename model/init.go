package model

import (
	"github.com/jinzhu/gorm"
 	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
	"weather-push/util"
)

// 数据库连接单例
var DB *gorm.DB

// Database 在中间件中初始化 sqlite 连接
func Database(path string) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	db.LogMode(true)

	// 连接池
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Second)

	DB = db

	// 自动建表
	migration()
}
