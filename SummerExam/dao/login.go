package dao

import (
	"Advance/SummerExam/util/sqlConn"
	_ "github.com/go-sql-driver/mysql"
)

var DB = sqlConn.Init()

func Login(username string, password string) bool {
	var u user
	DB.Where("username = ?", username).First(&u)
	if u.Password == password {
		return true
	} else {
		return false
	}
}//登陆账户
