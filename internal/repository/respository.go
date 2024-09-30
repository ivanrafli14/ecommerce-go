package repository

import (
	"github.com.ivanrafli14/ecommerce-golang/pkg/database/mongodb"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	AuthRepository     IAuthRepository
	UserRepository     IUserRepository
	MerchantRepository IMerchantRepository
	ProductRepository  IProductRepository
	CategoryRepository ICategoryRepository
	OrderRepository    IOrderRepository
}

func NewRepository(db *sqlx.DB, mongodb mongodb.IMongoDB) Repository {
	authRepo := NewAuthRepository(db)
	userRepo := NewUserRepository(db)
	merchantRepo := NewMerchantRepository(db)
	productRepo := NewProductRepository(db)
	categoryRepo := NewCategoryRepository(db)
	orderRepo := NewOrderRepository(mongodb)

	return Repository{
		AuthRepository:     authRepo,
		UserRepository:     userRepo,
		MerchantRepository: merchantRepo,
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
		OrderRepository:    orderRepo,
	}
}
