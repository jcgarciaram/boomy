package boomyDB

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// Residence refers to a specific residence in a complex. It can contain one or more parking spaces
type Residence struct {
	gorm.Model
	Reference         string // Apartment number, house number, etc.
	FloorNum          int
	PrimaryResidentID string

	// Hydrate Down
	Residents     []Resident     `gorm:"-" json:"-"`
	ParkingSpaces []ParkingSpace `gorm:"-" json:"-"`
}

// Residences is a slice of Residence
type Residences []Residence

func (o *Residence) Create(conn utils.Conn) error {
	return conn.Create(o).Error
}

func (o *Residence) Update(conn utils.Conn) error {
	return conn.Update(o).Error
}

func (o *Residence) First(conn utils.Conn, id uint) error {
	return conn.First(o, id).Error
}

func (o *Residence) Validate() error {
	return nil
}

func (os Residences) Find(conn utils.Conn) error {
	return conn.Find(os).Error
}
