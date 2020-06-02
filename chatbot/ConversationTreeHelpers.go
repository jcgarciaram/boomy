package chatbot

// BuildConversationTrees builds entire conversation trees and is kept in memory for later access
import (
	"fmt"
	"log"
)

// BuildConversationTreesFromNodes builds entire conversation tree from nodes and quick replies passed
func BuildConversationTreesFromNodes(ctns ConversationTreeNodes, qrs QuickReplies) {

	// builds quickReplyMap which helps when building the tree
	for i, qr := range qrs {
		quickReplyMap[qr.ID] = &qrs[i]
	}

	// Iterate through all nodes, find root of tree, build maps
	for i, ctn := range ctns {

		// Add ConversationTreeNode (ctns) to convNodeMap
		convNodeIDMap[ctn.ID] = &ctns[i]

		// Register method if
		// Get conversationTree from convTreeMap. If not there, create
		ct, ok := convTreeIDMap[ctn.ConversationTreeID]
		if !ok {
			ct := ConversationTree{}
			fmt.Println("ctn.ConversationTreeID", ctn.ConversationTreeID)
			if err := db.First(&ct, ctn.ConversationTreeID); err != nil {
				log.Fatal("Error getting conversation tree:", err)
			}

			convTreeIDMap[ctn.ConversationTreeID] = &ct
			convTreeNicknameMap[ct.Nickname] = &ct
		}

		if ctn.IsRootNode {
			ct.RootNode = &ctns[i]
		}

	}

	// Iterate through all conversation trees and build them
	for _, ct := range convTreeIDMap {
		buildTreeFromNodesHelper(ct.RootNode)
	}
}

func buildTreeFromNodesHelper(rootNode *ConversationTreeNode) {

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

		childNode := convNodeIDMap[cnID]
		rootNode.ChildrenNodes[i] = childNode

		// If we haven't visited the child node, recursively call this function with it
		if !childNode.Visited {
			buildTreeFromNodesHelper(childNode)
		}

	}

}

// PrintConversationTrees iterates through a tree and prints out the tree using DFS
func PrintConversationTrees() {
	for ID, tree := range convTreeIDMap {

		fmt.Printf("\n\nPrinting tree: %d\n\n", ID)

		parentChildNodeMap := make(map[uint]map[uint]struct{})
		printTree(tree.RootNode, parentChildNodeMap)
	}
}

// PrintConversationTrees iterates through a tree and prints out the tree using DFS
func printTree(n *ConversationTreeNode, parentChildNodeMap map[uint]map[uint]struct{}) {

	n.Print()

	// If we don't have any child nodes, we can return
	if len(n.ChildrenNodes) == 0 {
		fmt.Printf("\tNo children nodes\n\n")
		return
	}

	fmt.Printf("\t%d children nodes:\n\n", len(n.ChildrenNodes))

	for _, cn := range n.ChildrenNodes {

		if innerMap, ok := parentChildNodeMap[n.ID]; !ok {
			innerMap = make(map[uint]struct{})
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

func buildTreeFromRootNode(rootNode *ConversationTreeNode) {

	convNodeIDMap[rootNode.ID] = rootNode

	// Ensure we don't visit the same node twice
	rootNode.Visited = true

	// Populate quickReplyMap
	for i, qr := range rootNode.QuickReplies {
		quickReplyMap[qr.ID] = rootNode.QuickReplies[i]
	}

	// If we don't have any child nodes, we can return
	if len(rootNode.ChildrenNodes) == 0 {
		return
	}

	// Iterate through child nodes if we haven't visited it, recursively call this function with it
	for _, cn := range rootNode.ChildrenNodes {
		if !cn.Visited {
			buildTreeFromRootNode(cn)
		}
	}

}

func saveTreeFromRootNode(rootNode *ConversationTreeNode) {

	fmt.Printf("\n\nrootNode: %v\n\n", *rootNode)

	if err := rootNode.Save(db); err != nil {
		log.Fatal("Error saving node:", err)
	}

	// Ensure we don't visit the same node twice
	rootNode.Visited = true

	// Save Quick Replies
	for _, qr := range rootNode.QuickReplies {
		if err := qr.Save(db); err != nil {
			log.Fatal("Error saving Quick Reply:", err)
		}
	}

	// If we don't have any child nodes, we can return
	if len(rootNode.ChildrenNodes) == 0 {
		return
	}

	// Iterate through child nodes if we haven't visited it, recursively call this function with it
	for _, cn := range rootNode.ChildrenNodes {
		if !cn.Visited {
			saveTreeFromRootNode(cn)
		}
	}

}

// GetHydratedConversationTree gets the tree from convTreeIDMap
func GetHydratedConversationTree(ID uint) (*ConversationTree, error) {
	ct, ok := convTreeIDMap[ID]
	if !ok {
		return nil, fmt.Errorf("Converation Tree with ID %d does not exist", ID)
	}
	return ct, nil
}
