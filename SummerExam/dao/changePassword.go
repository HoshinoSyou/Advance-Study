package dao

import (
	"Advance/SummerExam/sqlConn"
)

func Verify(username string, oldPassword string) bool {
	var u user
	db := sqlConn.Init()
	db.Where("username = ? AND password = ?", username, oldPassword).First(&u)
	if u.ID >= 0 {
		return true
	} else {
		return false
	}
}

func ChangePassword(username string, newPassword string) bool {
	var u user
	db := sqlConn.Init()
	db.Where("username = ?", username).First(&u)
	if u.ID >= 0 {
		db.Model(&u).Update("Password", newPassword)
		return true
	} else {
		return false
	}
}
