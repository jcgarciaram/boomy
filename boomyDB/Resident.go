package boomyDB

import (
	"log"

	"github.com/jcgarciaram/boomy/chatbot"

	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/utils"
	uuid "github.com/satori/go.uuid"
)

// Resident refers to a resident in the apartment building
type Resident struct {
	ID          string `dynamo:"ID,hash"`
	ResidenceID string `dynamo:"ResidenceID"`

	FirstName   string `dynamo:"FirstName"`
	MiddleName  string `dynamo:"MiddleName"`
	LastName    string `dynamo:"LastName"`
	Email       string `dynamo:"Email"`
	PhoneNumber string `dynamo:"PhoneNumber"`

	ValidationCode string `dynamo:"ValidationCode" json:"-"`

	// Hydrate Up
	Residence Residence `dynamo:"-" json:"-"`

	Conversation chatbot.Conversation `dynamo:"Conversation" json:"-"`
}

// Residents is a slice of Resident
type Residents []Resident

// Save puts Resident struct o in Dynamo
func (o *Resident) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}
	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a Resident struct from Dynamo and unmarshals results into o
func (o *Resident) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *Resident) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *Resident) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all Resident records from Dynamo and unmarshals results into o
func (oSlice *Residents) GetAll() error {
	var o Resident
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	if len(*oSlice) == 0 {
		*oSlice = make([]Resident, 0)
	}
	return nil
}

// GetConversation gets a resident's current conversation
func (o *Resident) GetConversation(ID string) (chatbot.Conversation, error) {
	var c chatbot.Conversation

	if err := o.Get(ID); err != nil {
		log.Println(err)
		return c, err
	}

	return o.Conversation, nil
}
