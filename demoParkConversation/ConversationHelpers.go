package demoParkConversation

import "fmt"

// getResponse queries Dynamo database to get the current state of the conversation
// Then looks at the conversation tree and comes up with a response.
func getResponse(senderId string, message string) (string, []*QuickReply, error) {

	// Get conversation
	var c Conversation
	if err := c.Get(senderId); err != nil && err.Error() != "dynamo: no item found" {
		return "", nil, err
	}

	// If conversation doesn't exit, user has never registered. Begin new user conversation
	if c.SenderID == "" {
		fmt.Println("New user!")

		c.SenderID = senderId

		// Get Conversation Tree by type
		var ct2 ConversationTree
		if err := ct2.GetOneByField("ConversationType", ConversationTypeNewUser); err != nil {
			return "", nil, err
		}

		// Get conversation tree map from map
		ct := convTreeMap[ct2.ID]

		// Save conversation
		c.ConversationTreeID = ct.ID
		c.CurrentNodeID = ct.RootNode.ID
		if err := c.Save(); err != nil {
			return "", nil, err
		}

		textResponse := ct.RootNode.ResponseText
		quickReplies := ct.RootNode.QuickReplies

		return textResponse, quickReplies, nil
	}

	// Because we already have a conversation started with this recipient, we can analyze the message they've sent and come up with our response
	currNode := convNodeMap[c.CurrentNodeID]

	// fmt.Printf("\n\ncurrNode: %v\n\n", currNode)

	// Validate response
	if ok, mess := currNode.ValidateResponse(message); !ok {
		return mess, currNode.QuickReplies, nil
	}

	var nextNode *ConversationTreeNode
	// If current node does not have any children, reset conversation to root
	if len(currNode.ChildrenNodes) == 0 {

		// Get Conversation Tree
		ct := convTreeMap[currNode.ConversationTreeID]
		nextNode = ct.RootNode

	} else {

		// Get next node
		// fmt.Println("Iterating through child nodes")
		for _, n := range currNode.ChildrenNodes {

			// fmt.Printf("\n\nchildNode: %v\n\n", n)

			if n.ParentNodeResponse == nil {
				nextNode = n
				break
			}

			if n.ParentNodeResponse == message {
				nextNode = n
				break
			}
		}
	}

	if nextNode == nil {
		return "Something bad happened. Help me...", nil, nil
	}

	// Save new state of conversation
	c.CurrentNodeID = nextNode.GetID()
	if err := c.Save(); err != nil {
		return "", nil, err
	}

	return nextNode.ResponseText, nextNode.QuickReplies, nil

}
