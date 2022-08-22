package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"order-service/internal/service"
	"reflect"
	"strings"
	"testing"
)

func TestHandleOrderData(t *testing.T) {

	setupTests()

	invalidRequestBody := `
	{
    "bank_id":"test",
	"amount":"test"
}`
	validationViolationRequestBody := `
	{ "bank_id":"test"}`

	validRequestBody := `
	{
    "bank_id":"test",
    "amount":2000
}`

	type test struct {
		name             string
		input            string
		mockHandler      Handler
		mockValidator    Validator
		mockOrderService service.OrderService
		want             Response
	}

	tests := []test{
		{
			name:             `invalidRequestBody not saved`,
			input:            invalidRequestBody,
			mockHandler:      new(mockErrorRestHandler),
			mockValidator:    new(mockValidator),
			mockOrderService: new(mockOrderService),
			want:             Response{Code: 400, Message: `Not able to parse request body.`},
		},
		{
			name:             ErrorInvalidRequestBody,
			input:            validationViolationRequestBody,
			mockHandler:      new(mockRestHandler),
			mockValidator:    new(mockErrorValidator),
			mockOrderService: new(mockOrderService),
			want:             Response{Code: 400, Message: ""},
		},
		{
			name:             `Error Order Service`,
			input:            validRequestBody,
			mockHandler:      new(mockRestHandler),
			mockValidator:    new(mockValidator),
			mockOrderService: new(mockErrorOrderService),
			want:             Response{Code: 500, Message: `Error Creating Order`},
		},
		{
			name:             `validRequestBody saved successfully`,
			input:            validRequestBody,
			mockHandler:      new(mockRestHandler),
			mockValidator:    new(mockValidator),
			mockOrderService: new(mockOrderService),
			want:             Response{Code: 201, Message: `Order has been created successfully. OrderID : ID`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "http://dummy", strings.NewReader(tt.input))
			if err != nil {
				log.Println(err)
			}

			orderController := &OrderController{orderService: tt.mockOrderService}
			orderController.createOrderInternal(w, request, tt.mockHandler, tt.mockValidator)

			got := Response{}
			unmarshalErr := json.Unmarshal(w.Body.Bytes(), &got)

			if unmarshalErr != nil {
				t.Fatalf("Error unmarshalling Response")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleOrderData() = %v, want %v", got, tt.want)
			}

		})
	}
}
