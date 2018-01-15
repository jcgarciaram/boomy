package main

/*
import (
	// "time"
	"net/http"

	"bitbucket.org/wmsight/flourish-api/apiutils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func VerifyPermissions(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the variables from the request
		vars := mux.Vars(r)
		facilityId := vars["facility"]

		if facilityId != "" {
			// Verify that facility passed is an ObjectId
			if ok := bson.IsObjectIdHex(facilityId); !ok {
				http.Error(w, "Invalid facility", http.StatusBadRequest)
				return
			}

			// Verify if user has permission to perform whatever they are attempting
			if ok := apiutils.VerifyPerm(r, facilityId); !ok {
				http.Error(w, "User does not have permission to perform this action or there was an error processing the request", http.StatusUnauthorized)
				return
			}
		}

		inner.ServeHTTP(w, r)
	})
}
*/
