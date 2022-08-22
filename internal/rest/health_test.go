package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	mocket "github.com/selvatico/go-mocket"
)

func setupTests() {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
}

func TestHealthHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	setupTests()
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := HealthHandler()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"status":"UP"}`
	data, _ := json.Marshal(CheckResponse{Status: "UP"})
	if rr.Body.String() != expected || rr.Header().Get("Content-Type") != "application/json" || rr.Header().Get("Content-Length") != strconv.Itoa(len(data)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
