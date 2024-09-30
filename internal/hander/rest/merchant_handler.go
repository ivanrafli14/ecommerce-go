package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *Rest) CreateMerchant(c *gin.Context) {
	var merchantReq entity.MerchantRequest
	if err := c.ShouldBindJSON(&merchantReq); err != nil {
		err = response.ErrBadRequest
		response.Failed(c, err, err.Error())
		return
	}

	merchantReq.AuthID = c.GetInt("AuthID")
	merchantReq.Role = c.GetString("AuthRole")
	log.Println(merchantReq.Role)
	err := r.service.MerchantService.CreateMerchant(&merchantReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "create merchants success", nil)
	return

}

func (r *Rest) GetMerchant(c *gin.Context) {
	authID := c.GetInt("AuthID")
	merchant, err := r.service.MerchantService.GetMerchantByAuthID(authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "get merchant success", merchant)
}

func (r *Rest) UpdateMerchant(c *gin.Context) {
	var merchantReq entity.MerchantRequest
	if err := c.ShouldBindJSON(&merchantReq); err != nil {
		err = response.ErrBadRequest
		response.Failed(c, err, err.Error())
	}
	merchantReq.AuthID = c.GetInt("AuthID")
	err := r.service.MerchantService.UpdateMerchant(&merchantReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "update merchant success", nil)
}
