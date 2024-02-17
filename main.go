package main

import (
	"go_rest_api/config"
	"go_rest_api/controllers"
	"go_rest_api/database/migrations"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	migrations.Migration()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetAPost)
	r.PATCH("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run() // listen and serve on 127.0.0.1:4000
}
