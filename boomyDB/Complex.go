package boomyDB

import (
	"log"

	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/utils"
	uuid "github.com/satori/go.uuid"
)

// Complex refers to an entire complex which could contain multiple parking decks
type Complex struct {
	ID               string `dynamo:"ID,hash"`
	Name             string `dynamo:"Name" validate:"string,min=0,max=30"`
	AddressLine1     string `dynamo:"AddressLine1" validate:"string,min=0,max=30"`
	AddressLine2     string `dynamo:"AddressLine2" validate:"string,min=0,max=30"`
	City             string `dynamo:"City" validate:"string,min=0,max=30"`
	State            string `dynamo:"State" validate:"string,min=0,max=10"`
	ZipCode          string `dynamo:"ZipCode" validate:"string,min=0,max=5"`
	PrimaryContactID string `dynamo:"PrimaryContactID"`

	// Hydrate Down
	ParkingDecks []ParkingDeck `dynamo:"-" json:"-"`
	Residences   []Residence   `dynamo:"-" json:"-"`
}

// Complexes is a slice of Complex
type Complexes []Complex

// Save puts Complex struct o in Dynamo
func (o *Complex) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a Complex struct from Dynamo and unmarshals results into o
func (o *Complex) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *Complex) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *Complex) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all Complex records from Dynamo and unmarshals results into o
func (oSlice *Complexes) GetAll() error {
	var o Complex
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}

	if len(*oSlice) == 0 {
		*oSlice = make([]Complex, 0)
	}

	return nil
}
