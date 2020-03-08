package cmd

import (
	"Advance/workOne/middleware"
	"Advance/workOne/userService"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	acc := r.Group("/userAccount")
	{
		acc.GET("/login/:username/:password", userService.Login)
		acc.GET("/register/:username/:password", userService.Register)
	}
	ui := r.Group("/userInformation")
	{
		ui.GET("/user/query/:username", userService.QueryUser)
		ui.GET("/user/upload", middleware.Token, userService.UploadInformation)
		ui.GET("/user/set/:username", middleware.Token, userService.SetInformation)
	}
	r.Run(":8080")
}
