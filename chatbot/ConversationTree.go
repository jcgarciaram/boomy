package chatbot

import (
	"fmt"
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/demoPark/dynahelpers"
	"github.com/jcgarciaram/demoPark/utils"
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

// BuildConversationTrees builds entire conversation trees and is kept in memory for later access
func BuildConversationTrees(ctns ConversationTreeNodes) {

	// Iterate through all nodes, find root of tree, build maps
	for i, ctn := range ctns {

		// Add ConversationTreeNode (ctns) to convNodeMap
		convNodeMap[ctn.ID] = &ctns[i]

		// Get conversationTree from convTreeMap. If not there, create
		ct, ok := convTreeMap[ctn.ConversationTreeID]
		if !ok {
			ct = &ConversationTree{}
			convTreeMap[ctn.ConversationTreeID] = ct
		}

		if ctn.IsRootNode {
			ct.RootNode = &ctns[i]
		}

	}

	// Iterate through all conversation trees and build them
	for _, ct := range convTreeMap {
		buildTreeHelper(ct.RootNode)
	}

	// PrintConversationTrees()

}

func buildTreeHelper(rootNode *ConversationTreeNode) {

	// Ensure we don't visit the same node twice
	rootNode.Visited = true

	// Populate QuickReplies slice
	rootNode.QuickReplies = make([]*QuickReply, len(rootNode.QuickReplyIDs))
	for i, qrID := range rootNode.QuickReplyIDs {
		rootNode.QuickReplies[i] = quickReplyMap[qrID]
	}

	// If we don't have any child nodes, we can return
	if len(rootNode.ChildrenNodeIDs) == 0 {
		return
	}

	// Iterate through child node IDS and populate ChildresNodes
	rootNode.ChildrenNodes = make([]*ConversationTreeNode, len(rootNode.ChildrenNodeIDs))
	for i, cnID := range rootNode.ChildrenNodeIDs {

		childNode := convNodeMap[cnID]
		rootNode.ChildrenNodes[i] = childNode

		// If we haven't visited the child node, recursively call this function with it
		if !childNode.Visited {
			buildTreeHelper(childNode)
		}

	}

}

// PrintConversationTrees iterates through a tree and prints out the tree using DFS
func PrintConversationTrees() {
	for ID, tree := range convTreeMap {

		fmt.Printf("\n\nPrinting tree: %s\n\n", ID)

		parentChildNodeMap := make(map[string]map[string]struct{})
		printTree(tree.RootNode, parentChildNodeMap)
	}
}

// PrintConversationTrees iterates through a tree and prints out the tree using DFS
func printTree(n *ConversationTreeNode, parentChildNodeMap map[string]map[string]struct{}) {

	n.Print()

	// If we don't have any child nodes, we can return
	if len(n.ChildrenNodes) == 0 {
		fmt.Printf("\tNo children nodes\n\n")
		return
	}

	fmt.Printf("\t%d children nodes:\n\n", len(n.ChildrenNodes))

	for _, cn := range n.ChildrenNodes {

		if innerMap, ok := parentChildNodeMap[n.ID]; !ok {
			innerMap = make(map[string]struct{})
			innerMap[cn.ID] = struct{}{}

			parentChildNodeMap[n.ID] = innerMap

		} else if _, ok = innerMap[cn.ID]; !ok {
			innerMap[cn.ID] = struct{}{}

			parentChildNodeMap[n.ID] = innerMap
		} else {
			continue
		}

		printTree(cn, parentChildNodeMap)

	}
}
