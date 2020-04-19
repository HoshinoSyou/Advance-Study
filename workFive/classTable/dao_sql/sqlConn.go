package dao_sql

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	open, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/bilibiliproject?charset=utf8")
	if err != nil {
		log.Fatalf("连接数据库失败喵！错误信息：%v", err)
		return db
	}
	return open
}
