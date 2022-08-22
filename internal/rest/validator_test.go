package rest

import (
	"context"
	"order-service/internal/middleware/httplogger"
	"testing"
)

//Test file for validator

func TestDefaultValidator_validateStruct(t *testing.T) {

	con := context.Background()
	logger := httplogger.GetLogger(con)

	tests := []struct {
		name    string
		input   interface{}
		wantErr bool
	}{
		{
			name: `valid struct with json tag`,
			input: struct {
				Name string `validate:"required" json:"name.firstname"`
			}{
				Name: "aaa",
			},
			wantErr: false,
		},
		{
			name: `valid struct with json tag throws error`,
			input: struct {
				Name string `validate:"required" json:"name.firstname"`
			}{
				Name: ``,
			},
			wantErr: true,
		}, {
			name: `valid struct without json tag`,
			input: struct {
				Name string `validate:"required"`
			}{
				Name: `aaa`,
			},
			wantErr: false,
		}, {
			name: `valid struct with json:"-"`,
			input: struct {
				Name string `json:"-"`
			}{
				Name: "X",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultValidator := &DefaultValidator{}
			if err := defaultValidator.ValidateStruct(tt.input, logger); (err != nil) != tt.wantErr {
				t.Errorf("DefaultValidator.ValidateStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
