package boomyAPI

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	apih "github.com/jcgarciaram/boomy/boomyAPI/APIHelpers"
	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/chatbot"
	"github.com/jcgarciaram/boomy/utils"
)

// ResidentBeginConversation creates a new Resident and begins a new conversation with them
func ResidentBeginConversation(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	o.FirstName = "Juan"

	newUserConvTree, ok := chatbot.GetBuiltConversationTreeByID(ConversationTreeIDNewResident)
	if !ok {
		http.Error(w, fmt.Sprintf("New resident conversation not registered. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	o.Conversation = chatbot.Conversation{
		ConversationTreeID: ConversationTreeIDNewResident,
	}

	tx := db.BeginTx(context.Background(), &sql.TxOptions{})
	if tx.Error != nil {
		http.Error(w, fmt.Sprintf("There was an error initializing transaction. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	if err := apih.PostHelper(r, &o, tx); err != nil {
		http.Error(w, fmt.Sprintf("There was an error saving the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(o.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error generation JWT")

		http.Error(w, fmt.Sprintf("There was an error saving the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	message := newUserConvTree.RootNode.ResponseText
	qrs := chatbot.QuickReplyStringSlice(newUserConvTree.RootNode.QuickReplies)

	response := Message{
		Message:      message,
		QuickReplies: qrs,
		Token:        token,
	}

	if tx.Commit().Error != nil {
		http.Error(w, fmt.Sprintf("There was an error saving the %s. Please contact your administrator.\n", utils.GetType(o)), http.StatusInternalServerError)
		return
	}

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// PostResident creates a new resident
func PostResident(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

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

// PutResident updates an existing resident
func PutResident(w http.ResponseWriter, r *http.Request) {

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

	// Get the variables from the request
	vars := mux.Vars(r)
	IDStr := vars["resident"]
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

// GetResident gets a specific resident
func GetResident(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.Resident

	// Get the variables from the request
	vars := mux.Vars(r)
	IDStr := vars["resident"]
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

// GetResidents gets all residents from Dynamo
func GetResidents(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.Residents

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

// ResidentSetFirstName sets the first name for Resident
func ResidentSetFirstName(conn utils.Conn, r interface{}, firstName string) error {

	rsdnt, ok := r.(*boomyDB.Resident)
	if !ok {
		return fmt.Errorf("interface not resident")
	}

	// fmt.Println("ResidentSetFirstName rsdnt", rsdnt)

	rsdnt.FirstName = firstName

	// fmt.Println("ResidentSetFirstName after rsdnt", rsdnt)

	return rsdnt.Update(conn)
}

// ResidentSetLastName sets the last name for Resident
func ResidentSetLastName(conn utils.Conn, r interface{}, lastName string) error {
	rsdnt, ok := r.(*boomyDB.Resident)
	if !ok {
		return fmt.Errorf("interface not resident")
	}

	// fmt.Println("ResidentSetFirstName rsdnt", rsdnt)

	rsdnt.LastName = lastName

	// fmt.Println("ResidentSetFirstName after rsdnt", rsdnt)

	return rsdnt.Update(conn)
}

// ResidentSendValidationCode validates phone number sent,
// generates a random code, sends the code and sets the phone number for Resident
func ResidentSendValidationCode(conn utils.Conn, r interface{}, phoneNumber string) error {
	rsdnt, ok := r.(*boomyDB.Resident)
	if !ok {
		return fmt.Errorf("interface not resident")
	}

	// Validate phone number
	if err := validatePhoneNumber(phoneNumber); err != nil {
		return err
	}

	// Generate random code
	randomCode, err := utils.RandomSecret(5)
	if err != nil {
		return err
	}

	// Send validation code
	if err := sendValidationCode(phoneNumber, randomCode); err != nil {
		return err
	}

	// Update Resident in Dynamo
	rsdnt.ValidationCode = randomCode
	rsdnt.PhoneNumber = phoneNumber

	return rsdnt.Update(conn)
}

// ResidentCheckValidationCode verifies the code typed by the user is the same that was sent
func ResidentCheckValidationCode(conn utils.Conn, r interface{}, code string) error {
	rsdnt, ok := r.(*boomyDB.Resident)
	if !ok {
		return fmt.Errorf("interface not resident")
	}

	// Validate code
	if strings.ToUpper(rsdnt.ValidationCode) != strings.ToUpper(code) {
		return fmt.Errorf("code is not valid")
	}

	return nil
}

// ResidentRegenerateValidationCode generates a new validation code and sends it to the phone number
// saved for the Resident
func ResidentRegenerateValidationCode(conn utils.Conn, r interface{}, null string) error {
	rsdnt, ok := r.(*boomyDB.Resident)
	if !ok {
		return fmt.Errorf("interface not resident")
	}

	// Generate random code
	randomCode, err := utils.RandomSecret(5)
	if err != nil {
		return err
	}

	// Send validation code
	if err := sendValidationCode(rsdnt.PhoneNumber, randomCode); err != nil {
		return err
	}

	// Update Resident in Dynamo
	rsdnt.ValidationCode = randomCode

	return rsdnt.Update(conn)
}
