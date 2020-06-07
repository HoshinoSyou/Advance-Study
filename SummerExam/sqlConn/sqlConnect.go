package sqlConn

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	open, err := gorm.Open("mysql", "root:syouZX@(127.0.0.1:3306)/bilibiliproject?charset=utf8")
	if err != nil {
		log.Fatal("连接数据库失败喵！错误信息：", err)
		return db
	}
	db = open
	return db
}
