package validation

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"time"
)

func ValidateUserReq(userReq entity.UserRequest) error {
	if userReq.Name == "" {
		return response.ErrNameRequired
	}
	if err := validateDateOfBirth(userReq.DateOfBirth); err != nil {
		return err
	}

	if err := validatePhoneNumber(userReq.PhoneNumber); err != nil {
		return err
	}

	if err := ValidateGender(userReq.Gender); err != nil {
		return err
	}
	if userReq.Address == "" {
		return response.ErrAddressRequired
	}
	if userReq.Address == "" {
		return response.ErrAddressRequired
	}
	return nil

}

func validateDateOfBirth(dateOfBirth string) error {
	if dateOfBirth == "" {
		return response.ErrDateOfBirthRequired
	}
	_, err := time.Parse("2006-01-02", dateOfBirth)

	if err != nil {
		return response.ErrDateOfBirthInvalid
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	if phoneNumber == "" {
		return response.ErrPhoneNumberRequired
	}
	if len(phoneNumber) < 9 {
		return response.ErrPhoneNumberInvalidLength
	}
	return nil
}

func ValidateGender(gender string) error {
	if gender == "" {
		return response.ErrGenderRequired
	}

	if !(gender == "male" || gender == "female") {
		return response.ErrGenderInvalid
	}
	return nil
}
