package chatbot

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// InitializeConv builds the conversation tree
func InitializeConv(dbGorm *gorm.DB) chan struct{} {
	db = dbGorm
	convInitChan := make(chan struct{})
	go func() {

		// Create all table
		var ctn ConversationTreeNode
		var c Conversation
		var ct ConversationTree
		var qr QuickReply

		err := db.AutoMigrate(ctn, c, ct, qr).Error
		if err != nil {
			log.Fatal(err)
		}

		// Get ConversationTreeNodes
		var ctns ConversationTreeNodes
		if err := db.Find(&ctns).Error; err != nil {
			log.Fatal(err)
		}

		// Get QuickReplies
		var qrs QuickReplies
		if err := db.Find(&qrs).Error; err != nil {
			log.Fatal(err)
		}

		BuildConversationTreesFromNodes(ctns, qrs)

		convInitChan <- struct{}{}
	}()

	return convInitChan
}
