package boomyAPI

// Message is the structure sent back to the caller in JSON format
type Message struct {
	Message      string
	QuickReplies []string
	Token        string
}
