package main

import (
	"github.com/DaytonBere/go-crud/initializers"
	"github.com/DaytonBere/go-crud/models"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main () {
	initializers.DB.AutoMigrate(&models.Post{})
}