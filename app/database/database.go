package database

import (
	"log"

	"github.com/jinzhu/gorm"
	// Import the postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database holds the database context
type Database struct {
	Ref *gorm.DB
}

// New returns a database instance
func New() *Database {
	var db Database
	var err error
	// db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	db.Ref, err = gorm.Open("postgres", "host=127.0.0.1 user=alextanhongpin password=123456 dbname=grpc-openid sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return &db
}

// func (db *Database) Setup () {
// }
