package users

import "github.com/Coffie/fortress/database"
import "github.com/jinzhu/gorm"

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (g *UserService) Users() []User {
	var dbUsers []database.User
	g.db.Find(&dbUsers)
	allUsers := []User{}
	for _, user := range dbUsers {
		allUsers = append(allUsers, NewUser(g.db, user))
	}
	return allUsers
}

func (g *UserService) AddUser(name string, email string) User {
	user := database.User{
		Name:  name,
		Email: email,
	}
	g.db.Create(&user)
	return NewUser(g.db, user)
}

func (g *UserService) GetUser(email string) User {
	var user database.User
	g.db.First(&user, &database.User{Email: email})
	return NewUser(g.db, user)
}
