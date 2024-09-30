package validation

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
)

func ValidateProductReq(productReq entity.ProductRequest) error {
	if productReq.Name == "" {
		return response.ErrNameProductRequired
	}
	if productReq.Description == "" {
		return response.ErrDescriptionProductRequired
	}
	if productReq.Price == 0 {
		return response.ErrPriceProductRequired
	}
	if productReq.Stock == 0 {
		return response.ErrStockProductRequired
	}
	if productReq.CategoryID == 0 {
		return response.ErrCategoryIDProductRequired
	}
	if productReq.ImageUrl == "" {
		return response.ErrImageUrlProductRequired
	}
	return nil
}
