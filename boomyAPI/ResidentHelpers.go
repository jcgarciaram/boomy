package boomyAPI

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/jcgarciaram/boomy/utils"
)

// getResidentFromJWT will read from the token and return Resident passed in request
func getResidentIDFromJWT(r *http.Request) (uint, error) {

	interfaceAPI, err := utils.GetCustomStruct(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error retrieving Residents struct")

		return 0, err
	}

	fmt.Printf("\n\ninterfaceAPI %v %t\n\n", interfaceAPI, interfaceAPI)

	ID, ok := interfaceAPI.(float64)
	if !ok {
		logrus.Warn("Error converting custom_struct to int")

		return 0, fmt.Errorf("error retrieving Residents struct")
	}

	return uint(ID), nil

}

func validatePhoneNumber(phoneNumber string) error {
	return nil
}

func sendValidationCode(phoneNumber, code string) error {
	fmt.Println("CODE:", code)
	return nil
}
