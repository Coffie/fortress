package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDB(host string, port string, user string, dbname string, password string) *gorm.DB {
	opts := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
	conn, err := gorm.Open("postgres", opts)
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database: %w", err))
	}
	return conn
}

func DropAll(db *gorm.DB) {
	db.DropTableIfExists(&OwnedGarment{})
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Garment{})
	db.DropTableIfExists(&GarmentModel{})
	db.DropTableIfExists("migrations")
}
