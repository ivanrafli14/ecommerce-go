package service

import (
	"errors"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/bcrypt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/jwt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/redis"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com.ivanrafli14/ecommerce-golang/pkg/validation"
	"log"
	"time"
)

type IAuthService interface {
	Login(authReq entity.AuthRequest) (entity.AuthLoginResponse, error)
	Register(authReq entity.AuthRequest) error
	UpdateRole(authID int) error
}

type AuthService struct {
	ar     repository.IAuthRepository
	bcrypt bcrypt.Interface
	jwt    jwt.Interface
	redis  redis.Interface
}

func NewAuthService(repo repository.IAuthRepository, bcrypt bcrypt.Interface, jwt jwt.Interface, redis redis.Interface) IAuthService {
	return &AuthService{
		ar:     repo,
		bcrypt: bcrypt,
		jwt:    jwt,
		redis:  redis,
	}
}

func (s *AuthService) Register(authReq entity.AuthRequest) error {
	err := validation.ValidateAuthReq(authReq)
	if err != nil {
		return err
	}

	auth, err := s.ar.GetAuthByEmail(authReq.Email)
	if err != nil {
		if err != response.ErrEmailNotFound {

			return err
		}
	}

	if auth != nil {
		return response.ErrEmailAlreadyUsed
	}

	hashedPassword, err := s.bcrypt.GenerateFromPassword(authReq.Password)

	if err != nil {
		return err
	}

	authModel := &entity.Auth{
		Email:    authReq.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	err = s.ar.CreateAuth(authModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(authReq entity.AuthRequest) (entity.AuthLoginResponse, error) {
	err := validation.ValidateAuthReq(authReq)
	var authResp entity.AuthLoginResponse

	if err != nil {
		return authResp, err
	}

	authRegistered, err := s.ar.GetAuthByEmail(authReq.Email)
	if err != nil {
		return authResp, err
	}

	if err := s.bcrypt.CompareHashAndPassword(authRegistered.Password, authReq.Password); err != nil {
		return authResp, response.ErrPasswordNotMatch
	}

	token, err := s.jwt.CreateJWTToken(authRegistered.ID, authRegistered.Role)
	if err != nil {
		return authResp, err
	}

	if err := s.redis.SetData(token, true, 1*time.Hour); err != nil {
		log.Println(err)
		return authResp, err
	}

	authResp = entity.AuthLoginResponse{
		Token: token,
		Role:  authRegistered.Role,
	}
	return authResp, nil

}

func (s *AuthService) UpdateRole(authID int) error {
	auth, err := s.ar.GetAuthByID(authID)
	if err != nil {
		return err
	}
	if auth.Role == "merchant" {
		return errors.New("user already as a merchant")
	}

	err = s.ar.UpdateRoleAuth(auth.ID)
	if err != nil {
		return err
	}
	return nil
}
