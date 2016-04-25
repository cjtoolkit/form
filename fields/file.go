package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"mime/multipart"
	"strings"
)

type File struct {
	Name       string                 // Mandatory
	Label      string                 // Mandatory
	File       **multipart.FileHeader // Mandatory
	Err        *error                 // Mandatory
	Required   bool
	Mime       []string
	SizeInByte int64
	Extra      func()
}

type fileJson struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Required bool     `json:"required"`
	Mime     []string `json:"mime"`
	Size     int64    `json:"size"`
}

const (
	FILE_CONTENT_TYPE = "Content-Type"
)

func (f File) MarshalJSON() ([]byte, error) {
	return json.Marshal(fileJson{
		Type:     "file",
		Name:     f.Name,
		Required: f.Required,
		Mime:     f.Mime,
		Size:     f.SizeInByte,
	})
}

func (f File) PreCheck() {
	switch {
	case "" == strings.TrimSpace(f.Name):
		panic(form.ErrorPreCheck("File Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(f.Label):
		panic(form.ErrorPreCheck("File Field: " + f.Name + ": Label cannot be empty string"))
	case nil == f.File:
		panic(form.ErrorPreCheck("File Field: " + f.Name + ": File cannot be nil value"))
	case nil == f.Err:
		panic(form.ErrorPreCheck("File Field: " + f.Name + ": Err cannot be nil value"))
	}
}

func (f File) GetErrorPtr() *error {
	return f.Err
}

func (f File) PopulateNorm(values form.ValuesInterface) {
	*f.File = form.GetOneFile(values, f.Name)
}

func (f File) Transform() {
	// Do nothing
}

func (f File) ReverseTransform() {
	// Do nothing
}

func (f File) ValidateModel() {
	f.validateRequired()
	f.validateMime()
	f.validateSizeInByte()
	execFnIfNotNil(f.Extra)
}

func (f File) validateRequired() {
	switch {
	case !f.Required:
		return
	case nil == *f.File:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": f.Label,
			},
		})
	}
}

func (f File) validateMime() {

}

func (f File) getFileSize() (size int64) {
	if nil == *f.File {
		return
	}

	file, err := (*f.File).Open()
	if nil != err {
		return
	}
	defer file.Close()

	size, err = file.Seek(0, 2)
	if nil != err {
		size = 0
		return
	}

	file.Seek(0, 0)

	return
}

func (f File) validateSizeInByte() {

}
