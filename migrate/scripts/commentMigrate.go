package scripts

import (
	"log"

	"github.com/DaytonBere/go-crud/initializers"
	"github.com/DaytonBere/go-crud/models"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func CommentMigrate () {
	log.Println("Migrating Comment table...")
	initializers.DB.AutoMigrate(&models.Comment{})
	log.Println("Done migrating Comment table")
}