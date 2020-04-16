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
	for _, garment := range dbGarments {
		garments = append(garments, NewGarment(g.db, garment))
	}
	return garments
}

func (g *GarmentModel) GetGarment(size string, color string) Garment {
	var garment database.Garment
	g.db.Model(&database.GarmentModel{ID: g.Id()}).
		Where(&database.Garment{Size: size, Color: color}).
		Related(&garment)
	return NewGarment(g.db, garment)
}

func (g *GarmentModel) AddGarment(size string, color string, quantity int) Garment {
	garment := database.Garment{
		GarmentModelID: g.model.ID,
		Quantity:       quantity,
		Color:          color,
		Size:           size,
	}
	g.db.Create(&garment)
	return NewGarment(g.db, garment)
}
