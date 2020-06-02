package boomyAPI

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/boomyAPI/APIHelpers"
	"github.com/jcgarciaram/boomy/boomyDB"
	"github.com/jcgarciaram/boomy/chatbot"
)

// ResidentPostMessage receives a message from the caller and gets response from chatbot
func ResidentPostMessage(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	residentID, err := getResidentIDFromJWT(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting residentID from JWT")

		http.Error(w, "Error getting residentID from JWT", http.StatusInternalServerError)
		return
	}

	tx := db.BeginTx(context.Background(), &sql.TxOptions{})
	if tx.Error != nil {
		http.Error(w, fmt.Sprintf("There was an error initializing transaction. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	var rsdnt boomyDB.Resident
	if err := rsdnt.First(tx, residentID); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting resident from database")

		http.Error(w, "Error getting resident from database", http.StatusInternalServerError)
		return
	}

	// String to unmarshal message
	var m Message

	if err := APIHelpers.UnmarshalBody(r, &m); err != nil {
		http.Error(w, "Error reading message", http.StatusInternalServerError)
		return
	}

	message, qrs, err := chatbot.GetResponse(&rsdnt, &rsdnt.Conversation, m.Message)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting resident from database")

		http.Error(w, "Error getting resident from database", http.StatusInternalServerError)
		return
	}

	// Save resident to save the new state
	if err := rsdnt.Update(tx); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error saving resident")

		http.Error(w, "Error saving resident", http.StatusInternalServerError)
		return
	}

	if tx.Commit().Error != nil {
		http.Error(w, fmt.Sprintf("There was an error. Please contact your administrator.\n"), http.StatusInternalServerError)
		return
	}

	response := Message{
		Message:      message,
		QuickReplies: qrs,
	}

	// Encode message into JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}

// ResidentGetConversation returns the current state of the conversation
func ResidentGetConversation(w http.ResponseWriter, r *http.Request) {
	// Set content type returned to JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	residentID, err := getResidentIDFromJWT(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting residentID from JWT")

		http.Error(w, "Error getting residentID from JWT", http.StatusInternalServerError)
		return
	}

	var rsdnt boomyDB.Resident
	if err := rsdnt.First(db, residentID); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting resident from database")

		http.Error(w, "Error getting resident from database", http.StatusInternalServerError)
		return
	}

	message, qrs, err := chatbot.GetCurrentResponse(&rsdnt, &rsdnt.Conversation)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error getting resident from database")

		http.Error(w, "Error getting resident from database", http.StatusInternalServerError)
		return
	}

	response := Message{
		Message:      message,
		QuickReplies: qrs,
	}

	// Encode message into JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error encoding retStruct")
	}
}
