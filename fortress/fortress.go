package fortress

import (
	"github.com/Coffie/fortress/models"
	"github.com/jinzhu/gorm"
)

type Fortress struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Fortress {
	return &Fortress{db: db}
}

func (f *Fortress) AddThing(thing models.FortressThing) (models.FortressThing, error) {
	err := f.db.Create(&thing).Error
	return thing, err
}

func (f *Fortress) GetThing(id uint) (models.FortressThing, error) {
	var thing models.FortressThing
	err := f.db.Where(&models.FortressThing{ID: id}).First(&thing).Error
	return thing, err
}

func (f *Fortress) DeleteThing(id uint) error {
	return f.db.Delete(&models.FortressThing{}, &models.FortressThing{ID: id}).Error
}
