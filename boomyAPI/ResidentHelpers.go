package boomyAPI

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/utils"
)

// getResidentFromJWT will read from the token and return Resident passed in request
func getResidentIDFromJWT(r *http.Request) (string, error) {

	interfaceAPI, err := utils.GetCustomStruct(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error retrieving Residents struct")

		return "", err
	}

	ID, ok := interfaceAPI.(string)
	if !ok {
		logrus.Warn("Error converting custom_struct to string")

		return "", fmt.Errorf("error retrieving Residents struct")
	}

	return ID, nil

}

func validatePhoneNumber(phoneNumber string) error {
	return nil
}

func sendValidationCode(phoneNumber, code string) error {
	fmt.Println("CODE:", code)
	return nil
}
