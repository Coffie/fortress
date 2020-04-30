package tshirts

import (
	"github.com/pkg/errors"

	"github.com/Coffie/fortress/models"
	"github.com/jinzhu/gorm"
)

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

func (t *TshirtService) DeleteTshirtGroup(name string) error {
	return t.db.Delete(&models.TshirtGroup{}, &models.TshirtGroup{Name: name}).Error
}

func (t *TshirtService) AddTshirt(tshirt models.Tshirt) (models.Tshirt, error) {
	err := t.db.Create(&tshirt).Error
	if err != nil {
		return models.Tshirt{}, err
	}
	return tshirt, nil
}

func (t *TshirtService) ListTshirts(tshirtGroupName string) ([]models.Tshirt, error) {
	tshirts := []models.Tshirt{}
	err := t.db.Transaction(func(tx *gorm.DB) error {
		tshirtGroup := models.TshirtGroup{}
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where(&models.TshirtGroup{Name: tshirtGroupName}).
			First(&tshirtGroup).Error; err != nil {
			return err
		}
		if err := tx.Model(&tshirtGroup).Related(&tshirts).Error; err != nil {
			return err
		}
		return nil
	})
	return tshirts, errors.Wrap(err, "could not list tshirts")
}

func (t *TshirtService) DeleteTshirt(tshirtGroupName string, size string, color string) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		tshirtGroup := models.TshirtGroup{}
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where(&models.TshirtGroup{Name: tshirtGroupName}).
			First(&tshirtGroup).Error; err != nil {
			return err
		}
		if err := t.db.
			Delete(&models.Tshirt{}, &models.Tshirt{TshirtGroupID: tshirtGroup.ID, Size: size, Color: color}).
			Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (t *TshirtService) AddFlag(flag models.Flag) (models.Flag, error) {
	err := t.db.Create(&flag).Error
	return flag, err
}

func (t *TshirtService) GetFlag(name string) (models.Flag, error) {
	flag := models.Flag{}
	err := t.db.Where(&models.Flag{Name: name}).First(&flag).Error
	return flag, err
}

func (t *TshirtService) ListFlags() ([]models.Flag, error) {
	flags := []models.Flag{}
	err := t.db.Find(&flags).Error
	return flags, err
}

func (t *TshirtService) DeleteFlag(name string) error {
	return t.db.Delete(&models.Flag{}, &models.Flag{Name: name}).Error
}
