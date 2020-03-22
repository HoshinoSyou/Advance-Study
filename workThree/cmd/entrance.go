package cmd

import (
	"Advance/workThree/router"
	"Advance/workThree/service"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	go service.Manage()
	r.GET("/chat", router.Chat)
	r.Run(":16638")
}
