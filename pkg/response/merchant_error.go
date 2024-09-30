package response

import (
	"errors"
	"net/http"
)

var (
	ErrNameMerchantRequired          = errors.New("name is required")
	ErrAddressMerchantRequired       = errors.New("address is required")
	ErrPhoneNumberMerchantRequired   = errors.New("phone_number is required")
	ErrPhoneNumberMerchantInvalidLen = errors.New("phone_number length must be greater than equal 10")
	ErrImageUrlMerchantRequired      = errors.New("image_url is required")
	ErrCityMerchantRequired          = errors.New("city is required")
	ErrMerchantAlreadyExits          = errors.New("merchant already exits")
)

var (
	ErrorNameMerchantRequired     = NewError("bad request", "40001", http.StatusBadRequest)
	ErrorAddressMerchantRequired  = NewError("bad request", "40002", http.StatusBadRequest)
	ErrorPhoneMerchantRequired    = NewError("bad request", "40003", http.StatusBadRequest)
	ErrorPhoneMerchantLength      = NewError("bad request", "40004", http.StatusBadRequest)
	ErrorImageUrlMerchantRequired = NewError("bad request", "40005", http.StatusBadRequest)
	ErrorCityMerchantRequired     = NewError("bad request", "40006", http.StatusBadRequest)
	ErrorMerchantAlreadyExits     = NewError("conflict", "40901", http.StatusConflict)
)
