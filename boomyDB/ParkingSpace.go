package boomyDB

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// ParkingSpace refers to a specific parking space that belongs to a ParkingDeck which belongs to a ParkingSpace
type ParkingSpace struct {
	gorm.Model
	Reference     string // Parking space number
	ComplexID     uint
	ParkingDeckID uint
	ParkingTypeID int //Compact, EV, etc.
	FloorNumber   int
	ResidenceID   uint
	FloatSpace    bool // Parking space that does not belong to any one resident

	// Hydrate Up
	ParkingDeck ParkingDeck `gorm:"-" json:"-"`
	Residence   Residence   `gorm:"-" json:"-"`
}

// ParkingSpaces is a slice of ParkingSpace
type ParkingSpaces []ParkingSpace

func (o *ParkingSpace) Create(conn utils.Conn) error {
	return conn.Create(o).Error
}

func (o *ParkingSpace) Update(conn utils.Conn) error {
	return conn.Update(o).Error
}

func (o *ParkingSpace) First(conn utils.Conn, id uint) error {
	return conn.First(o, id).Error
}

func (o *ParkingSpace) Validate() error {
	return nil
}

func (os ParkingSpaces) Find(conn utils.Conn) error {
	return conn.Find(os).Error
}
