package dao

import (
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(username string, password string) bool {
	DB.AutoMigrate(&user{})
	var u = user{
		Model:    gorm.Model{},
		Username: username,
		Password: password,
	}
	var us user
	DB.Where("username = ?", username).First(&us)
	if us.ID <= 0 {
		DB.Create(&u)
		return true
	} else {
		return false
	}
} //注册账户
