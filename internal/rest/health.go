package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//CheckResponse - Struct type representing the response from health check
type CheckResponse struct {
	Status string `json:"status"`
}

//HealthHandler  for health check
func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response CheckResponse
		status := http.StatusOK
		response.Status = "UP"
		data, _ := json.Marshal(response)
		writeJSONResponse(w, status, data)
	}
}

func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Something wrong happened while writing the JSON response"))
	}
}
