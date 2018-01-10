package demoParkAPI

import (
	"github.com/jcgarciaram/demoPark/chatbot"
)

func newUserConversation() {
	ct := chatbot.ConversationTree{
		ConversationType: chatbot.ConversationTypeNewUser,
	}

	node1 := &chatbot.ConversationTreeNode{
		ResponseText:      "Hey there! Excited to get started? Because we are! First thing's first. What's your phone number?",
		ExpectedReplyType: chatbot.ExpectedReplyTypePhoneNumber,
	}

	ct.SetRootNode(node1)

	node2 := &chatbot.ConversationTreeNode{
		ResponseText:      "Awesome! Do you authorize us to send you a quick code through text message? We promise we won't send you text messages without asking your permission first.",
		ExpectedReplyType: chatbot.ExpectedReplyTypeQuickReply,
	}

	node2.AddQuickReplies(
		chatbot.NewQuickReply("Yes"),
		chatbot.NewQuickReply("No"),
	)

	node1.AddChildNode(node2)

	node3 := &chatbot.ConversationTreeNode{
		ResponseText:       "Cool. Check your messages. We just sent you a code. Can you type it below?",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeAny,
		ParentNodeResponse: "Yes",
	}

	node2.AddChildNode(node3)

	node4 := &chatbot.ConversationTreeNode{
		ResponseText:       "Are you sure? We need to confirm the number you gave us belongs to you.",
		ExpectedReplyType:  chatbot.ExpectedReplyTypeQuickReply,
		ParentNodeResponse: "No",
	}

	node4.AddQuickReplies(
		chatbot.NewQuickReply("Yes"),
		chatbot.NewQuickReply("No"),
	)

	node2.AddChildNode(node4)

	node4.AddChildNode(node3)
	node4.AddChildNode(node4)

	// ct.Save()
	// node1.Save()
	// node2.Save()
	// node3.Save()
	// node4.Save()

	ct.Register()
}
