package garment

import (
	"github.com/Coffie/fortress/database"
	"github.com/jinzhu/gorm"
)

type Flag struct {
	db    *gorm.DB
	model database.Flag
}

func NewFlag(db *gorm.DB, model database.Flag) Flag {
	return Flag{db, model}
}

func (f *Flag) Country() string {
	return f.model.Country
}

func (f *Flag) Name() string {
	return f.model.Name
}
