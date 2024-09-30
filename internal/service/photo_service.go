package service

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/cloudinary"
)

type IPhotoService interface {
	UploadPhoto(photoReq entity.PhotoRequest) (string, error)
}

type PhotoService struct {
	cloudinary cloudinary.ICloudinary
}

func NewPhotoService(cloudinary cloudinary.ICloudinary) IPhotoService {
	return &PhotoService{cloudinary: cloudinary}
}

func (s *PhotoService) UploadPhoto(photoReq entity.PhotoRequest) (string, error) {
	url, err := s.cloudinary.UploadPhoto(photoReq)
	if err != nil {
		return "", err
	}
	return url, nil
}
