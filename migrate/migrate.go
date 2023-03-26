package main

import (
	"log"

	"github.com/DaytonBere/go-crud/migrate/scripts"
)

func main () {
	log.Println("Starting migration scripts...")
	scripts.PostMigrate()
	scripts.CommentMigrate()
	log.Println("Done with migration scripts")
}