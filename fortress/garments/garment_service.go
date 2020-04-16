package garments

import "github.com/Coffie/fortress/database"
import "github.com/jinzhu/gorm"

type GarmentService struct {
	db *gorm.DB
}

func NewGarmentService(db *gorm.DB) *GarmentService {
	return &GarmentService{db: db}
}

func (g *GarmentService) GarmentModels() []GarmentModel {
	var dbModels []database.GarmentModel
	var models []GarmentModel

	g.db.Find(&dbModels)
	for _, model := range dbModels {
		models = append(models, NewGarmentModel(g.db, model))
	}
	return models
}

func (g *GarmentService) AddGarmentModel(name string) GarmentModel {
	model := database.GarmentModel{Name: name}
	g.db.Create(&model)
	return NewGarmentModel(g.db, model)
}

func (g *GarmentService) GetGarmentModel(name string) GarmentModel {
	var model database.GarmentModel
	g.db.Where(&database.GarmentModel{Name: name}).First(&model)
	return NewGarmentModel(g.db, model)
}
