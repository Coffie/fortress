package main

import "fmt"
import "github.com/Coffie/fortress/database"
import "github.com/Coffie/fortress/fortress/garments"
import "github.com/Coffie/fortress/fortress/users"
import "github.com/jinzhu/gorm"

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
	g := garmentModel.GetGarment("xs", "blues")
	fmt.Println("garment to buy: ", g)
	erl.Buy(garment)
	fmt.Println("owned garments:", erl.OwnedGarments())
}
