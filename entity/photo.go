package entity

import "mime/multipart"

type PhotoRequest struct {
	File   *multipart.FileHeader `form:"file"`
	AuthID int                   `form:"-" json:"-"`
}
