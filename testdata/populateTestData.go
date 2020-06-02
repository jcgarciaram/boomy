package main

import (
	"github.com/jcgarciaram/boomy/boomyDB"
)

func main() {
	c := boomyDB.Complex{
		Name:         "Pipo's Place",
		AddressLine1: "800 Pipo Street",
		AddressLine2: "Apt. 1409",
		City:         "Caguas",
		State:        "Puerto Rico",
		ZipCode:      "00725",
	}

	pd1 := boomyDB.ParkingDeck{
		ComplexID: c.GetID(),
		Name:      "Deck A",
		NumFloors: 5,
		NumSpaces: 100,
	}

	pd2 := boomyDB.ParkingDeck{
		ComplexID: c.GetID(),
		Name:      "Deck A",
		NumFloors: 5,
		NumSpaces: 100,
	}

}
