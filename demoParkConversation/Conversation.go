package demoParkConversation

import (
	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
	uuid "github.com/satori/go.uuid"

	"log"
)

// Conversation is a single conversation with a recipient
type Conversation struct {
	SenderID           string `dynamo:"ID,hash"`
	ConversationTreeID string `dynamo:"ConversationTreeID"`
	CurrentNodeID      string `dynamo:"CurrentNodeID"`
}

// Conversations is a slice of Conversation
type Conversations []Conversation

// Save puts struct o in Dynamo
func (o *Conversation) Save() error {
	if o.SenderID == "" {
		o.SenderID = uuid.NewV4().String()
	}

	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a struct from Dynamo and unmarshals results into o
func (o *Conversation) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *Conversation) GetID() string {
	if o.SenderID == "" {
		o.SenderID = uuid.NewV4().String()
	}

	return o.SenderID
}

// Validate validates an object
func (o *Conversation) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all records from Dynamo and unmarshals results into o
func (oSlice *Conversations) GetAll() error {
	var o Conversation
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}

	if len(*oSlice) == 0 {
		*oSlice = make([]Conversation, 0)
	}

	return nil
}
