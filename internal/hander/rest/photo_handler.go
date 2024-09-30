package rest

import (
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Rest) UploadImage(c *gin.Context) {
	var photoReq entity.PhotoRequest

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to get file"})
		return
	}

	// Manually populate the Photo struct
	photoReq.File = file
	photoReq.AuthID = c.GetInt("AuthID")

	url, err := r.service.PhotoService.UploadPhoto(photoReq)
	if err != nil {
		response.Failed(c, err, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "upload image success", url)
}
