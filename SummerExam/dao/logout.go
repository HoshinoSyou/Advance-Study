package dao


func Logout(username string,password string) bool {
	var u user
	DB.Where("username = ?", username).First(&u)
	if u.Password == password {
		DB.Delete(&u)
		return true
	} else {
		return false
	}
}//注销账户
