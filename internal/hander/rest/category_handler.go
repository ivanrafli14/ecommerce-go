package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) GetAllCategories(c *gin.Context) {
	categories, err := r.service.CategoryService.GetAllCategories()
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "get categories success", categories)
}
