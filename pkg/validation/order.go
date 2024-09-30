package validation

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
)

func ValidateCheckoutReq(checkoutReq entity.CheckoutRequest) error {
	if checkoutReq.ProductID == 0 {
		return response.ErrProductIDOrderRequired
	}
	if checkoutReq.Quantity == 0 {
		return response.ErrQuantityOrderRequired
	}
	return nil
}
