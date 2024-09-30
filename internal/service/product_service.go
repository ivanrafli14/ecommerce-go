package service

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/meilisearch"
	"github.com.ivanrafli14/ecommerce-golang/pkg/validation"
	"github.com/google/uuid"
	"log"
	"time"
)

type IProductService interface {
	CreateProduct(productReq entity.ProductRequest) error
	GetProductByProductID(productID int, authID int) (*entity.ProductDetailResponse, error)
	GetProductBySKU(SKU string) (entity.DetailProductSKUResponse, error)
	UpdateProduct(productReq entity.ProductRequest, productID int) error
	ListProducts(query string, limit, page, authID int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error)
	ListProductsAll(query string, limit int, page int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error)
}

type ProductService struct {
	productRepo  repository.IProductRepository
	merchantRepo repository.IMerchantRepository
	categoryRepo repository.ICategoryRepository
	meilisearch  meilisearch.SearchEngine
}

func NewProductService(pr repository.IProductRepository, mr repository.IMerchantRepository, cr repository.ICategoryRepository, meilisearch meilisearch.SearchEngine) IProductService {
	return &ProductService{
		productRepo:  pr,
		merchantRepo: mr,
		categoryRepo: cr,
		meilisearch:  meilisearch,
	}
}

func (s *ProductService) CreateProduct(productReq entity.ProductRequest) error {
	if err := validation.ValidateProductReq(productReq); err != nil {
		return err
	}

	categoryName, err := s.categoryRepo.GetCategoryNameByID(productReq.CategoryID)
	if err != nil {
		log.Println(err)
		return err
	}

	merchant, err := s.merchantRepo.GetMerchantByAuthId(productReq.AuthID)
	if err != nil {
		return err
	}

	product := entity.Product{
		SKU:         uuid.New().String(),
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		Stock:       productReq.Stock,
		CategoryID:  productReq.CategoryID,
		ImageUrl:    productReq.ImageUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		MerchantID:  merchant.ID,
	}

	if err := s.productRepo.CreateProduct(&product); err != nil {

		return err
	}

	err = s.meilisearch.StoreData(product, categoryName)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProductByProductID(productID int, authID int) (*entity.ProductDetailResponse, error) {
	merchant, err := s.merchantRepo.GetMerchantByAuthId(authID)
	if err != nil {
		return nil, err
	}

	product, err := s.productRepo.GetProductByProductID(productID, merchant.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProductBySKU(sku string) (entity.DetailProductSKUResponse, error) {
	productRes, err := s.productRepo.GetProductBySKU(sku)
	if err != nil {
		return entity.DetailProductSKUResponse{}, err
	}
	return productRes, nil
}

func (s *ProductService) UpdateProduct(productReq entity.ProductRequest, productID int) error {
	if err := validation.ValidateProductReq(productReq); err != nil {
		return err
	}
	categoryName, err := s.categoryRepo.GetCategoryNameByID(productReq.CategoryID)
	if err != nil {
		return err
	}

	merchant, err := s.merchantRepo.GetMerchantByAuthId(productReq.AuthID)
	if err != nil {
		return err
	}

	product := entity.Product{
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		Stock:       productReq.Stock,
		CategoryID:  productReq.CategoryID,
		ImageUrl:    productReq.ImageUrl,
		UpdatedAt:   time.Now(),
		MerchantID:  merchant.ID,
	}

	err = s.productRepo.UpdateProduct(product, productID)
	if err != nil {
		return err
	}

	err = s.meilisearch.UpdateData(product, categoryName, productID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) ListProducts(query string, limit, page int, authID int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error) {
	merchant, err := s.merchantRepo.GetMerchantByAuthId(authID)
	if err != nil {
		return []entity.MeilisearchPayloadResponse{}, entity.MeilisearchPagination{}, err
	}

	product, pagination, err := s.meilisearch.SearchQuery(query, limit, page, &merchant.ID)
	if err != nil {
		return []entity.MeilisearchPayloadResponse{}, entity.MeilisearchPagination{}, err
	}
	return product, pagination, nil
}

func (s *ProductService) ListProductsAll(query string, limit, page int) ([]entity.MeilisearchPayloadResponse, entity.MeilisearchPagination, error) {

	product, pagination, err := s.meilisearch.SearchQuery(query, limit, page, nil)
	if err != nil {
		return []entity.MeilisearchPayloadResponse{}, entity.MeilisearchPagination{}, err
	}
	return product, pagination, nil
}
