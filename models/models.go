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
	TshirtGroupID uint
	Size          string
	Color         string
}

type FortressThing struct {
	ID       uint
	Capacity uint
}

type FortressThingInstance struct {
	ID              uint
	FortressThingID uint
	SequenceNumber  uint
}
