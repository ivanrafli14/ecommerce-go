package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) Login(c *gin.Context) {
	var authReq entity.AuthRequest
	if err := c.ShouldBindJSON(&authReq); err != nil {
		err = response.ErrorBadRequest
		response.Failed(c, err, "bad request")
		return
	}

	res, err := r.service.AuthService.Login(authReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, 200, "login success", res)
}

func (r *Rest) Register(c *gin.Context) {
	var authReq entity.AuthRequest
	if err := c.ShouldBindJSON(&authReq); err != nil {
		err = response.ErrorBadRequest
		response.Failed(c, err, "bad request")
		return
	}
	err := r.service.AuthService.Register(authReq)
	if err != nil {

		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, 201, "register success", nil)
}

func (r *Rest) UpdateRole(c *gin.Context) {
	authID := c.GetInt("AuthID")
	err := r.service.AuthService.UpdateRole(authID)

	if err != nil {

		if err.Error() == "user already as a merchant" {
			newError := response.NewError("bad request", "40001", http.StatusBadRequest)
			response.ErrorMapping[err.Error()] = newError
			response.Failed(c, err, err.Error())
		} else {
			myErr := response.ErrGeneral
			response.Failed(c, myErr, err.Error())
		}
		return

	}
	c.Set("AuthRole", "merchant")
	response.Success(c, 200, "update role success", nil)
}
