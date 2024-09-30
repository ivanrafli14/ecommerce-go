package cloudinary

import (
	"context"
	"github.com.ivanrafli14/ecommerce-golang/entity"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"log"
	"strconv"
)

type ICloudinary interface {
	UploadPhoto(photoReq entity.PhotoRequest) (string, error)
}

type CloudinaryClient struct {
	client *cloudinary.Cloudinary
}

func NewCloudinaryClient(config config.CloudinaryConfig) ICloudinary {
	cldSecret := config.APISecret
	cldName := config.Name
	cldKey := config.APIKey

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		log.Fatal(err)
	}

	return &CloudinaryClient{
		client: cld,
	}
}

func (c *CloudinaryClient) UploadPhoto(photoReq entity.PhotoRequest) (string, error) {
	log.Println(photoReq.AuthID)
	filename := uuid.NewString()
	authIDStr := strconv.Itoa(photoReq.AuthID)
	res, err := c.client.Upload.Upload(context.Background(), photoReq.File, uploader.UploadParams{
		Folder:   "E-Commerce/" + authIDStr,
		PublicID: filename,
	})

	if err != nil {
		return "", err
	}
	log.Println(res)

	if len(res.Eager) > 0 {
		// will return secure url with transformation
		return res.Eager[0].SecureURL, nil
	}

	url := res.SecureURL

	return url, nil
}
