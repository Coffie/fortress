package models

type Flag struct {
	ID      uint
	Name    string
	Country string
}

type TshirtGroup struct {
	ID     uint
	FlagID uint
	Name   string
}

type Tshirt struct {
	ID            uint
	TshirtGroupID uint `gorm:"foreignkey:ID"`
	Size          string
	Color         string
}

type User struct {
	ID                  uint
	Name                string
	Tshirts []Tshirt `gorm:"foreignkey:ID"`
}
