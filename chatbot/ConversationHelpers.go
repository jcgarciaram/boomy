package chatbot

import (
	"fmt"
	"log"
)

// GetResponse queries Dynamo database to get the current state of the conversation
// Then looks at the conversation tree and comes up with a response.
func GetResponse(o interface{}, c *Conversation, message string) (string, []string, error) {

	fmt.Println("GetResponse.Conversation.ConversationTreeID:", c.ConversationTreeID)
	fmt.Println("GetResponse.Conversation.CurrentNodeID:", c.CurrentNodeID)

	// If we don't have a node yet, then we start at the root.
	var currNode *ConversationTreeNode
	if c.CurrentNodeID == 0 {
		currNode = convTreeIDMap[c.ConversationTreeID].RootNode
	} else {
		currNode = convNodeIDMap[c.CurrentNodeID]
	}

	fmt.Printf("\n\ncurrNode: %v\n\n", currNode)

	// If we are expecting a QuickReply
	if currNode.ExpectedReplyType == ExpectedReplyTypeQuickReply {

		// Iterate through the QuickReplies in this node, find the correct one and
		// execute the response handler method if necessary
		var foundReply bool
		for _, qr := range currNode.QuickReplies {
			if qr.ReplyText == message {
				foundReply = true

				if qr.ResponseHandlerMethod.MethodName != "" {
					if err := qr.ResponseHandler(o, message); err != nil {
						log.Fatal("somethig went horribly wrong")
					}
				}
				break
			}
		}
		if !foundReply {
			response := "Whoops. I can only understand one of the options below. Can you please select one?\n\n"
			return response + currNode.ResponseText, QuickReplyStringSlice(currNode.QuickReplies), nil
		}

	} else if currNode.ExpectedReplyType == ExpectedReplyTypeAny {
		if currNode.ResponseHandlerMethod.MethodName != "" {
			if err := currNode.ResponseHandler(o, message); err != nil {
				log.Println(err)
				message = "0"
			} else {
				message = "1"
			}
		}
	}

	fmt.Printf("\n\nmessage: %v\n\n", message)

	var nextNode *ConversationTreeNode
	// If current node does not have any children, reset conversation to root
	if len(currNode.ChildrenNodes) == 0 {

		// Get Conversation Tree
		ct := convTreeIDMap[currNode.ConversationTreeID]
		nextNode = ct.RootNode

	} else {

		// Get next node
		fmt.Println("Iterating through child nodes")
		for _, n := range currNode.ChildrenNodes {

			fmt.Printf("\n\nchildNode: %v\n\n", n)

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

	fmt.Printf("\n\nnextNode: %v\n\n", nextNode)

	// Save new state of conversation
	c.CurrentNodeID = nextNode.ID

	return nextNode.ResponseText, QuickReplyStringSlice(nextNode.QuickReplies), nil

}

// GetCurrentResponse queries Dynamo database to get the current state of the conversation and returns current reponse
func GetCurrentResponse(o interface{}, c *Conversation) (string, []string, error) {

	// If we don't have a node yet, then we start at the root.
	var currNode *ConversationTreeNode
	if c.CurrentNodeID == 0 {
		currNode = convTreeIDMap[c.ConversationTreeID].RootNode
	} else {
		currNode = convNodeIDMap[c.CurrentNodeID]
	}

	return currNode.ResponseText, QuickReplyStringSlice(currNode.QuickReplies), nil
}
