package boomyAPI

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

// Initialize initializes all necessary tables
func Initialize(dbGorm *gorm.DB) {
	db = dbGorm
	newUserConvTree := buildNewUserConversation(db)
	newUserConvTree.Register()
}
