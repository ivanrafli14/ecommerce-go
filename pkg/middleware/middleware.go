package middleware

import (
	"github.com.ivanrafli14/ecommerce-golang/internal/repository"
	"github.com.ivanrafli14/ecommerce-golang/pkg/jwt"
	"github.com.ivanrafli14/ecommerce-golang/pkg/redis"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	Authentication(c *gin.Context)
	Authorization(allowedRole string) gin.HandlerFunc
}
type middleware struct {
	redis    redis.Interface
	jwtAuth  jwt.Interface
	authRepo repository.IAuthRepository
}

func Init(jwtAuth jwt.Interface, redis redis.Interface, authRepo repository.IAuthRepository) Interface {
	return &middleware{
		redis:    redis,
		jwtAuth:  jwtAuth,
		authRepo: authRepo,
	}
}
