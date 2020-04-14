package tshirts

import "github.com/Coffie/fortress/models"
import "github.com/jinzhu/gorm"

type TshirtService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TshirtService {
	return &TshirtService{db}
}

func (t *TshirtService) GetTshirtGroup(name string) (models.TshirtGroup, error) {
	var tshirtGroup models.TshirtGroup
	err := t.db.Where(&models.TshirtGroup{Name: name}).First(&tshirtGroup).Error
	return tshirtGroup, err
}

func (t *TshirtService) ListTshirtGroups() ([]models.TshirtGroup, error) {
	var tshirtGroups []models.TshirtGroup
	err := t.db.Find(&tshirtGroups).Error
	return tshirtGroups, err
}

func (t *TshirtService) AddTshirtGroup(tshirtGroup models.TshirtGroup) (models.TshirtGroup, error) {
	err := t.db.Create(&tshirtGroup).Error
	return tshirtGroup, err
}

func (t *TshirtService) AddTshirt(tshirt models.Tshirt) (models.Tshirt, error) {
	err := t.db.Create(&tshirt).Error
	return tshirt, err
}

func (t *TshirtService) ListTshirts(tshirtGroupName string) ([]models.Tshirt, error) {
	tshirts := []models.Tshirt{}
	err := t.db.Model(&models.TshirtGroup{Name: tshirtGroupName}).Related(&tshirts).Error
	return tshirts, err
}

func (t *TshirtService) AddFlag(flag models.Flag) (models.Flag, error) {
	err := t.db.Create(&flag).Error
	return flag, err
}
