package chatbot

import (
	"reflect"

	"github.com/jcgarciaram/boomy/utils"

	"github.com/jinzhu/gorm"
)

// QuickReply are packaged responses
type QuickReply struct {
	gorm.Model
	NodeID    uint   `gorm:"NodeID"`
	ReplyText string `gorm:"ResponseText"`

	ResponseHandlerMethod ResponseHandlerMethod `gorm:"ResponseHandlerMethod"`
}

// QuickReplies is a slice of QuickReply
type QuickReplies []QuickReply

// QuickReplyPointers is a slice of *QuickReply
type QuickReplyPointers []*QuickReply

// Save creates if ID = 0 or update id ID > 0
func (o *QuickReply) Save(conn utils.Conn) error {
	if o.ID == 0 {
		return db.Create(o).Error
	}

	return db.Update(o).Error
}

// Validate validates an object
func (o *QuickReply) Validate() error {
	for _, err := range utils.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// NewQuickReply initializes a QuickReply with a new ID and saves to Dynamo
func NewQuickReply(conn utils.Conn, replyText string, method func(utils.Conn, interface{}, string) error) *QuickReply {
	var o QuickReply
	o.ReplyText = replyText

	if method != nil {
		methodName := RegisterMethod(method)

		// Reflect magic to get the function's name
		o.ResponseHandlerMethod.MethodName = methodName
	}

	o.Save(conn)

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

	methodName := o.ResponseHandlerMethod.MethodName
	methodValue := methodMap[methodName]

	errInterface := methodValue.Call([]reflect.Value{reflect.ValueOf(r), reflect.ValueOf(s)})[0].Interface()

	err, ok := errInterface.(error)
	if !ok {
		return nil
	}

	return err

}
