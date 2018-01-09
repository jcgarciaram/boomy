package demoParkConversation

import (
	"fmt"
	"log"

	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
	uuid "github.com/satori/go.uuid"
)

// ConversationTreeNode is saved in Dynamo. Used by ConversationTree to build up the tree in memory
type ConversationTreeNode struct {
	ID                   string      `dynamo:"ID,hash"`
	ConversationTreeID   string      `dynamo:"ConversationTreeID"`
	ResponseText         string      `dynamo:"ResponseText"`
	ExpectedReplyType    int         `dynamo:"ExpectedReplyType"`
	QuickReplyIDs        []string    `dynamo:"QuickReplyIDs"`
	ChildrenNodeIDs      []string    `dynamo:"ChildrenNodeIDs"`
	DayIntervalIncrement int         `dynamo:"DayIntervalIncrement"`
	IsRootNode           bool        `dynamo:"IsRootNode"`
	ParentNodeResponse   interface{} `dynamo:"ParentNodeResponse"`

	// We can store a struct that will contain the name of the type that contains the function and the name of function itself in strings
	// We then need another function that has a default variable of all necessary types and depending on the type uses that default variable and reflection below to call the method. All methods should return bool
	// See the following example:
	//
	// https://play.golang.org/p/no9z-0wvkmy
	//
	// As part of initializing the dynamoDB package, we can add all the types to a map that will contain the reflect.Type of the type. This map needs to be visible to demoParkConversation
	//
	// https://play.golang.org/p/2TugHjN1IUk
	//
	validateResponseMethod

	// Hydrate Up
	ConversationTree *ConversationTree `dynamo:"-"`

	// Hydrate Down
	Visited       bool                    `dynamo:"-"`
	QuickReplies  []*QuickReply           `dynamo:"-"`
	ChildrenNodes []*ConversationTreeNode `dynamo:"-"`
}

// ConversationTreeNodes is a slice of ConversationTreeNode
type ConversationTreeNodes []ConversationTreeNode

// Save puts struct o in Dynamo
func (o *ConversationTreeNode) Save() error {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	if err := dynahelpers.DynamoPut(o); err != nil {
		log.Printf("Error saving object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Get gets a struct from Dynamo and unmarshals results into o
func (o *ConversationTreeNode) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// Validate validates an object
func (o *ConversationTreeNode) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *ConversationTreeNode) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// GetAll gets all records from Dynamo and unmarshals results into o
func (oSlice *ConversationTreeNodes) GetAll() error {
	var o ConversationTreeNode
	if err := dynahelpers.DynamoGetAll(o, oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}

	if len(*oSlice) == 0 {
		*oSlice = make([]ConversationTreeNode, 0)
	}

	return nil
}

// AddChildNode adds a new child node to ConversationTreeNode
func (o *ConversationTreeNode) AddChildNode(ctn *ConversationTreeNode) {

	o.ChildrenNodeIDs = append(o.ChildrenNodeIDs, ctn.GetID())
	o.ChildrenNodes = append(o.ChildrenNodes, ctn)

	ctn.ConversationTreeID = o.ConversationTreeID
	ctn.ConversationTree = o.ConversationTree
}

// AddQuickReplies adds a quick replies to ConversationTreeNode
func (o *ConversationTreeNode) AddQuickReplies(qrs ...*QuickReply) {

	for _, qr := range qrs {
		o.QuickReplyIDs = append(o.QuickReplyIDs, qr.GetID())
		o.QuickReplies = append(o.QuickReplies, qr)

		qr.NodeID = o.GetID()
	}

}

// ValidateResponse validates that the response is a valid response
func (o *ConversationTreeNode) ValidateResponse(r string) (bool, string) {

	// If reply can be anything
	if o.ExpectedReplyType == ExpectedReplyTypeAny {
		return true, ""

		// If expected reply has to be an email
	} else if o.ExpectedReplyType == ExpectedReplyTypeEmail {
		var ev dynahelpers.EmailValidator
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

// Print prints to console a ConversationTreeNode in a readable manner.
func (o *ConversationTreeNode) Print() {
	fmt.Printf("Node ID: %s\n", o.ID)
	fmt.Printf("\tResponse: %s\n", o.ResponseText)
	fmt.Printf("\tIsRootNode: %v\n", o.IsRootNode)
	fmt.Printf("\tParentNodeResponse: %v\n", o.ParentNodeResponse)
	fmt.Printf("\tvalidateResponseMethod.typeName: %v\n", o.validateResponseMethod.typeName)
	fmt.Printf("\tvalidateResponseMethod.methodName: %v\n", o.validateResponseMethod.methodName)
	if len(o.QuickReplies) > 0 {
		fmt.Printf("\tQuick Replies:\n")
		for i, qr := range o.QuickReplies {
			fmt.Printf("\t\t%d - %s\n", i, qr.ReplyText)
		}
	}
}
