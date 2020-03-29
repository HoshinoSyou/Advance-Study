package main

import (
	"Advance/WebFrame/webogo"
	"net/http"
)

func main() {
	r := webogo.Default()
	r.GET("/user/id", func(c *webogo.Context) {
		id := c.Query("username")
		c.JSON(200, webogo.F{
			"status":  http.StatusOK,
			"message": "该用户id为" + id,
		})
	})
	r.GET("/user/onlineOrNot", func(c *webogo.Context) {
		b := c.GetBool("username")
		if b {
			c.XML(200, webogo.F{
				"status":  http.StatusOK,
				"message": "该用户在线",
			})
		} else {
			c.XML(200, webogo.F{
				"status":  http.StatusOK,
				"message": "该用户不在线",
			})
		}
	})
	r.Run(":16637")
}
