package response

import (
	"errors"
	"net/http"
)

var (
	ErrNameProductRequired        = errors.New("name is required")
	ErrDescriptionProductRequired = errors.New("description is required")
	ErrPriceProductRequired       = errors.New("price is required")
	ErrPriceProductInvalid        = errors.New("price is invalid")
	ErrStockProductRequired       = errors.New("stock is required")
	ErrStockProductInvalid        = errors.New("stock is invalid")
	ErrCategoryIDProductRequired  = errors.New("category_id is required")
	ErrImageUrlProductRequired    = errors.New("image_url is required")
	ErrCategoryIDProductNotFound  = errors.New("category_id is not found")
	ErrProductNotFound            = errors.New("product is not found in this resources")
)

var (
	ErrorPriceProductRequired       = NewError("bad request", "40001", http.StatusBadRequest)
	ErrorPriceProductInvalid        = NewError("bad request", "40002", http.StatusBadRequest)
	ErrorStockProductRequired       = NewError("bad request", "40003", http.StatusBadRequest)
	ErrorStockProductInvalid        = NewError("bad request", "40004", http.StatusBadRequest)
	ErrorNameProductRequired        = NewError("bad request", "40005", http.StatusBadRequest)
	ErrorDescriptionProductRequired = NewError("bad request", "40006", http.StatusBadRequest)
	ErrorImageUrlProductRequired    = NewError("bad request", "40007", http.StatusBadRequest)
	ErrorCategoryIDProductRequired  = NewError("bad request", "40008", http.StatusBadRequest)
	ErrorCategoryIDProductNotFound  = NewError("bad request", "40009", http.StatusBadRequest)
	ErrorProductNotFound            = NewError("not found", "40401", http.StatusNotFound)
)
