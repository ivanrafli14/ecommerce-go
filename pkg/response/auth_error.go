package response

import (
	"errors"
	"net/http"
)

var (
	ErrEmailNotFound         = errors.New("email not found")
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password length must be greater than equal 6")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")
	ErrJWTEmpty              = errors.New("jwt is empty,please provide jwt token")
	ErrJWTExpired            = errors.New("jwt has expired. Please login to generate new one")
	ErrJWTInvalid            = errors.New("jwt is invalid. Please provide valid token")
)

var (
	ErrorEmailNotFound         = NewError("not found", "40401", http.StatusNotFound)
	ErrorEmailRequired         = NewError("bad request", "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError("bad request", "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError("bad request", "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError("bad request", "40004", http.StatusBadRequest)
	ErrorEmailAlreadyUsed      = NewError("duplicate entry", "40901", http.StatusConflict)

	ErrorJWTEmpty         = NewError("unauthorized", "40101", http.StatusUnauthorized)
	ErrorJWTExpired       = NewError("expired jwt", "40102", http.StatusUnauthorized)
	ErrorJWTInvalid       = NewError("invalid jwt", "40103", http.StatusUnauthorized)
	ErrorPasswordNotMatch = NewError("unauthorized", "40104", http.StatusUnauthorized)
)
