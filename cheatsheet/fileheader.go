package cheatsheet

import (
	"mime/multipart"
)

/*
InputFile
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"file":
		"size": *int64
		"sizeErr": *string
		"accept": *[]string
		"acceptErr": *string
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type FileHeader *multipart.FileHeader
