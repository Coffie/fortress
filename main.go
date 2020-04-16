package main

import (
	"fmt"

	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/fortress/garments"
	"github.com/Coffie/fortress/fortress/users"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var garmentService *garments.GarmentService
var userService *users.UserService

func init() {
	db = database.NewDB("0.0.0.0", "5432", "postgres", "postgres", "")
	database.DropAll(db)
	database.Migrate(db)

	garmentService = garments.NewGarmentService(db)
	userService = users.NewUserService(db)
}

func main() {
	erl := userService.AddUser("Erlend", "erlend@gmail.com")

	gmodel := garmentService.AddGarmentModel("myfirstmodel")
	garment := gmodel.AddGarment("xs", "blue", 10)
	fmt.Println("owned garments:", erl.OwnedGarments())
	ownedGarment, _ := erl.Buy(garment)
	fmt.Println("purchased garment:", erl.OwnedGarments()[0])
	fmt.Println("purchased garment is owned garment:", erl.OwnedGarments()[0] == ownedGarment)

	garmentModel := garmentService.GetGarmentModel("myfirstmodel")
	g := garmentModel.Garments()[0]
	fmt.Println("garment to buy: ", g)
	erl.Buy(garment)
	fmt.Println("owned garments:", erl.OwnedGarments())
}
