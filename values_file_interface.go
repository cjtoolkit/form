package form

import (
	"mime/multipart"
)

type ValuesFileInterface interface {
	ValuesInterface
	GetOneFile(name string) *multipart.FileHeader
	GetAllFile(name string) []*multipart.FileHeader
}
