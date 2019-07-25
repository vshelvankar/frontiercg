package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrInvalidRequest helper to send error response for invalid requests
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     err.Error(),
	}
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error  `json:"-"`             // low-level runtime error
	HTTPStatusCode int    `json:"status_code"`   // http response status code
	StatusText     string `json:"error_message"` // user-level status message
}

// Render method for error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
