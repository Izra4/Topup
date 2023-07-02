package initializers

import (
	"TopUpWEb/database"
	"log"
)

func SyncDb() {
	db := database.InitDB()
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to Migrate")
	}
}
