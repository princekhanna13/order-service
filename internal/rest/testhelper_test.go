package rest

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"order-service/internal/model"
)

/**
Use these default mocks while testing.
This file contains an valid mock implementation and error mock implementation.
So we can test the valid case as well as invalid case.
*/

type mockErrorRestHandler struct{}

func (fetcher *mockErrorRestHandler) ParseRequest(_ io.Reader, _ interface{}, _ *logrus.Entry) error {
	return errors.New(InvalidRequestBody)
}

func (fetcher *mockErrorRestHandler) RespondCreated(w http.ResponseWriter, body interface{}, _ *logrus.Entry) {
	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	}
}

type mockRestHandler struct{}

func (fetcher *mockRestHandler) ParseRequest(reqBody io.Reader, out interface{}, _ *logrus.Entry) error {
	_ = json.NewDecoder(reqBody).Decode(&out)
	return nil
}

func (fetcher *mockRestHandler) RespondCreated(w http.ResponseWriter, body interface{}, _ *logrus.Entry) {
	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(body)
	}

}

type mockErrorValidator struct{}

func (mockErrorValidator *mockErrorValidator) ValidateStruct(object interface{}, logger *logrus.Entry) error {
	return errors.New(``)
}

type mockValidator struct{}

func (mockValidator *mockValidator) ValidateStruct(object interface{}, logger *logrus.Entry) error {
	return nil
}

type mockOrderService struct{}

func (os mockOrderService) CreateOrder(_ model.Order) (string, error) {
	return "ID", nil
}

type mockErrorOrderService struct{}

func (os mockErrorOrderService) CreateOrder(_ model.Order) (string, error) {
	return "", errors.New("Error Creating Order")
}
