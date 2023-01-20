package main

import (
	"go-crud-gin-gorm/initializers"
	"go-crud-gin-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
