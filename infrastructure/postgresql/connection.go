package postgresql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Register dialect
)

func Open(dbHost, dbName, dbUser, dbPassword string) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("host=%s dbname=%s user=%s password=%s", dbHost, dbName, dbUser, dbPassword)
	log.Printf("dS: %s", dataSource)
	return gorm.Open("postgres", dataSource)
}
