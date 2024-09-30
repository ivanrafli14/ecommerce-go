package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/internal/service"
	"github.com.ivanrafli14/ecommerce-golang/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	route := r.router.Group("/v1")

	auth := route.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login", r.Login)
	auth.PATCH("/role", r.middleware.Authentication, r.UpdateRole)

	user := route.Group("/users")
	user.Use(r.middleware.Authentication)
	user.Use(r.middleware.Authorization("user"))
	user.POST("/profile", r.CreateProfile)
	user.PUT("/profile", r.UpdateProfile)
	user.GET("/profile", r.GetProfile)

	merchant := route.Group("/merchants")
	merchant.Use(r.middleware.Authentication)
	merchant.Use(r.middleware.Authorization("merchant"))
	merchant.POST("/profile", r.CreateMerchant)
	merchant.GET("/profile", r.GetMerchant)
	merchant.PUT("/profile", r.UpdateMerchant)

	product := route.Group("/products")
	product.Use(r.middleware.Authentication)
	product.GET("/detail/:sku", r.middleware.Authorization("user"), r.GetProductBySKU)

	product.Use(r.middleware.Authorization("merchant"))
	product.GET("", r.ListProduct)
	product.POST("", r.CreateProduct)
	product.PUT("/id/:product_id", r.UpdateProduct)
	product.GET("/id/:product_id", r.GetProductByID)

	category := route.Group("/categories")
	category.GET("", r.GetAllCategories)

	order := route.Group("/orders")
	order.Use(r.middleware.Authentication)
	order.POST("", r.middleware.Authorization("user"), r.Checkout)
	order.GET("/merchant", r.middleware.Authorization("merchant"), r.GetOrderMerchant)
	order.GET("/user", r.middleware.Authorization("user"), r.GetOrderUser)

	route.POST("/orders/webhook", r.WebhookOrder)

	route.POST("/files/image", r.middleware.Authentication, r.UploadImage)
	route.GET("/search/products", r.SearchProduct)

}

func (r *Rest) Serve() {
	r.router.Run(":8081")
}
