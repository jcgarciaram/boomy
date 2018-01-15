package chatbot

import (
	"log"
	"reflect"
	"runtime"

	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/utils"
	uuid "github.com/satori/go.uuid"
)

// QuickReply are packaged responses
type QuickReply struct {
	ID        string `dynamo:"ID,hash"`
	NodeID    string `dynamo:"NodeID"`
	ReplyText string `dynamo:"ResponseText"`

	responseHandlerMethod
}

// QuickReplies is a slice of QuickReply
type QuickReplies []QuickReply

// QuickReplyPointers is a slice of *QuickReply
type QuickReplyPointers []*QuickReply

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
func NewQuickReply(replyText string, method func(interface{}, string) error) *QuickReply {
	var o QuickReply
	o.GetID()
	o.ReplyText = replyText

	if method != nil {
		RegisterMethod(method)

		// Reflect magic to get the function's name
		o.responseHandlerMethod.methodName = runtime.FuncForPC(reflect.ValueOf(method).Pointer()).Name()
	}

	// o.Save()
	return &o
}

// buildMap builds quickReplyMap which helps when building the tree
func (oSlice QuickReplies) buildMap() {
	for i, qr := range oSlice {
		quickReplyMap[qr.ID] = &oSlice[i]
	}
}

// QuickReplyStringSlice returns a string slice of all the quick reply responses
func QuickReplyStringSlice(qrs []*QuickReply) []string {
	s := make([]string, len(qrs))
	for i, qr := range qrs {
		s[i] = qr.ReplyText
	}
	return s
}

// ResponseHandler handles the reponse received using the method defined in validateResponseMethod
func (o *QuickReply) ResponseHandler(r interface{}, s string) error {

	methodName := o.responseHandlerMethod.methodName
	methodValue := methodMap[methodName]

	errInterface := methodValue.Call([]reflect.Value{reflect.ValueOf(r), reflect.ValueOf(s)})[0].Interface()

	err, ok := errInterface.(error)
	if !ok {
		return nil
	}

	return err

}
