package cmd

import (
	"Advance/workTwo/middleware"
	"Advance/workTwo/service"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.GET("/login", service.Login)
		u.POST("/register", service.Register)
	}
	a := r.Group("/activity")
	{
		a.GET("/Sort", service.Sort)
		a.POST("/join", middleware.Token, service.Join)
		a.PUT("/vote", middleware.Token, service.Vote)
		a.DELETE("/exit", middleware.Token, service.Exit)
	}
}
