package boomyDB

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// ParkingDeck represents a specific parking deck on a complex
type ParkingDeck struct {
	gorm.Model
	ComplexID uint
	Name      string
	NumFloors int
	NumSpaces int

	// Hydrate Up
	Complex Complex `gorm:"-" json:"-"`

	// Hydrate Down
	ParkingSpaces []ParkingSpace `gorm:"-" json:"-"`
}

// ParkingDecks is a slice of ParkingDeck
type ParkingDecks []ParkingDeck

func (o *ParkingDeck) Create(conn utils.Conn) error {
	return conn.Create(o).Error
}

func (o *ParkingDeck) Update(conn utils.Conn) error {
	return conn.Update(o).Error
}

func (o *ParkingDeck) First(conn utils.Conn, id uint) error {
	return conn.First(o, id).Error
}

func (o *ParkingDeck) Validate() error {
	return nil
}

func (os ParkingDecks) Find(conn utils.Conn) error {
	return conn.Find(os).Error
}
