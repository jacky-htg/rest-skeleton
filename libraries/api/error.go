package api

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Error struct {
	Err           error
	Status        string
	MessageStatus string
	HTTPStatus    int
}

func (err *Error) Error() string {
	return err.Err.Error()
}

func ErrNew(err error, status string, messageStatus string, httpStatus int) error {
	return &Error{err, status, messageStatus, httpStatus}
}

func ErrBadRequest(err error, message string) error {
	if len(message) <= 0 || message == "" {
		message = StatusMessageBadRequest
	}
	return &Error{err, StatusCodeBadRequest, message, http.StatusBadRequest}
}

// ErrNotFound wraps a provided error with an HTTP status code and custome status code for not found. This
// function should be used when handlers encounter expected errors.
func ErrNotFound(err error, message string) error {
	if len(message) <= 0 || message == "" {
		message = StatusMessageNotFound
	}
	return &Error{err, StatusCodeNotFound, message, http.StatusNotFound}
}

// ErrForbidden wraps a provided error with an HTTP status code and custome status code for forbidden. This
// function should be used when handlers encounter expected errors.
func ErrForbidden(err error, message string) error {
	if len(message) <= 0 || message == "" {
		message = StatusMessageForbidden
	}
	return &Error{err, StatusCodeForbidden, message, http.StatusForbidden}
}
