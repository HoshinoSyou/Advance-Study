package cmd

import (
	"Advance/SummerExam/middleware"
	"Advance/SummerExam/routers"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	acc := r.Group("/userAccount")
	{
		acc.GET("/login/:username/:password", routers.Login)
		acc.POST("/register/:username/:password", routers.Register)
		acc.PUT("/changePassword/:oldPassword/:newPassword/:onceNew", middleware.Token, routers.ChangePassword)
		acc.DELETE("deletePassword/:username/:password", middleware.Token, routers.Logout)
	}
	r.Group("/room")
	{
		acc.POST("/create", middleware.Token, routers.CreateRoom)
		acc.GET("/enter", middleware.Token, routers.EnterRoom)
		acc.DELETE("/exit", middleware.Token, routers.ExitRoom)
	}
}
