package chatbot

func BuildNewUserConversation() error {
	ct := ConversationTree{
		ConversationType: ConversationTypeNewUser,
	}

	node1 := &ConversationTreeNode{
		ResponseText:      "Hey there! Excited to get started? Because we are! First thing's first. What's your phone number?",
		ExpectedReplyType: ExpectedReplyTypePhoneNumber,
	}

	ct.SetRootNode(node1)

	node2 := &ConversationTreeNode{
		ResponseText:      "Awesome! Do you authorize us to send you a quick code through text message? We promise we won't send you text messages without asking your permission first.",
		ExpectedReplyType: ExpectedReplyTypeQuickReply,
	}

	node2.AddQuickReplies(
		NewQuickReply("Yes"),
		NewQuickReply("No"),
	)

	node1.AddChildNode(node2)

	node3 := &ConversationTreeNode{
		ResponseText:       "Cool. Check your messages. We just sent you a code. Can you type it below?",
		ExpectedReplyType:  ExpectedReplyTypeAny,
		ParentNodeResponse: "Yes",
	}

	node2.AddChildNode(node3)

	node4 := &ConversationTreeNode{
		ResponseText:       "Are you sure? We need to confirm the number you gave us belongs to you.",
		ExpectedReplyType:  ExpectedReplyTypeQuickReply,
		ParentNodeResponse: "No",
	}

	node4.AddQuickReplies(
		NewQuickReply("Yes"),
		NewQuickReply("No"),
	)

	node2.AddChildNode(node4)

	node4.AddChildNode(node3)
	node4.AddChildNode(node4)

	ct.Save()
	node1.Save()
	node2.Save()
	node3.Save()
	node4.Save()

	return nil
}
