package main

import (

"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test",Test)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Test (c *gin.Context) {
	c.JSON(200, gin.H{
		"message" : "Hello Test",

	})
}
