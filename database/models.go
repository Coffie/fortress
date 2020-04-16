package database

type User struct {
	ID    uint
	Name  string
	Email string
}

type OwnedGarment struct {
	ID             uint
	GarmentID      uint
	UserID         uint
	SequenceNumber uint
}

type GarmentModel struct {
	ID     uint
	Name   string
	Images []byte
}

type Garment struct {
	ID             uint
	GarmentModelID uint
	GarmentType    string
	Quantity       int
}

type Flag struct {
	ID      uint
	Country string
	Name    string
}

type FlagTshirt struct {
	ID        uint
	GarmentID uint
	FlagID    uint
	Size      string
	Color     string
}
