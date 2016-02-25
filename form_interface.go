package form

import (
	"mime/multipart"
	"net/url"
)

type FormInterface interface {
	SetForm(form url.Values)
	SetMultipartForm(form *multipart.Form)
	Transform(form FormBuilderInterface) bool
	TransformSingle(field FormFieldInterface) error
	Validate(form FormBuilderInterface) bool
	ValidateSingle(field FormFieldInterface) error
}
