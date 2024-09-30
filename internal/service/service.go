package service

import (
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/bcrypt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/cloudinary"
	"github.com.ivanrafli14/ecommerce-golang/pkg/jwt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/meilisearch"
	"github.com.ivanrafli14/ecommerce-golang/pkg/payment_gateway"
	"github.com.ivanrafli14/ecommerce-golang/pkg/redis"
)

type Service struct {
	AuthService     IAuthService
	UserService     IUserService
	MerchantService IMerchantService
	ProductService  IProductService
	CategoryService ICategoryService
	OrderService    IOrderService
	PhotoService    IPhotoService
}

type InitParam struct {
	Repository     repository.Repository
	Bcrypt         bcrypt.Interface
	Jwt            jwt.Interface
	Redis          redis.Interface
	Meilisearch    meilisearch.SearchEngine
	PaymentGateway payment_gateway.IPaymentGateway
	Cloudinary     cloudinary.ICloudinary
}

func NewService(param InitParam) *Service {
	authService := NewAuthService(param.Repository.AuthRepository, param.Bcrypt, param.Jwt, param.Redis)
	userService := NewUserService(param.Repository.UserRepository)
	merchantService := NewMerchantService(param.Repository.MerchantRepository)
	productService := NewProductService(param.Repository.ProductRepository, param.Repository.MerchantRepository, param.Repository.CategoryRepository, param.Meilisearch)
	categoryService := NewCategoryService(param.Repository.CategoryRepository)
	orderService := NewOrderService(param.Repository.OrderRepository, param.Repository.MerchantRepository, param.Repository.ProductRepository, param.Repository.AuthRepository, param.Repository.UserRepository, param.PaymentGateway, param.Meilisearch)
	photoService := NewPhotoService(param.Cloudinary)

	return &Service{
		AuthService:     authService,
		UserService:     userService,
		MerchantService: merchantService,
		ProductService:  productService,
		CategoryService: categoryService,
		OrderService:    orderService,
		PhotoService:    photoService,
	}
}
