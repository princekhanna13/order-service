package rest

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse - renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"`    // low-level runtime error
	HTTPStatusCode int   `json:"code"` // http response status code

	StatusText string `json:"-"`                 // user-level status message
	ErrorText  string `json:"message,omitempty"` // application-level error message, for debugging
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Render - func to render the error status text from code
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest - func to send error invalid status
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrInternalServer - func to send internal server error
func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error.",
		ErrorText:      err.Error(),
	}
}
