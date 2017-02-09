package dbconfig

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// Ignore
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lkkadiri/goboot/models"
)

var db *gorm.DB

// Init the app
func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=goboot_user dbname=goboot password=goboot sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	if !db.HasTable(models.User{}) {
		db.CreateTable(models.User{})
		fmt.Print("Users table migrated.\n")
	}

	return db
}
