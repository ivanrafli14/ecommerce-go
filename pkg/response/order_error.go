package response

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidQuantity        = errors.New("invalid quantity")
	ErrMaxReachQuantity       = errors.New("max reach quantity")
	ErrProductOrderNotFound   = errors.New("product not found")
	ErrProductIDOrderRequired = errors.New("product_id is required")
	ErrQuantityOrderRequired  = errors.New("quantity is required")
)

var (
	ErrorInvalidQuantity        = NewError("bad request", "40001", http.StatusBadRequest)
	ErrorMaxReachQuantity       = NewError("bad request", "40002", http.StatusBadRequest)
	ErrorProductOrderNotFound   = NewError("bad request", "40003", http.StatusBadRequest)
	ErrorProductIDOrderRequired = NewError("bad request", "40004", http.StatusBadRequest)
	ErrorQuantityOrderRequired  = NewError("bad request", "40005", http.StatusBadRequest)
)
