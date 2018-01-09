package demoParkAPI

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	apih "github.com/jcgarciaram/demoPark/demoParkAPI/APIHelpers"
	"github.com/jcgarciaram/demoPark/demoParkDB"
	"github.com/jcgarciaram/demoPark/utils"
)

// PostParkingSpace creates a new ParkingSpace
func PostParkingSpace(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal body of request into
	var o demoParkDB.ParkingSpace

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.PostHelper(r, &o); err != nil {
		http.Error(w, fmt.Sprintf("There was an error saving the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(o); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// PutParkingSpace updates an existing ParkingSpace
func PutParkingSpace(w http.ResponseWriter, r *http.Request) {

	// Get the variables from the request
	vars := mux.Vars(r)
	ID := vars["parkingspace"]

	// Struct to unmarshal body of request into
	var o demoParkDB.ParkingSpace

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.PutHelper(r, &o, ID); err != nil {
		http.Error(w, fmt.Sprintf("There was an updating the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(o); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// GetParkingSpace gets a specific ParkingSpace
func GetParkingSpace(w http.ResponseWriter, r *http.Request) {

	// Get the variables from the request
	vars := mux.Vars(r)
	ID := vars["parkingspace"]

	// Struct to unmarshal result from Dynamo into
	var o demoParkDB.ParkingSpace

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.GetHelper(r, &o, ID); err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(o); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// GetParkingSpaces gets all ParkingSpaces from Dynamo
func GetParkingSpaces(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o demoParkDB.ParkingSpaces

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.GetAllHelper(r, &o); err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(o); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}
