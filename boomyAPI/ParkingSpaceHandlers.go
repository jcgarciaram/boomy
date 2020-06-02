package boomyAPI

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	apih "github.com/jcgarciaram/boomy/boomyAPI/APIHelpers"
	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/utils"
)

// PostParkingSpace creates a new parkingSpace
func PostParkingSpace(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Struct to unmarshal body of request into
	var o boomyDB.ParkingSpace

	tx := db.BeginTx(context.Background(), &sql.TxOptions{})
	if tx.Error != nil {
		http.Error(w, fmt.Sprintf("There was an error initializing transaction. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	if err := apih.PostHelper(r, &o, tx); err != nil {
		http.Error(w, fmt.Sprintf("There was an error saving the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	if tx.Commit().Error != nil {
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

// PutParkingSpace updates an existing parkingSpace
func PutParkingSpace(w http.ResponseWriter, r *http.Request) {

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Struct to unmarshal body of request into
	var o boomyDB.ParkingSpace

	// Get the variables from the request
	vars := mux.Vars(r)
	IDStr := vars["parkingSpace"]
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID passed in request. Please contact your administrator.\n"), http.StatusBadRequest)
		return
	}

	tx := db.BeginTx(context.Background(), &sql.TxOptions{})
	if tx.Error != nil {
		http.Error(w, fmt.Sprintf("There was an error initializing transaction. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	if err := apih.PutHelper(r, &o, uint(ID), tx); err != nil {
		http.Error(w, fmt.Sprintf("There was an updating the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	if tx.Commit().Error != nil {
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

// GetParkingSpace gets a specific parkingSpace
func GetParkingSpace(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.ParkingSpace

	// Get the variables from the request
	vars := mux.Vars(r)
	IDStr := vars["parkingSpace"]
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID passed in request. Please contact your administrator.\n"), http.StatusBadRequest)
		return
	}

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.GetHelper(r, &o, uint(ID), db); err != nil {
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

// GetParkingSpaces gets all parking spaces from Dynamo
func GetParkingSpaces(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.ParkingSpaces

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := apih.GetAllHelper(r, &o, db); err != nil {
		http.Error(w, fmt.Sprintf("There was an error getting the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(o); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}
