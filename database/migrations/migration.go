package migrations

// package main

import (
	"go_rest_api/config"
	"go_rest_api/database/models"
)

// func init() {
// 	config.LoadEnv()
//     config.ConnectDB()
// }

//	func main() {
//		config.DB.AutoMigrate(&models.Post{})
//	}
func Migration() {
	config.DB.AutoMigrate(&models.Post{})
}
