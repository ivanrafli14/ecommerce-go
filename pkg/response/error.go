package response

import (
	"errors"
	"net/http"
)

type Error struct {
	Message   string
	ErrorCode string
	HttpCode  int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:   msg,
		ErrorCode: code,
		HttpCode:  httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrGeneral    = errors.New("unknown error")
	ErrRepository = errors.New("error repository")
	ErrBadRequest = errors.New("bad request")
)

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorRepository = NewError("internal server error", "50001", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)
