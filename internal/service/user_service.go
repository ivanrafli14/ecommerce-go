package service

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com.ivanrafli14/ecommerce-golang/pkg/validation"
	"log"
	"time"
)

type IUserService interface {
	CreateUser(userReq entity.UserRequest) error
	UpdateUser(userReq entity.UserRequest) error
	GetUser(AuthID int) (*entity.UserResponse, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) UpdateUser(userReq entity.UserRequest) error {
	if err := validation.ValidateUserReq(userReq); err != nil {
		return err
	}

	parsedDate, err := time.Parse("2006-01-02", userReq.DateOfBirth)
	if err != nil {
		return err
	}

	_, err = s.repo.FindByAuthID(userReq.AuthID)
	if err != nil {
		return err
	}

	user := &entity.User{
		Name:        userReq.Name,
		DateOfBirth: parsedDate,
		PhoneNumber: userReq.PhoneNumber,
		Gender:      userReq.Gender,
		Address:     userReq.Address,
		ImageUrl:    userReq.ImageUrl,
		AuthID:      userReq.AuthID,
	}

	err = s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) CreateUser(userReq entity.UserRequest) error {
	log.Println(userReq.Role)
	if userReq.Role != "user" {
		return response.ErrRoleInvalid
	}

	if err := validation.ValidateUserReq(userReq); err != nil {
		return err
	}
	parsedDate, err := time.Parse("2006-01-02", userReq.DateOfBirth)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = s.repo.FindByAuthID(userReq.AuthID)

	if err != response.ErrUserNotFound {
		return response.ErrUserAlreadyExists
	}

	user := &entity.User{
		Name:        userReq.Name,
		DateOfBirth: parsedDate,
		PhoneNumber: userReq.PhoneNumber,
		Gender:      userReq.Gender,
		Address:     userReq.Address,
		ImageUrl:    userReq.ImageUrl,
		AuthID:      userReq.AuthID,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(AuthID int) (*entity.UserResponse, error) {
	user, err := s.repo.FindByAuthID(AuthID)
	if err != nil {
		return nil, err
	}
	dateParse := user.DateOfBirth.Format("2006-01-02")
	userResponse := &entity.UserResponse{
		Name:        user.Name,
		DateOfBirth: dateParse,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Address:     user.Address,
		ImageUrl:    user.ImageUrl,
	}

	return userResponse, nil
}
