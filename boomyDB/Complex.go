package boomyDB

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// Complex refers to an entire complex which could contain multiple parking decks
type Complex struct {
	gorm.Model
	Name             string
	AddressLine1     string
	AddressLine2     string
	City             string
	State            string
	ZipCode          string
	PrimaryContactID uint

	// Hydrate Down
	ParkingDecks []ParkingDeck `gorm:"-" json:"-"`
	Residences   []Residence   `gorm:"-" json:"-"`
}

// Complexes is a slice of Complex
type Complexes []Complex

// TypeHelper will be used to be able to use shared common functions across objects
type TypeHelper interface {
	Create() error
	Update() error
	First(uint) error
	Validate() error
}

func (o *Complex) Create(conn utils.Conn) error {
	return conn.Create(o).Error
}

func (o *Complex) Update(conn utils.Conn) error {
	return conn.Update(o).Error
}

func (o *Complex) First(conn utils.Conn, id uint) error {
	return conn.First(o, id).Error
}

func (o *Complex) Validate() error {
	return nil
}

func (os Complexes) Find(conn utils.Conn) error {
	return conn.Find(os).Error
}
