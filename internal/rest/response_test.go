package rest

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-chi/render"
)

func TestErrInternalServer(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want render.Renderer
	}{
		{
			name: "expect ErrResponse with nil details",
			args: args{
				err: errors.New(`internal server error`),
			},
			want: &ErrResponse{
				Err:            errors.New(`internal server error`),
				HTTPStatusCode: 500,
				StatusText:     "Internal server error.",
				ErrorText:      "internal server error",
			},
		}, {
			name: "expect ErrResponse with one error details",
			args: args{
				err: errors.New(`internal server error`),
			},
			want: &ErrResponse{
				Err:            errors.New(`internal server error`),
				HTTPStatusCode: 500,
				StatusText:     "Internal server error.",
				ErrorText:      "internal server error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrInternalServer(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrInternalServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrInvalidRequest(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want render.Renderer
	}{
		{
			name: "expect ErrResponse with nil details",
			args: args{
				err: errors.New(`invalid request`),
			},
			want: &ErrResponse{
				Err:            errors.New(`invalid request`),
				HTTPStatusCode: 400,
				StatusText:     "Invalid request.",
				ErrorText:      "invalid request",
			},
		}, {
			name: "expect ErrResponse with one error details",
			args: args{
				err: errors.New(`invalid request`),
			},
			want: &ErrResponse{
				Err:            errors.New(`invalid request`),
				HTTPStatusCode: 400,
				StatusText:     "Invalid request.",
				ErrorText:      "invalid request",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrInvalidRequest(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrInvalidRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
