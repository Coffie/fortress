package garment

import (
	"github.com/Coffie/fortress/database"
	"github.com/Coffie/fortress/fortress/garments"
	"github.com/jinzhu/gorm"
)

type FlagTshirt struct {
	garments.Garment
	db    *gorm.DB
	model database.FlagTshirt
}

func NewFlagTshirt(db *gorm.DB, garment garments.Garment, model database.FlagTshirt) FlagTshirt {
	return FlagTshirt{garment, db, model}
}

func (f *FlagTshirt) Size() string {
	return f.model.Size
}

func (f *FlagTshirt) Color() string {
	return f.model.Color
}

func (f *FlagTshirt) Flag() Flag {
	var flag database.Flag
	f.db.Model(&database.FlagTshirt{ID: f.Id()}).Related(&flag)
	return NewFlag(f.db, flag)
}
