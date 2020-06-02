package chatbot

// ConversationTreeNodeQuickReply is the bridge table between
// a ConversationTreeNode and a QuickReply
type ConversationTreeNodeQuickReply struct {
	ConversationTreeNodeID uint
	QuickReplyID           uint
}
