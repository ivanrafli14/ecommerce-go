package service

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
)

type ICategoryService interface {
	GetAllCategories() ([]entity.Category, error)
}

type CategoryService struct {
	repo repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) ICategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (c *CategoryService) GetAllCategories() ([]entity.Category, error) {
	category, err := c.repo.GetAllCategories()
	if err != nil {
		return []entity.Category{}, err
	}
	return category, nil
}
