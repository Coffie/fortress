package fortress

import (
	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/fortress/garments"
	"github.com/Coffie/fortress/fortress/users"
	"github.com/jinzhu/gorm"
)

type Fortress struct {
	db *gorm.DB
}

func (f *Fortress) Users() []users.User {
	var dbUsers []database.User
	f.db.Find(&dbUsers)
	allUsers := []users.User{}
	for _, user := range dbUsers {
		allUsers = append(allUsers, users.New(f.db, user))
	}
	return allUsers
}

func (f *Fortress) AddUser(name string, email string) users.User {
	user := database.User{
		Name:  name,
		Email: email,
	}
	f.db.Create(&user)
	return users.New(f.db, user)
}

func (f *Fortress) GetUser(email string) users.User {
	var user database.User
	f.db.First(&user, &database.User{Email: email})
	return users.New(f.db, user)
}

func (f *Fortress) GarmentModels() []garments.GarmentModel {
	var dbModels []database.GarmentModel
	var models []garments.GarmentModel

	f.db.Find(&dbModels)
	for _, model := range dbModels {
		models = append(models, garments.NewGarmentModel(f.db, model))
	}
	return models
}

func (f *Fortress) AddGarmentModel(name string) garments.GarmentModel {
	model := database.GarmentModel{Name: name}
	f.db.Create(&model)
	return garments.NewGarmentModel(f.db, model)
}

func (f *Fortress) GetGarmentModel(name string) garments.GarmentModel {
	var model database.GarmentModel
	f.db.Where(&database.GarmentModel{Name: name}).First(&model)
	return garments.NewGarmentModel(f.db, model)
}
