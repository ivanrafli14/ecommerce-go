package validation

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
)

func ValidateMerchantReq(merchantReq entity.MerchantRequest) error {
	if merchantReq.Name == "" {
		return response.ErrNameMerchantRequired
	}
	if merchantReq.Address == "" {
		return response.ErrAddressMerchantRequired
	}
	if merchantReq.PhoneNumber == "" {
		return response.ErrPhoneNumberMerchantRequired
	}
	if len(merchantReq.PhoneNumber) < 9 {
		return response.ErrPhoneNumberMerchantInvalidLen
	}
	if merchantReq.City == "" {
		return response.ErrCityMerchantRequired
	}
	if merchantReq.ImageUrl == "" {
		return response.ErrImageUrlMerchantRequired
	}
	return nil
}
