package errors

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

const (
	// NotFoundErrCode is the Error code for car not found case
	NotFoundErrCode = "0001"
	// NotFoundErrMessage is the Error message for car not found case
	NotFoundErrMessage = "Car by id : %s not found"
	// MissingMakeFieldForCreateCode is the Error code for make field missing case in payload
	MissingMakeFieldForCreateCode = "0002"
	// MissingMakeFieldForCreateMsg is the Error message for make field missing case in payload
	MissingMakeFieldForCreateMsg = "Missing required Make field in request payload"
	// MissingModelFieldForCreateCode is the Error code for model field missing case in payload
	MissingModelFieldForCreateCode = "0003"
	// MissingModelFieldForCreateMsg is the Error message for make field missing case in payload
	MissingModelFieldForCreateMsg = "Missing required Model field in request payload"
	// MissingYearFieldForCreateCode is the Error code for year field missing case in payload
	MissingYearFieldForCreateCode = "0004"
	// MissingYearFieldForCreateMsg is the Error message for make field missing case in payload
	MissingYearFieldForCreateMsg = "Missing required Year field in request payload"
)

// CarError strct type for any car mgmt errors
type CarError struct {
	Code    string
	Message string
}

func (e CarError) Error() string {
	return fmt.Sprintf("%s. Code : %s", e.Message, e.Code)
}

// CarNotFoundErr is the helper error function for car not found case
func CarNotFoundErr(id string) CarError {
	return CarError{
		Message: fmt.Sprintf(NotFoundErrMessage, id),
		Code:    NotFoundErrCode,
	}
}

// New is the helper error function for car not found case
func New(code string, message string) CarError {
	return CarError{
		Message: message,
		Code:    code,
	}
}

// HTTPErrInvalidRequest helper to send error response for invalid requests
func HTTPErrInvalidRequest(err error) render.Renderer {
	// If generic error send without App Code since we might not be aware of it. Else add AppCode
	errRes := &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     err.Error(),
	}
	if carErr, ok := err.(CarError); ok {
		errRes.AppCode = carErr.Code
		errRes.StatusText = carErr.Message
	}
	return errRes
}

// HTTPErrInternalServer helper to send error response for internal server error
func HTTPErrInternalServer(err error) render.Renderer {
	// If generic error send without App Code since we might not be aware of it. Else add AppCode
	errRes := &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     err.Error(),
	}
	if carErr, ok := err.(CarError); ok {
		errRes.AppCode = carErr.Code
		errRes.StatusText = carErr.Message
	}
	return errRes
}

// HTTPErrNotFound helper to send error response for not found
func HTTPErrNotFound(err error) render.Renderer {
	// If generic error send without App Code since we might not be aware of it. Else add AppCode
	errRes := &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     err.Error(),
	}
	if carErr, ok := err.(CarError); ok {
		errRes.AppCode = carErr.Code
		errRes.StatusText = carErr.Message
	}
	return errRes
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error  `json:"-"`           // low-level runtime error
	HTTPStatusCode int    `json:"status_code"` // http response status code
	StatusText     string `json:"error_message"`
	AppCode        string `json:"error_code"` // user-level status message
}

// Render method for error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
