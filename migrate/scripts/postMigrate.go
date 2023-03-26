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

func PostMigrate () {
	log.Println("Migrating Post table...")
	initializers.DB.AutoMigrate(&models.Post{})
	log.Println("Done migrating Post table")
}