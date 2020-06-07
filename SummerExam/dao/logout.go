package dao

import "Advance/SummerExam/sqlConn"


func Logout(username string,password string) bool {
	db := sqlConn.Init()
	var u user
	db.Where("username = ?", username).First(&u)
	if u.Password == password {
		db.Delete(&u)
		return true
	} else {
		return false
	}
}
