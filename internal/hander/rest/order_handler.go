package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (r *Rest) Checkout(c *gin.Context) {
	var checkoutReq entity.CheckoutRequest
	if err := c.ShouldBindJSON(&checkoutReq); err != nil {
		response.Failed(c, err, "bad request")
		return
	}
	authID := c.GetInt("AuthID")

	invoiceUrl, err := r.service.OrderService.Checkout(checkoutReq, authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "create order success", invoiceUrl)

}

func (r *Rest) GetOrderMerchant(c *gin.Context) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "1"
	}
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	authID := c.GetInt("AuthID")
	orders, pagination, err := r.service.OrderService.ListOrdersMerchant(limit, page, authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.SuccessWithPagination(c, http.StatusOK, "get orders success", orders, pagination)
}

func (r *Rest) GetOrderUser(c *gin.Context) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "5"
	}
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	authID := c.GetInt("AuthID")
	orders, pagination, err := r.service.OrderService.ListOrdersUser(limit, page, authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.SuccessWithPagination(c, http.StatusOK, "get orders success", orders, pagination)
}

func (r *Rest) WebhookOrder(c *gin.Context) {
	var webhookReq entity.WebhookInvoiceRequest
	if err := c.ShouldBindJSON(&webhookReq); err != nil {
		response.Failed(c, err, "bad request")
		return
	}
	err := r.service.OrderService.WebhookOrder(webhookReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "webhook success", nil)

}
