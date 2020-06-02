package boomyDB

import (
	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// Resident refers to a resident in the apartment building
type Resident struct {
	gorm.Model
	ResidenceID uint

	FirstName   string
	MiddleName  string
	LastName    string
	Email       string
	PhoneNumber string

	ValidationCode string `json:"-"`

	// Hydrate Up
	Residence Residence `gorm:"-" json:"-"`

	Conversation chatbot.Conversation `gorm:"-" json:"-"`
}

// Residents is a slice of Resident
type Residents []Resident

func (o *Resident) Create(conn utils.Conn) error {
	return conn.Create(o).Error
}

func (o *Resident) Update(conn utils.Conn) error {
	return conn.Update(o).Error
}

func (o *Resident) First(conn utils.Conn, id uint) error {
	return conn.First(o, id).Error
}

func (o *Resident) Validate() error {
	return nil
}

func (os Residents) Find(conn utils.Conn) error {
	return conn.Find(&os).Error
}
