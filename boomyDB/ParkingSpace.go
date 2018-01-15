package boomyDB

import (
	"log"

	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/utils"
	uuid "github.com/satori/go.uuid"
)

// ParkingSpace refers to a specific parking space that belongs to a ParkingDeck which belongs to a Complex
type ParkingSpace struct {
	ID            string `dynamo:"ID,hash"`
	Reference     string `dynamo:"Reference"` // Parking space number
	ComplexID     string `dynamo:"ComplexID"`
	ParkingDeckID string `dynamo:"ParkingDeckID"`
	ParkingTypeID int    `dynamo:"ParkingTypeID"` //Compact, EV, etc.
	FloorNumber   int    `dynamo:"FloorNum"`
	ResidenceID   string `dynamo:"ResidenceID"`
	FloatSpace    bool   `dynamo:"FloatSpace"` // Parking space that does not belong to any one resident

	// Hydrate Up
	ParkingDeck ParkingDeck `dynamo:"-" json:"-"`
	Residence   Residence   `dynamo:"-" json:"-"`
}

// ParkingSpaces is a slice of ParkingSpace
type ParkingSpaces []ParkingSpace

// Save puts ParkingSpace struct o in Dynamo
func (o *ParkingSpace) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}
	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a ParkingSpace struct from Dynamo and unmarshals results into o
func (o *ParkingSpace) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *ParkingSpace) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *ParkingSpace) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all ParkingSpace records from Dynamo and unmarshals results into o
func (oSlice *ParkingSpaces) GetAll() error {
	var o ParkingSpace
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	if len(*oSlice) == 0 {
		*oSlice = make([]ParkingSpace, 0)
	}
	return nil
}
