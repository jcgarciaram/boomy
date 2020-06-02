package boomyDB

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// InitializeDB initializes all necessary tables
func InitializeDB(dbGorm *gorm.DB) chan struct{} {
	db = dbGorm
	dbInitChan := make(chan struct{})
	go func() {

		// Create all tables
		var c Complex
		// chatbot.RegisterType(c)

		var pd ParkingDeck
		// chatbot.RegisterType(pd)

		var ps ParkingSpace
		// chatbot.RegisterType(ps)

		var r Residence
		// chatbot.RegisterType(r)

		var rsdnt Resident
		// chatbot.RegisterType(rsdnt)

		db.AutoMigrate(c, pd, ps, r, rsdnt)

		dbInitChan <- struct{}{}
	}()
	return dbInitChan
}
