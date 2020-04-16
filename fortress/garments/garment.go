package garments

import "github.com/Coffie/fortress/database"
import "github.com/jinzhu/gorm"

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

func (g *Garment) AmountInStock() int {
	return g.model.Quantity - g.AmountInRotation()
}

func (g *Garment) InStock() bool {
	return g.AmountInStock() > 0
}

func (g *Garment) Id() uint {
	return g.model.ID
}

func (g *Garment) Color() string {
	return g.model.Color
}

func (g *Garment) Size() string {
	return g.model.Size
}

func (g *Garment) Quantity() int {
	return g.model.Quantity
}
