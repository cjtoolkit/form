package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"sort"
	"strings"
)

type StringSlice struct {
	Name           string    // Mandatory
	Label          string    // Mandatory
	Norm           *[]string // Mandatory
	Model          *[]string // Mandatory
	Err            *error    // Mandatory
	Suffix         *string
	Required       bool
	RequiredErrKey string
	Extra          func()
}

func NewStringSlice(name, label string, norm, model *[]string, err *error, options ...func(*StringSlice)) StringSlice {
	s := StringSlice{
		Name:  name,
		Label: label,
		Norm:  norm,
		Model: model,
		Err:   err,
	}

	s.PreCheck()

	for _, option := range options {
		option(&s)
	}

	return s
}

type stringSliceJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
}

func (s StringSlice) NameWithSuffix() string {
	return addSuffix(s.Name, s.Suffix)
}

func (s StringSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(stringSliceJson{
		Type:     "string_slice",
		Name:     s.Name,
		Required: s.Required,
		Success:  nil == *s.Err,
		Error:    getMessageFromError(*s.Err),
	})
}

func (s StringSlice) PreCheck() {
	switch {
	case "" == strings.TrimSpace(s.Name):
		panic(form.ErrorPreCheck("StringSlice Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(s.Label):
		panic(form.ErrorPreCheck("StringSlice Field: " + s.Name + ": Label cannot be empty string"))
	case nil == s.Norm:
		panic(form.ErrorPreCheck("StringSlice Field: " + s.Name + ": Norm cannot be nil value"))
	case nil == s.Model:
		panic(form.ErrorPreCheck("StringSlice Field: " + s.Name + ": Model cannot be nil value"))
	case nil == s.Err:
		panic(form.ErrorPreCheck("StringSlice Field: " + s.Name + ": Err cannot be nil value"))
	}
}

func (s StringSlice) GetErrorPtr() *error {
	return s.Err
}

func (s StringSlice) PopulateNorm(values form.ValuesInterface) {
	*s.Norm = values.GetAll(s.NameWithSuffix())
}

func (s StringSlice) Transform() {
	*s.Norm = nil
	for _, str := range *s.Model {
		*s.Norm = append(*s.Norm, strings.TrimSpace(str))
	}
	sort.Strings(*s.Norm)
}

func (s StringSlice) ReverseTransform() {
	*s.Model = nil
	for _, str := range *s.Norm {
		*s.Model = append(*s.Model, strings.TrimSpace(str))
	}
	sort.Strings(*s.Model)
}

func (s StringSlice) ValidateModel() {
	s.validateRequired()
	execFnIfNotNil(s.Extra)
}

func (s StringSlice) validateRequired() {
	switch {
	case !s.Required:
		return
	case 0 == len(*s.Model):
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_FIELD_REQUIRED, s.RequiredErrKey),
			Value: map[string]interface{}{
				"Label": s.Label,
			},
		})
	}
}
