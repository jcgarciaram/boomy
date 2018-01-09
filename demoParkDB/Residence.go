package demoParkDB

import (
	"log"

	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
	uuid "github.com/satori/go.uuid"
)

// Residence refers to a specific residence in a complex. It can contain one or more parking spaces
type Residence struct {
	ID                string `dynamo:"ID,hash"`
	Reference         string `dynamo:"Reference"` // Apartment number, house number, etc.
	FloorNum          int    `dynamo:"FloorNum"`
	PrimaryResidentID string `dynamo:"PrimaryResidentID"`

	// Hydrate Down
	Residents     []Resident     `dynamo:"-" json:"-"`
	ParkingSpaces []ParkingSpace `dynamo:"-" json:"-"`
}

// Residences is a slice of Residence
type Residences []Residence

// Save puts Residence struct o in Dynamo
func (o *Residence) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}
	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a Residence struct from Dynamo and unmarshals results into o
func (o *Residence) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *Residence) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *Residence) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all Residence records from Dynamo and unmarshals results into o
func (oSlice *Residences) GetAll() error {
	var o Residence
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	if len(*oSlice) == 0 {
		*oSlice = make([]Residence, 0)
	}
	return nil
}
