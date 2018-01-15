package chatbot

import (
	"fmt"
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/dynahelpers"
	"github.com/jcgarciaram/boomy/utils"
	uuid "github.com/satori/go.uuid"
)

var (
	convNodeMap   = make(map[string]*ConversationTreeNode)
	convTreeMap   = make(map[string]*ConversationTree)
	quickReplyMap = make(map[string]*QuickReply)
)

// ConversationTree is used internally to build the tree in memory
type ConversationTree struct {
	ID               string `dynamo:"ID,hash"`
	ConversationType int    `dynamo:"ConversationType"`
	RootNodeID       string `dynamo:"RootNodeID"`

	RootNode *ConversationTreeNode `dynamo:"-"`
}

// Save puts struct o in Dynamo
func (o *ConversationTree) Save() error {
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
func (o *ConversationTree) Get(ID string) error {
	if err := dynahelpers.DynamoGet(ID, o); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}
	return nil
}

// GetID gets a struct from Dynamo and unmarshals results into o
func (o *ConversationTree) GetID() string {
	if o.ID == "" {
		o.ID = uuid.NewV4().String()
	}

	return o.ID
}

// GetOneByField searches for a struct from Dynamo and unmarshals results into o. Should only return one value
func (o *ConversationTree) GetOneByField(fieldName string, value interface{}) error {

	var oSlice []ConversationTree
	if err := dynahelpers.DynamoGetByField(fieldName, value, o, &oSlice); err != nil {
		log.Printf("Error getting object of type %s\n", utils.GetType(o))
		return err
	}

	if len(oSlice) > 1 {
		logrus.WithFields(logrus.Fields{
			"fieldName": fieldName,
			"value":     value,
			"tableName": utils.GetType(o),
		}).Warn("More than one field returned from Dynamo")

		return fmt.Errorf("More than one field returned from Dynamo")
	}

	fmt.Println("oSlice:", oSlice)

	*o = oSlice[0]

	fmt.Println(o.ID)

	return nil
}

// Validate validates an object
func (o *ConversationTree) Validate() error {
	for _, err := range dynahelpers.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRootNode set a node as the root node for a conversation tree
func (o *ConversationTree) SetRootNode(ctn *ConversationTreeNode) {

	o.RootNodeID = ctn.GetID()
	o.RootNode = ctn

	ctn.ConversationTreeID = o.GetID()
	ctn.ConversationTree = o

	ctn.IsRootNode = true

}

// Register ConversationTree adds conversation tree to map which will be used
func (o *ConversationTree) Register() {
	convTreeMap[o.GetID()] = o

	fmt.Println("convTreeMap:", convTreeMap)

	buildTreeFromRootNode(o.RootNode)
}
