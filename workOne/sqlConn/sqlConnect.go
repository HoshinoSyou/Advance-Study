package sqlConn

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func Init() *gorm.DB {
	open, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/bilibiliproject?charset=utf8")
	if err != nil {
		log.Fatal("连接数据库失败喵！错误信息：", err)
		return db
	}
	db = open
	return db
}
