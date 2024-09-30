package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) CreateProfile(c *gin.Context) {
	var userReq entity.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		myErr := response.ErrorBadRequest
		response.Failed(c, myErr, "bad request")
		return
	}
	userReq.AuthID = c.GetInt("AuthID")
	userReq.Role = c.GetString("AuthRole")
	err := r.service.UserService.CreateUser(userReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "create user success", nil)

}

func (r *Rest) UpdateProfile(c *gin.Context) {
	var userReq entity.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		myErr := response.ErrorBadRequest
		response.Failed(c, myErr, "bad request")
		return
	}
	userReq.AuthID = c.GetInt("AuthID")
	err := r.service.UserService.UpdateUser(userReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "update user success", nil)

}

func (r *Rest) GetProfile(c *gin.Context) {
	authID := c.GetInt("AuthID")
	user, err := r.service.UserService.GetUser(authID)

	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "get user success", user)

}
