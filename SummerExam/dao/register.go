package dao

import (
	"Advance/SummerExam/sqlConn"
	"github.com/jinzhu/gorm"
)

type user struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(username string, password string) bool {
	var u user
	db := sqlConn.Init()
	db.Where("username = ?", username).First(&u)
	if u.ID >= 0 {
		return false
	} else {
		var us = user{
			Model:    gorm.Model{},
			Username: username,
			Password: password,
		}
		db.Create(&us)
		return true
	}
}
