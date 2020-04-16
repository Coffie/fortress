package users

import "github.com/Coffie/fortress/database"
import "github.com/Coffie/fortress/fortress/garments"
import "github.com/jinzhu/gorm"

type OwnedGarment struct {
	db    *gorm.DB
	model database.OwnedGarment
}

func NewOwnedGarment(db *gorm.DB, model database.OwnedGarment) OwnedGarment {
	return OwnedGarment{db: db, model: model}
}

func (og *OwnedGarment) Id() uint {
	return og.model.ID
}

func (og *OwnedGarment) SeqNr() int {
	return int(og.model.SequenceNumber)
}

func (og *OwnedGarment) Garment() garments.Garment {
	var garment database.Garment
	og.db.First(&garment, &database.OwnedGarment{ID: og.Id()})
	return garments.NewGarment(og.db, garment)
}
