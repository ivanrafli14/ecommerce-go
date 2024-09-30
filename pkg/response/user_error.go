package response

import (
	"errors"
	"net/http"
)

var (
	ErrNameRequired             = errors.New("name is required")
	ErrDateOfBirthRequired      = errors.New("date of birth is required")
	ErrDateOfBirthInvalid       = errors.New("date of birth is invalid")
	ErrPhoneNumberRequired      = errors.New("phone number is required")
	ErrPhoneNumberInvalidLength = errors.New("phone_number length must be greater than equal 10")
	ErrAddressRequired          = errors.New("address is required")
	ErrImageUrlRequired         = errors.New("image url is required")
	ErrGenderRequired           = errors.New("gender is required")
	ErrGenderInvalid            = errors.New("gender is invalid")
	ErrRoleInvalid              = errors.New("invalid role")
	ErrUserNotFound             = errors.New("user not found")
	ErrMerchantNotFound         = errors.New("merchant not found")
	ErrUserAlreadyExists        = errors.New("user already exists")
)

var (
	ErrorGenderRequired           = NewError("bad request", "40001", http.StatusBadRequest)
	ErrorGenderInvalid            = NewError("bad request", "40002", http.StatusBadRequest)
	ErrorPhoneNumberRequired      = NewError("bad request", "40003", http.StatusBadRequest)
	ErrorPhoneNumberInvalidLength = NewError("bad request", "40004", http.StatusBadRequest)
	ErrorNameRequired             = NewError("bad request", "40005", http.StatusBadRequest)
	ErrorAddressRequired          = NewError("bad request", "40006", http.StatusBadRequest)
	ErrorDateOfBirthRequired      = NewError("bad request", "40007", http.StatusBadRequest)
	ErrorDateOfBirthInvalid       = NewError("bad request", "40008", http.StatusBadRequest)
	ErrorImageUrlRequired         = NewError("bad request", "40009", http.StatusBadRequest)
	ErrorRoleInvalid              = NewError("unauthorized", "40102", http.StatusUnauthorized)
	ErrorUserNotFound             = NewError("not found", "40401", http.StatusNotFound)
	ErrorMerchantNotFound         = NewError("not found", "40402", http.StatusNotFound)
	ErrorUserAlreadyExists        = NewError("conflict", "40901", http.StatusConflict)
)
