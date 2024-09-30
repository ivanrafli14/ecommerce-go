package middleware

import (
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *middleware) Authentication(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		myErr := response.ErrJWTEmpty
		response.Failed(c, myErr, myErr.Error())
		c.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	AuthId, _, err := m.jwtAuth.VerifyJWTToken(token)
	if err != nil {
		response.Failed(c, err, err.Error())
		c.Abort()
		return
	}

	_, err = m.redis.GetData(token)

	if err != nil {
		response.Failed(c, err, err.Error())
		c.Abort()
		return
	}

	auth, err := m.authRepo.GetAuthByID(AuthId)
	if err != nil {
		response.Failed(c, err, err.Error())
		c.Abort()
		return
	}

	c.Set("AuthID", AuthId)
	c.Set("AuthRole", auth.Role)
	c.Next()
}
