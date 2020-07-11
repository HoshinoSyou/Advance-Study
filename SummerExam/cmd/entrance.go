package cmd

import (
	"Advance/SummerExam/routers"
	"Advance/SummerExam/util/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	acc := r.Group("/userAccount")
	{
		acc.GET("/login", routers.Login)
		acc.POST("/register", routers.Register)
		acc.Use(middleware.Token)
		acc.PUT("/changePassword", routers.ChangePassword)
		acc.DELETE("/deletePassword", routers.Logout)
	}//账户路由组
	roo := r.Group("/room")
	{
		roo.Use(middleware.Token)
		roo.POST("/create", routers.CreateRoom)
		roo.GET("/enter", routers.EnterRoom)
		roo.DELETE("/exit", routers.ExitRoom)
		roo.PUT("/begin",routers.Game)
	}//房间路由组
	r.Run(":8080")
}
