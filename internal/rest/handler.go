package rest

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

// generic constants used across functions
const (
	InvalidRequestBody        = `Not able to parse request body.`
	UnableToDecodeRequestBody = `Cannot decode request-body data json.`
	ErrorInvalidRequestBody   = `Error while validating request body`
)

//Handler - This interface contains the methods for -
// a) generic handling of response (201 Created) b) Generic body parser
type Handler interface {
	ParseRequest(reqBody io.Reader, out interface{}, logger *logrus.Entry) error
	RespondCreated(w http.ResponseWriter, body interface{}, logger *logrus.Entry)
}

//DefaultHandler - default implementation of handler
type DefaultHandler struct{}

//ParseRequest - This method contains the generic handler for parsing the http request
func (defaultHandler *DefaultHandler) ParseRequest(reqBody io.Reader, out interface{}, logger *logrus.Entry) error {

	bodyBuffer, err := ioutil.ReadAll(reqBody)
	if err != nil {
		//logger.WithError(err).Error(InvalidRequestBody)
		return err
	}

	err = json.Unmarshal(bodyBuffer, &out)
	if err != nil {
		logger.WithField(`request-body`, string(bodyBuffer)).WithError(err).Error(UnableToDecodeRequestBody)
		return err
	}
	return nil
}

//RespondCreated - This method contains the generic handler for parsing 201 Created response to be set to client for rest api
func (defaultHandler *DefaultHandler) RespondCreated(w http.ResponseWriter, body interface{}, logger *logrus.Entry) {
	w.Header().Set("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusCreated)
		enc := json.NewEncoder(w)
		err := enc.Encode(body)
		if err != nil {
			logger.WithError(err).Error("Error writing response")
		}
	}
}
