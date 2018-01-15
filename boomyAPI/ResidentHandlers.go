package boomyAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jcgarciaram/boomy/chatbot"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	apih "github.com/jcgarciaram/boomy/boomyAPI/APIHelpers"
	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/utils"
)

// ResidentBeginConversation creates a new Resident and begins a new conversation with them
func ResidentBeginConversation(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	o.Conversation = chatbot.Conversation{
		ConversationTreeID: newUserConvTree.GetID(),
	}

	o.FirstName = "Juan"

	// fmt.Println(o.Conversation)

	if err := apih.PostHelper(r, &o); err != nil {
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

	// Encode array returned into JSON and return
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// PostResident creates a new Resident
func PostResident(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

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

// PutResident updates an existing Resident
func PutResident(w http.ResponseWriter, r *http.Request) {

	// Get the variables from the request
	vars := mux.Vars(r)
	ID := vars["resident"]

	// Struct to unmarshal body of request into
	var o boomyDB.Resident

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

// GetResident gets a specific Resident
func GetResident(w http.ResponseWriter, r *http.Request) {

	// Get the variables from the request
	vars := mux.Vars(r)
	ID := vars["resident"]

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.Resident

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

// GetResidents gets all Residents from Dynamo
func GetResidents(w http.ResponseWriter, r *http.Request) {

	// Struct to unmarshal result from Dynamo into
	var o boomyDB.Residents

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

// ResidentSetFirstName sets the first name for Resident
func ResidentSetFirstName(r interface{}, firstName string) error {

	rsdnt := r.(*boomyDB.Resident)

	// fmt.Println("ResidentSetFirstName rsdnt", rsdnt)

	rsdnt.FirstName = firstName

	// fmt.Println("ResidentSetFirstName after rsdnt", rsdnt)

	return rsdnt.Save()
}

// ResidentSetLastName sets the last name for Resident
func ResidentSetLastName(r interface{}, lastName string) error {
	rsdnt := r.(*boomyDB.Resident)

	// fmt.Println("ResidentSetFirstName rsdnt", rsdnt)

	rsdnt.LastName = lastName

	// fmt.Println("ResidentSetFirstName after rsdnt", rsdnt)

	return rsdnt.Save()
}

// ResidentSendValidationCode validates phone number sent,
// generates a random code, sends the code and sets the phone number for Resident
func ResidentSendValidationCode(r interface{}, phoneNumber string) error {
	rsdnt := r.(*boomyDB.Resident)

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

	return rsdnt.Save()
}

// ResidentCheckValidationCode verifies the code typed by the user is the same that was sent
func ResidentCheckValidationCode(r interface{}, code string) error {
	rsdnt := r.(*boomyDB.Resident)

	// Validate code
	if strings.ToUpper(rsdnt.ValidationCode) != strings.ToUpper(code) {
		return fmt.Errorf("code is not valid")
	}

	return nil
}

// ResidentRegenerateValidationCode generates a new validation code and sends it to the phone number
// saved for the Resident
func ResidentRegenerateValidationCode(r interface{}, null string) error {
	rsdnt := r.(*boomyDB.Resident)

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

	return rsdnt.Save()
}
