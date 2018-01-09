package APIHelpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
)

func PostHelper(r *http.Request, o TypeHelper) error {
	if err := UnmarshalBody(r, &o); err != nil {
		return err
	}

	// Validate object
	if err := o.Validate(); err != nil {
		return err
	}

	// Save object
	return o.Save()
}

func PutHelper(r *http.Request, o TypeHelper, ID string) error {

	if err := o.Get(ID); err != nil {
		return err
	}

	if err := UnmarshalBody(r, &o); err != nil {
		return err
	}

	// Validate object
	if err := o.Validate(); err != nil {
		return err
	}

	// Save object
	return o.Save()
}

func GetHelper(r *http.Request, o TypeHelper, ID string) error {
	return o.Get(ID)
}

func GetAllHelper(r *http.Request, o SliceTypeHelper) error {
	return o.GetAll()
}

func UnmarshalBody(r *http.Request, o interface{}) error {
	// Read body which contains the fields
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error reading create body")

		return err
	}

	// Close body
	if err := r.Body.Close(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error closing request body")

		return err
	}

	// Unmarshal body into object pointer passed in request
	if err := json.Unmarshal(body, o); err != nil {

		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Error marshaling JSON to packageConfig struct")

		return err
	}

	return nil
}
