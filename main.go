package main

import (
	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/tshirts"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var tshirtService *tshirts.TshirtService

func init() {
	db = database.NewDB("0.0.0.0", "5432", "postgres", "postgres", "")
	database.DropAll(db)
	database.Migrate(db)

	tshirtService = tshirts.New()
}

func main() {

}
