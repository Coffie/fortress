package database

type User struct {
	ID    uint
	Name  string
	Email string
}

type GarmentModel struct {
	ID     uint
	Name   string
	Images []byte
}

type Garment struct {
	ID             uint
	GarmentModelID uint
	Quantity       int
	Color          string
	Size           string
}

type OwnedGarment struct {
	ID             uint
	GarmentID      uint
	UserID         uint
	SequenceNumber uint
}
