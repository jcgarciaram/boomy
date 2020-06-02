package chatbot

import (
	"github.com/jcgarciaram/boomy/utils"
	"github.com/jinzhu/gorm"
)

// Conversation is a single conversation with a recipient
type Conversation struct {
	gorm.Model
	ConversationTreeID uint
	CurrentNodeID      uint
}

// Conversations is a slice of Conversation
type Conversations []Conversation

// Validate validates an object
func (o *Conversation) Validate() error {
	for _, err := range utils.ValidateStruct(*o) {
		if err != nil {
			return err
		}
	}
	return nil
}
