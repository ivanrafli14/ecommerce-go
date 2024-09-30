package middleware

import (
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
)

func (m* middleware) Authorization(allowedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("AuthRole")
		if role != allowedRole {
			err:= response.ErrRoleInvalid
			response.Failed(c,err, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
