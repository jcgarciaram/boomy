package demoParkDB

import (
	"log"

	"github.com/jcgarciaram/demoPark/dynahelpers"
)

// InitializeDB
func InitializeDB(dynamoInitChan chan struct{}) chan struct{} {
	dbInitChan := make(chan struct{})
	go func() {

		<-dynamoInitChan

		// Create all tables
		var c Complex
		if err := dynahelpers.CreateTable(c); err != nil {
			log.Fatal(err)
		}

		var pd ParkingDeck
		if err := dynahelpers.CreateTable(pd); err != nil {
			log.Fatal(err)
		}

		var ps ParkingSpace
		if err := dynahelpers.CreateTable(ps); err != nil {
			log.Fatal(err)
		}

		var r Residence
		if err := dynahelpers.CreateTable(r); err != nil {
			log.Fatal(err)
		}

		var rsdnt Resident
		if err := dynahelpers.CreateTable(rsdnt); err != nil {
			log.Fatal(err)
		}

		dbInitChan <- struct{}{}
	}()
	return dbInitChan
}
