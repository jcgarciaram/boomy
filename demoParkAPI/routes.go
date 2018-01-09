package demoParkAPI

import (
	r "github.com/jcgarciaram/demoPark/demoParkAPI/routes"
)

var allRoutes = r.Routes{

	// Complex
	r.Route{
		Name:        "PostComplex",
		Method:      "POST",
		Pattern:     "/v1/api/demopark/complex",
		HandlerFunc: PostComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutComplex",
		Method:      "PUT",
		Pattern:     "/v1/api/demopark/complex/{complex}",
		HandlerFunc: PutComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetComplex",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/complex/{complex}",
		HandlerFunc: GetComplex,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetComplexes",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/complex",
		HandlerFunc: GetComplexes,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// ParkingDeck
	r.Route{
		Name:        "PostParkingDeck",
		Method:      "POST",
		Pattern:     "/v1/api/demopark/parkingdeck",
		HandlerFunc: PostParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutParkingDeck",
		Method:      "PUT",
		Pattern:     "/v1/api/demopark/parkingdeck/{parkingdeck}",
		HandlerFunc: PutParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingDeck",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/parkingdeck/{parkingdeck}",
		HandlerFunc: GetParkingDeck,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingDecks",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/parkingdeck",
		HandlerFunc: GetParkingDecks,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// ParkingSpace
	r.Route{
		Name:        "PostParkingSpace",
		Method:      "POST",
		Pattern:     "/v1/api/demopark/parkingspace",
		HandlerFunc: PostParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutParkingSpace",
		Method:      "PUT",
		Pattern:     "/v1/api/demopark/parkingspace/{parkingspace}",
		HandlerFunc: PutParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingSpace",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/parkingspace/{parkingspace}",
		HandlerFunc: GetParkingSpace,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetParkingSpaces",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/parkingspace",
		HandlerFunc: GetParkingSpaces,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// Residence
	r.Route{
		Name:        "PostResidence",
		Method:      "POST",
		Pattern:     "/v1/api/demopark/residence",
		HandlerFunc: PostResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutResidence",
		Method:      "PUT",
		Pattern:     "/v1/api/demopark/residence/{residence}",
		HandlerFunc: PutResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidence",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/residence/{residence}",
		HandlerFunc: GetResidence,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidences",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/residence",
		HandlerFunc: GetResidences,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	// Resident
	r.Route{
		Name:        "PostResident",
		Method:      "POST",
		Pattern:     "/v1/api/demopark/resident",
		HandlerFunc: PostResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "PutResident",
		Method:      "PUT",
		Pattern:     "/v1/api/demopark/resident/{resident}",
		HandlerFunc: PutResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResident",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/resident/{resident}",
		HandlerFunc: GetResident,
		VerifyJWT:   false,
		VerifyPerms: true,
	},

	r.Route{
		Name:        "GetResidents",
		Method:      "GET",
		Pattern:     "/v1/api/demopark/resident",
		HandlerFunc: GetResidents,
		VerifyJWT:   false,
		VerifyPerms: true,
	},
}

// GetRoutes returns local variable routes qhich contains all methods for the API
func GetRoutes() r.Routes {
	return allRoutes
}
