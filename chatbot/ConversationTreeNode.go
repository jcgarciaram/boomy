package chatbot

import (
	"fmt"
	"reflect"

	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// ConversationTreeNode is saved in Dynamo. Used by ConversationTree to build up the tree in memory
type ConversationTreeNode struct {
	gorm.Model
	ConversationTreeID   uint
	ResponseText         string
	ExpectedReplyType    int
	QuickReplyIDs        []uint `gorm:"-"`
	ChildrenNodeIDs      []uint `gorm:"-"`
	DayIntervalIncrement int
	IsRootNode           bool
	ParentNodeResponse   interface{} `gorm:"-"`

	// We can store a struct that will contain the name of the type that contains the function and the name of function itself in strings.
	// We then need another function that has a default variable of all necessary types and depending on the type uses that default variable and reflection below to call the method. All methods should return bool
	// See the following example:
	//
	// https://play.golang.org/p/no9z-0wvkmy
	//
	// As part of initializing the dynamoDB package, we can add all the types to a map that will contain the reflect.Type of the type. This map needs to be visible to boomyConversation
	//
	// https://play.golang.org/p/2TugHjN1IUk
	//
	ResponseHandlerMethod ResponseHandlerMethod

	// Hydrate Up
	ConversationTree *ConversationTree `gorm:"-"`

	// Hydrate Down
	Visited       bool                    `gorm:"-"`
	QuickReplies  []*QuickReply           `gorm:"-"`
	ChildrenNodes []*ConversationTreeNode `gorm:"-"`
}

// ConversationTreeNodes is a slice of ConversationTreeNode
type ConversationTreeNodes []ConversationTreeNode

// Save creates if ID = 0 or update id ID > 0
func (o *ConversationTreeNode) Save(conn utils.Conn) error {
	var err error
	if o.ID == 0 {
		err = conn.Create(o).Error
	} else {
		err = conn.Update(o).Error
	}
	if err != nil {
		return err
	}

	ctnqrs := make([]*ConversationTreeNodeQuickReply, len(o.QuickReplyIDs))
	for i, qrID := range o.QuickReplyIDs {
		ctnqr := &ConversationTreeNodeQuickReply{
			ConversationTreeNodeID: o.ID,
			QuickReplyID:           qrID,
		}
		ctnqrs[i] = ctnqr
	}
	if len(ctnqrs) > 0 {
		err = conn.Create(ctnqrs).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate validates an object
func (o *ConversationTreeNode) Validate() error {
	for _, err := range utils.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// AddChildNode adds a new child node to ConversationTreeNode
func (o *ConversationTreeNode) AddChildNode(ctn *ConversationTreeNode) {

	o.ChildrenNodeIDs = append(o.ChildrenNodeIDs, ctn.ID)
	o.ChildrenNodes = append(o.ChildrenNodes, ctn)

	ctn.ConversationTreeID = o.ConversationTreeID
	ctn.ConversationTree = o.ConversationTree
}

// AddQuickReplies adds a quick replies to ConversationTreeNode
func (o *ConversationTreeNode) AddQuickReplies(qrs ...*QuickReply) {

	for _, qr := range qrs {
		o.QuickReplyIDs = append(o.QuickReplyIDs, qr.ID)
		o.QuickReplies = append(o.QuickReplies, qr)

		qr.NodeID = o.ID
	}

}

// SetResponseHandlerMethod sets the responseHandlerMethod for a ConversationTreeNode
func (o *ConversationTreeNode) SetResponseHandlerMethod(method func(utils.Conn, interface{}, string) error) {

	methodName := RegisterMethod(method)

	// Reflect magic to get the function's name
	o.ResponseHandlerMethod.MethodName = methodName

}

// ValidateResponse validates that the response is a valid response
func (o *ConversationTreeNode) ValidateResponse(r string) (bool, string) {

	// If reply can be anything
	if o.ExpectedReplyType == ExpectedReplyTypeAny {
		return true, ""

		// If expected reply has to be an email
	} else if o.ExpectedReplyType == ExpectedReplyTypeEmail {
		var ev utils.EmailValidator
		if ok, _ := ev.Validate(r); !ok {
			return false, "Oh no, the message you sent is not a valid email. Please reply with a valid email."
		}

		// If expected reply is one of the Quick Replies
	} else if o.ExpectedReplyType == ExpectedReplyTypeQuickReply {

		if err := o.ValidateQuickReplyResponse(r, new(QuickReply)); err != nil {
			return false, "Whoops, I didn't understand that. Please select from one of the following quick responses..."
		}

	}

	return true, ""
}

// ValidateQuickReplyResponse validates that the response is a valid response
func (o *ConversationTreeNode) ValidateQuickReplyResponse(r string, qr *QuickReply) error {

	if o.ExpectedReplyType != ExpectedReplyTypeQuickReply {
		return fmt.Errorf("ExpectedReplyType is not ExpectedReplyTypeQuickReply")
	}
	// Iterate through quick replies and find the one expected
	for _, q := range o.QuickReplies {

		if q.ReplyText == r {
			qr = q
			return nil
		}
	}

	return fmt.Errorf("response does not match any of the quick reply responses")

}

// ResponseHandler handles the reponse received using the method defined in validateResponseMethod
func (o *ConversationTreeNode) ResponseHandler(r interface{}, s string) error {

	methodName := o.ResponseHandlerMethod.MethodName
	methodValue := methodMap[methodName]

	errInterface := methodValue.Call([]reflect.Value{reflect.ValueOf(r), reflect.ValueOf(s)})[0].Interface()

	err, ok := errInterface.(error)
	if !ok {
		return nil
	}

	return err

}

// Print prints to console a ConversationTreeNode in a readable manner.
func (o *ConversationTreeNode) Print() {
	fmt.Printf("Node ID: %d\n", o.ID)
	fmt.Printf("\tResponse: %s\n", o.ResponseText)
	fmt.Printf("\tIsRootNode: %v\n", o.IsRootNode)
	fmt.Printf("\tParentNodeResponse: %v\n", o.ParentNodeResponse)
	fmt.Printf("\tResponseHandlerMethod.MethodName: %v\n", o.ResponseHandlerMethod.MethodName)
	if len(o.QuickReplies) > 0 {
		fmt.Printf("\tQuick Replies:\n")
		for i, qr := range o.QuickReplies {
			fmt.Printf("\t\t%d - %s\n", i, qr.ReplyText)
		}
	}
}
