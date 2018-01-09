package demoParkDB

import (
	"log"

	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
	uuid "github.com/satori/go.uuid"
)

// ParkingDeck represents a specific parking deck on a complex
type ParkingDeck struct {
	ID        string `dynamo:"ID,hash"`
	ComplexID string `dynamo:"ComplexID"`
	Name      string `dynamo:"Name"`
	NumFloors int    `dynamo:"NumFloors"`
	NumSpaces int    `dynamo:"NumSpaces"`

	// Hydrate Up
	Complex Complex `dynamo:"-" json:"-"`

	// Hydrate Down
	ParkingSpaces []ParkingSpace `dynamo:"-" json:"-"`
}

// ParkingDecks is a slice of ParkingDeck
type ParkingDecks []ParkingDeck

// Save puts ParkingDeck struct o in Dynamo
func (o *ParkingDeck) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}
	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a ParkingDeck struct from Dynamo and unmarshals results into o
func (o *ParkingDeck) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *ParkingDeck) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *ParkingDeck) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all ParkingDeck records from Dynamo and unmarshals results into o
func (oSlice *ParkingDecks) GetAll() error {
	var o ParkingDeck
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	if len(*oSlice) == 0 {
		*oSlice = make([]ParkingDeck, 0)
	}
	return nil
}
