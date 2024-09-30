package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (r *Rest) CreateProduct(c *gin.Context) {
	var productReq entity.ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {
		myErr := response.ErrGeneral
		response.Failed(c, myErr, myErr.Error())
		return
	}
	productReq.AuthID = c.GetInt("AuthID")
	productReq.Role = c.GetString("AuthRole")

	err := r.service.ProductService.CreateProduct(productReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "create product success", nil)
}

func (r *Rest) UpdateProduct(c *gin.Context) {
	var productReq entity.ProductRequest
	if err := c.ShouldBindJSON(&productReq); err != nil {
		myErr := response.ErrGeneral
		response.Failed(c, myErr, myErr.Error())
		return
	}
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		myErr := response.ErrGeneral
		response.Failed(c, myErr, myErr.Error())
		return
	}
	productReq.AuthID = c.GetInt("AuthID")
	productReq.Role = c.GetString("AuthRole")

	err = r.service.ProductService.UpdateProduct(productReq, productID)
	if err != nil {
		log.Println(err.Error())
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "update product success", nil)
}

func (r *Rest) GetProductByID(c *gin.Context) {
	authID := c.GetInt("AuthID")
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		log.Println(err)
		myErr := response.ErrGeneral
		response.Failed(c, myErr, myErr.Error())
		return
	}

	product, err := r.service.ProductService.GetProductByProductID(productID, authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "get products success", product)
}

func (r *Rest) ListProduct(c *gin.Context) {
	query := c.Query("query")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	authID := c.GetInt("AuthID")
	if limitStr == "" {
		limitStr = "10"
	}

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

	products, pagination, err := r.service.ProductService.ListProducts(query, limit, page, authID)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.SuccessWithPagination(c, http.StatusOK, "get products success", products, pagination)

}

func (r *Rest) GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")
	product, err := r.service.ProductService.GetProductBySKU(sku)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "get products success", product)
}

func (r *Rest) SearchProduct(c *gin.Context) {
	query := c.Query("query")
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	if limitStr == "" {
		limitStr = "10"
	}

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

	products, pagination, err := r.service.ProductService.ListProductsAll(query, limit, page)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.SuccessWithPagination(c, http.StatusOK, "get products success", products, pagination)

}
