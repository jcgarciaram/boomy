package chatbot

import (
	"log"

	"github.com/jcgarciaram/boomy/dynahelpers"
)

// InitializeConv
func InitializeConv(dynamoInitChan chan struct{}) chan struct{} {
	convInitChan := make(chan struct{})
	go func() {

		<-dynamoInitChan

		// Create all tables
		var ctn ConversationTreeNode
		if err := dynahelpers.CreateTable(ctn); err != nil {
			log.Fatal(err)
		}

		var c Conversation
		if err := dynahelpers.CreateTable(c); err != nil {
			log.Fatal(err)
		}

		var ct ConversationTree
		if err := dynahelpers.CreateTable(ct); err != nil {
			log.Fatal(err)
		}

		var qr QuickReply
		if err := dynahelpers.CreateTable(qr); err != nil {
			log.Fatal(err)
		}

		// Get ConversationTreeNodes
		var ctns ConversationTreeNodes
		if err := ctns.GetAll(); err != nil {
			log.Fatal(err)
		}

		// Get QuickReplies
		var qrs QuickReplies
		if err := qrs.GetAll(); err != nil {
			log.Fatal(err)
		}

		// BuildQuickReply map
		// qrs.buildMap()

		// fmt.Println("Number of quick replies: ", len(quickReplyMap), len(qrs))
		// fmt.Println(qrs)

		BuildConversationTreesFromNodes(ctns, qrs)

		convInitChan <- struct{}{}
	}()

	return convInitChan
}
