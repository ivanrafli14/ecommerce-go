package main

import (
	"github.com.ivanrafli14/ecommerce-golang/internal/hander/rest"
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/internal/service"
	"github.com.ivanrafli14/ecommerce-golang/pkg/bcrypt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/cloudinary"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com.ivanrafli14/ecommerce-golang/pkg/database"
	"github.com.ivanrafli14/ecommerce-golang/pkg/database/mongodb"
	"github.com.ivanrafli14/ecommerce-golang/pkg/jwt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/meilisearch"
	"github.com.ivanrafli14/ecommerce-golang/pkg/middleware"
	"github.com.ivanrafli14/ecommerce-golang/pkg/payment_gateway"
	"github.com.ivanrafli14/ecommerce-golang/pkg/redis"
)

func main() {
	cfg := config.LoadConfig("cmd/api/config.yaml")
	db := database.ConnectDB(cfg.Database)
	jwtPkg := jwt.NewJWT(cfg.JWT)
	bcryptPkg := bcrypt.Init()
	redisPkg := redis.NewRedis(cfg.Redis)
	meilisearchPkg := meilisearch.NewMeilisearch(cfg.MeiliSearch)

	mongodbPkg := mongodb.NewMongoDBClient(cfg.MongoDB)
	xenditPkg := payment_gateway.NewXendit(cfg.PaymentGateway)
	cloudinaryPkg := cloudinary.NewCloudinaryClient(cfg.Cloudinary)

	repo := repository.NewRepository(db, mongodbPkg)
	middlewarePkg := middleware.Init(jwtPkg, redisPkg, repo.AuthRepository)
	svc := service.NewService(service.InitParam{
		Repository:     repo,
		Bcrypt:         bcryptPkg,
		Jwt:            jwtPkg,
		Redis:          redisPkg,
		Meilisearch:    meilisearchPkg,
		PaymentGateway: xenditPkg,
		Cloudinary:     cloudinaryPkg,
	})
	r := rest.NewRest(svc, middlewarePkg)
	r.MountEndpoint()
	r.Serve()

}
