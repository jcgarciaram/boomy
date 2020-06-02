package boomyAPI

import (
	r "github.com/jcgarciaram/boomy/boomyAPI/routes"
)

var allRoutes = r.Routes{

	// Complex
	r.Route{
		Name:        "PostComplex",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/complex",
		HandlerFunc: PostComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutComplex",
		Method:      "PUT",
		Pattern:     "/v1/api/boomy/complex/{complex}",
		HandlerFunc: PutComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetComplex",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/complex/{complex}",
		HandlerFunc: GetComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetComplexes",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/complex",
		HandlerFunc: GetComplexes,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// ParkingDeck
	r.Route{
		Name:        "PostParkingDeck",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/parkingdeck",
		HandlerFunc: PostParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutParkingDeck",
		Method:      "PUT",
		Pattern:     "/v1/api/boomy/parkingdeck/{parkingdeck}",
		HandlerFunc: PutParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingDeck",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/parkingdeck/{parkingdeck}",
		HandlerFunc: GetParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingDecks",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/parkingdeck",
		HandlerFunc: GetParkingDecks,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// ParkingSpace
	r.Route{
		Name:        "PostParkingSpace",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/parkingspace",
		HandlerFunc: PostParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutParkingSpace",
		Method:      "PUT",
		Pattern:     "/v1/api/boomy/parkingspace/{parkingspace}",
		HandlerFunc: PutParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingSpace",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/parkingspace/{parkingspace}",
		HandlerFunc: GetParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingSpaces",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/parkingspace",
		HandlerFunc: GetParkingSpaces,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// Residence
	r.Route{
		Name:        "PostResidence",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/residence",
		HandlerFunc: PostResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutResidence",
		Method:      "PUT",
		Pattern:     "/v1/api/boomy/residence/{residence}",
		HandlerFunc: PutResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidence",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/residence/{residence}",
		HandlerFunc: GetResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidences",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/residence",
		HandlerFunc: GetResidences,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// Resident
	r.Route{
		Name:        "PostResident",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/resident",
		HandlerFunc: PostResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutResident",
		Method:      "PUT",
		Pattern:     "/v1/api/boomy/resident/{resident}",
		HandlerFunc: PutResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResident",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/resident/{resident}",
		HandlerFunc: GetResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidents",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/resident",
		HandlerFunc: GetResidents,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "ResidentPostMessage",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/resident/message",
		HandlerFunc: ResidentPostMessage,
		VerifyJWT:   true,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "ResidentBeginConversation",
		Method:      "POST",
		Pattern:     "/v1/api/boomy/resident/conversation",
		HandlerFunc: ResidentBeginConversation,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "ResidentGetConversation",
		Method:      "GET",
		Pattern:     "/v1/api/boomy/conversation/resident",
		HandlerFunc: ResidentGetConversation,
		VerifyJWT:   true,
		VerifyPerms: true,
	},
}

// GetRoutes returns local variable routes qhich contains all methods for the API
func GetRoutes() r.Routes {
	return allRoutes
}
