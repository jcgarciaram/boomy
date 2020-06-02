package APIHelpers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jcgarciaram/boomy/utils"

	"github.com/Sirupsen/logrus"
)

func PostHelper(r *http.Request, o TypeHelper, conn utils.Conn) error {
	if err := UnmarshalBody(r, &o); err != nil {
		return err
	}

	// Validate object
	if err := o.Validate(); err != nil {
		return err
	}

	// Save object
	return o.Create(conn)
}

func PutHelper(r *http.Request, o TypeHelper, ID uint, conn utils.Conn) error {

	if err := o.First(conn, ID); err != nil {
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
	return o.Update(conn)
}

func GetHelper(r *http.Request, o TypeHelper, ID uint, conn utils.Conn) error {
	return o.First(conn, ID)
}

func GetAllHelper(r *http.Request, o SliceTypeHelper, conn utils.Conn) error {
	return o.Find(conn)
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
		}).Warn(fmt.Sprintf("Error marshaling JSON to %s struct", utils.GetType(o)))

		return err
	}

	return nil
}
