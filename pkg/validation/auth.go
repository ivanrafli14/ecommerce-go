package validation

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"strings"
)

func ValidateAuthReq(authReq entity.AuthRequest) error {
	if err := ValidateEmail(authReq.Email); err != nil {
		return err
	}
	if err := ValidatePassword(authReq.Password); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return response.ErrPasswordRequired
	}
	if len(password) < 6 {
		return response.ErrPasswordInvalidLength
	}
	return nil
}
