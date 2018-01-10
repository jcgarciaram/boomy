package chatbot

const (
	// ExpectedReplyTypeAny refers to a reply by the user where it is something they type
	ExpectedReplyTypeAny = 1

	// ExpectedReplyTypeQuickReply refers to a reply by the user where it is one of the pre-selected options
	ExpectedReplyTypeQuickReply = 2

	// ExpectedReplyTypeEmail refers to a reply that contains an email
	ExpectedReplyTypeEmail = 3

	// ExpectedReplyTypePhoneNumber expects a phone number in the message
	ExpectedReplyTypePhoneNumber = 5
)
