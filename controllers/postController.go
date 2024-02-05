package controllers

import (
	"go_rest_api/config"
	"go_rest_api/database/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// get post of body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := config.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(200, gin.H{"message": "Post created", "post": post})
}

func GetPosts(c *gin.Context) {
	// Get the post

	var posts []models.Post
	config.DB.Find(&posts)

	c.JSON(200, gin.H{"message": "Post successfully retreived", "posts": posts})
}

func GetAPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	config.DB.First(&post, id)

	c.JSON(200, gin.H{"message": "Post successfully retreived", "post": post})
}
