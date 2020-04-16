package garments

import (
	"github.com/Coffie/fortress/database"
	"github.com/jinzhu/gorm"
)

type Garment struct {
	db    *gorm.DB
	model database.Garment
}

func NewGarment(db *gorm.DB, model database.Garment) Garment {
	return Garment{db: db, model: model}
}

func (g *Garment) AmountInRotation() int {
	var ownedGarments []database.OwnedGarment
	var count int
	g.db.Model(&database.Garment{ID: g.Id()}).Related(&ownedGarments).Count(&count)
	return count
}

func (g *Garment) Id() uint {
	return g.model.ID
}

func (g *Garment) Quantity() int {
	return g.model.Quantity
}

func (g *Garment) Type() string {
	return g.model.GarmentType
}

func (g *Garment) AmountInStock() int {
	return g.model.Quantity - g.AmountInRotation()
}

func (g *Garment) InStock() bool {
	return g.AmountInStock() > 0
}
