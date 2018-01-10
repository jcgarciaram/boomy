package demoParkDB

import (
	"log"

	"github.com/jcgarciaram/demoPark/chatbot"
	"github.com/jcgarciaram/demoPark/dynahelpers"
)

// InitializeDB initializes all necessary tables
func InitializeDB(dynamoInitChan chan struct{}) chan struct{} {
	dbInitChan := make(chan struct{})
	go func() {

		<-dynamoInitChan

		// Create all tables
		var c Complex
		if err := dynahelpers.CreateTable(c); err != nil {
			log.Fatal(err)
		}
		chatbot.RegisterType(c)

		var pd ParkingDeck
		if err := dynahelpers.CreateTable(pd); err != nil {
			log.Fatal(err)
		}
		chatbot.RegisterType(pd)

		var ps ParkingSpace
		if err := dynahelpers.CreateTable(ps); err != nil {
			log.Fatal(err)
		}
		chatbot.RegisterType(ps)

		var r Residence
		if err := dynahelpers.CreateTable(r); err != nil {
			log.Fatal(err)
		}
		chatbot.RegisterType(r)

		var rsdnt Resident
		if err := dynahelpers.CreateTable(rsdnt); err != nil {
			log.Fatal(err)
		}
		chatbot.RegisterType(rsdnt)

		dbInitChan <- struct{}{}
	}()
	return dbInitChan
}
