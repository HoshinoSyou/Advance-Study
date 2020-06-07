package dao

import (
	"Advance/SummerExam/sqlConn"
	_ "github.com/go-sql-driver/mysql"
)

func Login(username string, password string) bool {
	db := sqlConn.Init()
	var u user
	db.Where("username = ?", username).First(&u)
	if u.Password == password {
		return true
	} else {
		return false
	}
}
