package rest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"order-service/internal/middleware/httplogger"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDefaultHandler_parseRequest(t *testing.T) {

	con := context.Background()
	logger := httplogger.GetLogger(con)

	type args struct {
		reqBody io.Reader
		out     interface{}
		logger  *logrus.Entry
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: `invalid json parse failure`,
			args: args{
				reqBody: strings.NewReader(`abracadabra`),
				out:     Response{},
				logger:  logger,
			},
			wantErr: true,
		},
		{
			name: `valid json parse success`,
			args: args{
				reqBody: strings.NewReader(`{}`),
				out:     Response{},
				logger:  logger,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultHandler := &DefaultHandler{}
			if err := defaultHandler.ParseRequest(tt.args.reqBody, tt.args.out, tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("DefaultHandler.ParseRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefaultHandler_respondCreated(t *testing.T) {

	con := context.Background()
	logger := httplogger.GetLogger(con)

	w := httptest.NewRecorder()

	type args struct {
		w      http.ResponseWriter
		body   interface{}
		logger *logrus.Entry
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: `response with a body is 201 Created`,
			args: args{
				w:      w,
				body:   Response{Code: http.StatusCreated, Message: "Order Created"},
				logger: logger,
			},
		}, {
			name: `response without a body is 204 NO CONTENT`,
			args: args{
				w:      w,
				body:   nil,
				logger: logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultHandler := &DefaultHandler{}
			defaultHandler.RespondCreated(tt.args.w, tt.args.body, tt.args.logger)
		})
	}
}
