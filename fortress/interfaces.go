package fortress

type Country string
type Color string

type User interface {
	Name() string
	Email() string
	Garments() []Garment
}

type Size interface {
	Name() string
	Spec() string
}

type Flag interface {
	Name() string
	Country() Country
	Image() []byte
}

type Garment interface {
	Model() GarmentModel
	Owner() User
	SeqNr() int
}

type GarmentModel interface {
	Color() Color
	Size() Size
	Image() []byte
	AmountInStock() int
	AmountInRotation() int
	InStock() bool
	GarmentCollection() GarmentCollection
}

type GarmentCollection interface {
	Name() string
	GarmentModels() []GarmentModel
	GetGarmentModel(Name string) GarmentModel
	GetRandomGarmentModel(size Size) GarmentModel
	AddGarmentModel(GarmentModel)
	Flag() Flag
}

type GarmentStore interface {
	GarmentCollections() []GarmentCollection
	AddCollection(collection GarmentCollection)
	GetGarmentCollection(Name string) GarmentCollection
	GetRandomGarmentCollection() GarmentCollection

	Users() []User
	AddUser(user User)
	GetUser(name User)

	BuyGarment(user User, model GarmentModel)
}
