package main

import (
	"go_rest_api/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong yea",
		})
	})
	r.Run() // listen and serve on 127.0.0.1:4000
}
