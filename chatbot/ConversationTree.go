package chatbot

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

var (
	convNodeIDMap       = make(map[uint]*ConversationTreeNode)
	convTreeIDMap       = make(map[uint]*ConversationTree)
	convTreeNicknameMap = make(map[string]*ConversationTree)
	quickReplyMap       = make(map[uint]*QuickReply)
)

// ConversationTree is used internally to build the tree in memory
type ConversationTree struct {
	gorm.Model
	Nickname   string
	RootNodeID uint

	RootNode *ConversationTreeNode `gorm:"-"`
}

// Validate validates an object
func (o *ConversationTree) Validate() error {
	for _, err := range utils.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRootNode set a node as the root node for a conversation tree
func (o *ConversationTree) SetRootNode(ctn *ConversationTreeNode) {

	o.RootNodeID = ctn.ID
	o.RootNode = ctn

	ctn.ConversationTreeID = o.ID
	ctn.ConversationTree = o

	ctn.IsRootNode = true

}

// Register ConversationTree adds conversation tree to map which will be used
func (o *ConversationTree) Register() {
	convTreeIDMap[o.ID] = o
	buildTreeFromRootNode(o.RootNode)
}

// GetBuiltConversationTreeByID queries the map to find conversation tree.
// Returns false if not found
func GetBuiltConversationTreeByID(ID uint) (*ConversationTree, bool) {
	ct, ok := convTreeIDMap[ID]
	return ct, ok
}
