package dao

import "github.com/jinzhu/gorm"

func ChangePassword(username string, oldPassword string, newPassword string) bool {
	var u = user{
		Model:    gorm.Model{},
		Username: username,
		Password: oldPassword,
	}
	res := DB.NewRecord(u)
	if res {
		DB.Model(&user{
			Model:    gorm.Model{},
			Username: username,
			Password: oldPassword,
		}).Where("Username = ?", username).Update("Password", newPassword)
		return true
	} else {
		return false
	}
}//修改密码
