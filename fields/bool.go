package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"strings"
)

type Bool struct {
	Name           string // Mandatory
	Label          string // Mandatory
	Norm           string // Mandatory
	Model          *bool  // Mandatory
	Err            error  // Mandatory
	Value          string // Mandatory
	Suffix         *string
	Required       bool
	RequiredErrKey string
	Bridge         *form.FieldBridge
}

func NewBool(name, label, value string, model *bool, options ...BoolOption) *Bool {
	b := &Bool{
		Name:  name,
		Label: label,
		Value: value,
		Model: model,
	}

	b.PreCheck()

	b.Bridge = form.NewFieldBridge(form.InputCheckbox, false, &b.Label, &b.Name, &b.Suffix, &b.Value, nil)

	for _, option := range options {
		option(b)
	}

	return b
}

type boolJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	Value    string `json:"value"`
}

func (b *Bool) NameWithSuffix() string {
	return addSuffix(b.Name, b.Suffix)
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(boolJson{
		Type:     "bool",
		Name:     b.Name,
		Required: b.Required,
		Success:  nil == *b.Err,
		Error:    getMessageFromError(*b.Err),
		Value:    b.Value,
	})
}

func (b *Bool) PreCheck() {
	switch {
	case "" == strings.TrimSpace(b.Name):
		panic(form.ErrorPreCheck("Bool Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(b.Label):
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Label cannot be empty string"))
	case nil == b.Model:
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Model cannot be nil value"))
	case "" == strings.TrimSpace(b.Value):
		panic(form.ErrorPreCheck("Bool Field: " + b.Name + ": Value cannot be empty string"))
	}
}

func (b *Bool) GetErrorPtr() *error {
	return &b.Err
}

func (b *Bool) PopulateNorm(values form.ValuesInterface) {
	b.Norm = values.GetOne(b.NameWithSuffix())
}

func (b *Bool) Transform() {
	b.Norm = ""
	b.Bridge.SetChecked(false)
	if !*b.Model {
		return
	}
	b.Norm = b.Value
	b.Bridge.SetChecked(true)
}

func (b *Bool) ReverseTransform() {
	*b.Model = strings.TrimSpace(b.Norm) == b.Value
}

func (b *Bool) ValidateModel() {
	b.validateRequired()
}

func (b *Bool) validateRequired() {
	switch {
	case !b.Required:
		return
	case !*b.Model:
		panic(&form.ErrorValidateModel{
			Key:   UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_FIELD_REQUIRED, b.RequiredErrKey),
			Value: b,
		})
	}
}

func (b *Bool) GetBridge() *form.FieldBridge {
	return b.Bridge
}
