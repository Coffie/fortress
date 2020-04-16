package garments

import "github.com/Coffie/fortress/database"
import "github.com/jinzhu/gorm"

type GarmentModel struct {
	db    *gorm.DB
	model database.GarmentModel
}

func NewGarmentModel(db *gorm.DB, model database.GarmentModel) GarmentModel {
	return GarmentModel{
		db:    db,
		model: model,
	}
}

func (g *GarmentModel) Id() uint {
	return g.model.ID
}

func (g *GarmentModel) Name() string {
	return g.model.Name
}

func (g *GarmentModel) Garments() []Garment {
	var dbGarments []database.Garment
	g.db.Model(&database.GarmentModel{ID: g.Id()}).Related(&dbGarments)
	garments := []Garment{}
	for _, currentGarment := range dbGarments {
		garments = append(garments, NewGarment(g.db, currentGarment))
	}
	return garments
}

func (g *GarmentModel) AddGarment(quantity int, garmentType string) Garment {
	dbGarment := database.Garment{
		GarmentModelID: g.model.ID,
		GarmentType:    garmentType,
		Quantity:       quantity,
	}
	g.db.Create(&dbGarment)
	return NewGarment(g.db, dbGarment)
}
