package demoParkConversation

import (
	"log"

	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
	uuid "github.com/satori/go.uuid"
)

// QuickReply are packaged responses
type QuickReply struct {
	ID        string `dynamo:"ID,hash"`
	NodeID    string `dynamo:"NodeID"`
	ReplyText string `dynamo:"ResponseText"`
}

// QuickReplies is a slice of QuickReply
type QuickReplies []QuickReply

// Save puts QuickReply struct o in Dynamo
func (o *QuickReply) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a QuickReply struct from Dynamo and unmarshals results into o
func (o *QuickReply) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *QuickReply) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// Validate validates an object
func (o *QuickReply) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetAll gets all QuickReply records from Dynamo and unmarshals results into o
func (oSlice *QuickReplies) GetAll() error {
	var o QuickReply
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}

	if len(*oSlice) == 0 {
		*oSlice = make([]QuickReply, 0)
	}

	return nil
}

// NewQuickReply initializes a QuickReply with a new ID and saves to Dynamo
func NewQuickReply(replyText string) *QuickReply {
	var o QuickReply
	o.GetID()
	o.ReplyText = replyText

	o.Save()
	return &o
}

// buildMap builds quickReplyMap which helps when building the tree
func (oSlice QuickReplies) buildMap() {
	for i, qr := range oSlice {
		quickReplyMap[qr.ID] = &oSlice[i]
	}
}
