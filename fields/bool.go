package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"strings"
)

type Bool struct {
	Name     string  // Mandatory
	Label    string  // Mandatory
	Norm     *string // Mandatory
	Model    *bool   // Mandatory
	Err      *error  // Mandatory
	Value    string  // Mandatory
	Required bool
}

type boolJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	Value    string `json:"value"`
}

func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(boolJson{
		Type:     "bool",
		Name:     b.Name,
		Required: b.Required,
		Success:  nil == *b.Err,
		Error:    getMessageFromError(*b.Err),
		Value:    b.Value,
	})
}

func (b Bool) PreCheck() {
	switch {
	case "" == strings.TrimSpace(b.Name):
		panic(form.ErrorPreCheck("Bool Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(b.Label):
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Label cannot be empty string"))
	case nil == b.Norm:
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Norm cannot be nil value"))
	case nil == b.Model:
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Model cannot be nil value"))
	case nil == b.Err:
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Err cannot be nil value"))
	case "" == strings.TrimSpace(b.Value):
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Value cannot be empty string"))
	}
}

func (b Bool) GetErrorPtr() *error {
	return b.Err
}

func (b Bool) PopulateNorm(values form.ValuesInterface) {
	*b.Norm = values.GetOne(b.Name)
}

func (b Bool) Transform() {
	if !*b.Model {
		return
	}
	*b.Norm = b.Value
}

func (b Bool) ReverseTransform() {
	*b.Model = strings.TrimSpace(*b.Norm) == b.Value
}

func (b Bool) ValidateModel() {
	b.validateRequired()
}

func (b Bool) validateRequired() {
	switch {
	case !b.Required:
		return
	case !*b.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": b.Label,
			},
		})
	}
}