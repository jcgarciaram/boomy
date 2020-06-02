package boomyAPI

import (
	"fmt"

	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jcgarciaram/boomy/utils"
)

func buildNewUserConversation(conn utils.Conn) *chatbot.ConversationTree {

	newUserConvTree := &chatbot.ConversationTree{}

	//TODO: look for error here
	conn.First(newUserConvTree, ConversationTreeIDNewResident)
	if newUserConvTree.ID != 0 {
		return newUserConvTree
	}

	// newUserConvTree.ID = ConversationTreeIDNewResident
	newUserConvTree.Nickname = ConversationTreeNicknameNewResident

	// Declare Root node
	node1 := &chatbot.ConversationTreeNode{
		ResponseText:      "Hey there! Excited to get started? Because we are!\n\nFirst thing's first. We need to setup your profile. I'm going to ask you some questions. Just reply with your answer. If you need to fix your any of previous answers you can just type \"FIX\". Got it?",
		ExpectedReplyType: chatbot.ExpectedReplyTypeQuickReply,
	}

	node1.AddQuickReplies(
		chatbot.NewQuickReply(conn, "Yes", nil),
		chatbot.NewQuickReply(conn, "No", nil),
	)

	// First Name
	node2 := &chatbot.ConversationTreeNode{
		ResponseText:       "Awesome! What's your first name?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "Yes",
	}
	node2.SetResponseHandlerMethod(ResidentSetFirstName)

	// Are you sure?
	node3 := &chatbot.ConversationTreeNode{
		ResponseText:       "No? Umm... I kind of need to your information. Swear I won't do anything weird with it. It's like when you signed up for Facebook. They asked you questions right? So, you ready?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeQuickReply,
		ParentNodeResponse: "No",
	}

	node3.AddQuickReplies(
		chatbot.NewQuickReply(conn, "Yes", nil),
		chatbot.NewQuickReply(conn, "No", nil),
	)

	// Last Name
	node4 := &chatbot.ConversationTreeNode{
		ResponseText:       "And your last name?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "1",
	}

	node4.SetResponseHandlerMethod(ResidentSetLastName)

	// Phone Number
	node5 := &chatbot.ConversationTreeNode{
		ResponseText:       "Nice to meet you! And finally, what's your cell phone number where you can receive text messages?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "1",
	}

	node5.SetResponseHandlerMethod(ResidentSendValidationCode)

	// Invalid Phone
	node6 := &chatbot.ConversationTreeNode{
		ResponseText:       "Uh oh. It seems that phone number was not valid. Can you double-check and give me your number one more time?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "0",
	}

	node6.SetResponseHandlerMethod(ResidentSendValidationCode)

	// Sent code
	node7 := &chatbot.ConversationTreeNode{
		ResponseText:       "Cool! We sent you a code via text message. Can you type it in?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "1",
	}
	node7.SetResponseHandlerMethod(ResidentCheckValidationCode)

	// Invalid code
	node8 := &chatbot.ConversationTreeNode{
		ResponseText:       "Hmm... That doesn't look like the code we sent you. What would you like to do?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeQuickReply,
		ParentNodeResponse: "0",
	}

	node8.AddQuickReplies(
		chatbot.NewQuickReply(conn, "Try again", nil),
		chatbot.NewQuickReply(conn, "Send me a new code", ResidentRegenerateValidationCode),
	)

	// Sent code - again
	node9 := &chatbot.ConversationTreeNode{
		ResponseText:       "OK no prob! We generated a new code and sent it via text message. Can you type it in?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "Send me a new code",
	}
	node9.SetResponseHandlerMethod(ResidentCheckValidationCode)

	// Try again
	node10 := &chatbot.ConversationTreeNode{
		ResponseText:       "Alrighty! When you're ready retype the code we sent you...",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "Try again",
	}
	node10.SetResponseHandlerMethod(ResidentCheckValidationCode)

	// Success!!
	node11 := &chatbot.ConversationTreeNode{
		ResponseText:       "Success!! You have been registered.",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "1",
	}

	//TODO: look for errors here
	conn.Create(newUserConvTree)

	// Set root node
	newUserConvTree.SetRootNode(node1)
	conn.Save(newUserConvTree)
	conn.Save(node1)

	conn.Create(node1)
	conn.Create(node2)
	conn.Create(node3)
	conn.Create(node4)
	conn.Create(node5)
	conn.Create(node6)
	conn.Create(node7)
	conn.Create(node8)
	conn.Create(node9)
	conn.Create(node10)
	conn.Create(node11)

	node1.AddChildNode(node2)
	node1.AddChildNode(node3)

	node2.AddChildNode(node4)

	node3.AddChildNode(node2)
	node3.AddChildNode(node3)

	node4.AddChildNode(node5)

	node5.AddChildNode(node6)
	node5.AddChildNode(node7)

	node6.AddChildNode(node6)
	node6.AddChildNode(node7)

	node7.AddChildNode(node8)
	node7.AddChildNode(node11)

	node8.AddChildNode(node9)
	node8.AddChildNode(node10)

	node9.AddChildNode(node8)
	node9.AddChildNode(node11)

	node10.AddChildNode(node8)
	node10.AddChildNode(node11)

	// Update conv tree with root node
	fmt.Println("updating with root node ID", conn.Save(newUserConvTree).Error, newUserConvTree.RootNodeID)

	return newUserConvTree
}
