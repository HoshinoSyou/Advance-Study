package sqlConn

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func Init() *gorm.DB {
	open, err := gorm.Open("mysql", "root:@(localhost)/wuzichess?charset=utf8")
	if err != nil {
		log.Fatal("连接数据库失败喵！错误信息：", err)
		return DB
	}
	DB = open
	return DB
}//连接并初始化数据库
