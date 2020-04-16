package users

import (
	"fmt"

	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/fortress/garments"
	"github.com/jinzhu/gorm"
)

type User struct {
	db    *gorm.DB
	model database.User
}

func NewUser(db *gorm.DB, model database.User) User {
	return User{db: db, model: model}
}

func (u *User) Id() uint {
	return u.model.ID
}

func (u *User) Name() string {
	return u.model.Name
}

func (u *User) Email() string {
	return u.model.Email
}

func (u *User) OwnedGarments() []OwnedGarment {
	var dbOwnedGarments []database.OwnedGarment
	u.db.Model(&database.User{ID: u.Id()}).Related(&dbOwnedGarments)

	var ownedGarments []OwnedGarment
	for _, og := range dbOwnedGarments {
		ownedGarments = append(ownedGarments, NewOwnedGarment(u.db, og))
	}
	return ownedGarments
}

func (u *User) Buy(garment garments.Garment) (OwnedGarment, error) {
	// TODO: this should all be done in a transaction
	amountInStock := garment.AmountInStock()
	if amountInStock == 0 {
		return OwnedGarment{}, fmt.Errorf("No more garments in stock")
	}
	seqNr := uint(garment.Quantity() - amountInStock + 1)
	ownedGarment := u.createOwnedGarment(seqNr, garment.Id())
	return NewOwnedGarment(u.db, ownedGarment), nil
}

func (u *User) createOwnedGarment(seqNr uint, garmentId uint) database.OwnedGarment {
	ownedGarment := database.OwnedGarment{
		GarmentID:      garmentId,
		UserID:         u.Id(),
		SequenceNumber: seqNr,
	}
	u.db.Create(&ownedGarment)
	return ownedGarment
}
