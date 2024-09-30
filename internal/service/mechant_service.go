package service

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com.ivanrafli14/ecommerce-golang/pkg/validation"
)

type IMerchantService interface {
	CreateMerchant(merchantReq *entity.MerchantRequest) error
	GetMerchantByAuthID(authID int) (*entity.Merchant, error)
	UpdateMerchant(merchantReq *entity.MerchantRequest) error
}

type MerchantService struct {
	repo repository.IMerchantRepository
}

func NewMerchantService(repo repository.IMerchantRepository) IMerchantService {
	return &MerchantService{
		repo: repo,
	}
}

func (s *MerchantService) CreateMerchant(merchantReq *entity.MerchantRequest) error {
	if merchantReq.Role != "merchant" {
		return response.ErrRoleInvalid
	}

	err := validation.ValidateMerchantReq(*merchantReq)
	if err != nil {
		return err
	}
	_, err = s.repo.GetMerchantByAuthId(merchantReq.AuthID)

	if err != response.ErrMerchantNotFound {
		return response.ErrorMerchantAlreadyExits
	}
	merchant := &entity.Merchant{
		Name:        merchantReq.Name,
		PhoneNumber: merchantReq.PhoneNumber,
		Address:     merchantReq.Address,
		City:        merchantReq.City,
		ImageUrl:    merchantReq.ImageUrl,
		AuthID:      merchantReq.AuthID,
	}
	err = s.repo.CreateMerchant(merchant)
	if err != nil {
		return err
	}
	return nil

}

func (s *MerchantService) GetMerchantByAuthID(authID int) (*entity.Merchant, error) {
	merchant, err := s.repo.GetMerchantByAuthId(authID)
	if err != nil {
		return nil, err
	}
	return merchant, nil
}

func (s *MerchantService) UpdateMerchant(merchantReq *entity.MerchantRequest) error {
	err := validation.ValidateMerchantReq(*merchantReq)
	if err != nil {
		return err
	}

	_, err = s.repo.GetMerchantByAuthId(merchantReq.AuthID)

	if err != nil {
		return err
	}
	merchant := &entity.Merchant{
		Name:        merchantReq.Name,
		PhoneNumber: merchantReq.PhoneNumber,
		Address:     merchantReq.Address,
		City:        merchantReq.City,
		ImageUrl:    merchantReq.ImageUrl,
		AuthID:      merchantReq.AuthID,
	}

	err = s.repo.UpdateMerchant(merchant)
	if err != nil {
		return err
	}
	return nil
}
