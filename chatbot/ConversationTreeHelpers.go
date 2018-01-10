package chatbot

// BuildConversationTrees builds entire conversation trees and is kept in memory for later access
import (
	"fmt"
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
		buildTreeFromNodesHelper(ct.RootNode)
	}

	// PrintConversationTrees()

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

		childNode := convNodeMap[cnID]
		rootNode.ChildrenNodes[i] = childNode

		// If we haven't visited the child node, recursively call this function with it
		if !childNode.Visited {
			buildTreeFromNodesHelper(childNode)
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

func buildTreeFromRootNode(rootNode *ConversationTreeNode) {

	// Ensure we don't visit the same node twice
	rootNode.Visited = true

	// Populate quickReplyMap
	for i, qr := range rootNode.QuickReplies {
		quickReplyMap[qr.GetID()] = rootNode.QuickReplies[i]
	}

	// If we don't have any child nodes, we can return
	if len(rootNode.ChildrenNodes) == 0 {
		return
	}

	// Iterate through child nodes and populate convNodeMap
	for i, cn := range rootNode.ChildrenNodes {

		convNodeMap[cn.GetID()] = rootNode.ChildrenNodes[i]

		// If we haven't visited the child node, recursively call this function with it
		if !cn.Visited {
			buildTreeFromRootNode(cn)
		}

	}

}
