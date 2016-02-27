package fields

import (
	"github.com/cjtoolkit/form"
	"strings"
	"unicode/utf8"
)

/*
Implement:
	FormFieldInterface in "github.com/cjtoolkit/form"
*/
type Text struct {
	Name           string  // Mandatory
	Label          string  // Mandatory
	Norm           *string // Mandatory
	Model          *string // Mandatory
	Err            *error  // Mandatory
	Required       bool
	MinChar        int
	MaxChar        int
	MustMatchLabel string
	MustMatchModel *string
}

func (t Text) PreCheck() {
	switch {
	case "" == strings.TrimSpace(t.Name):
		panic(form.ErrorPreCheck("Text Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(t.Label):
		panic(form.ErrorPreCheck("Text Field: " + t.Name + ": Label cannot be empty string"))
	case nil == t.Norm:
		panic(form.ErrorPreCheck("Text Field: " + t.Name + ": Norm cannot be nil value"))
	case nil == t.Model:
		panic(form.ErrorPreCheck("Text Field: " + t.Name + ": Model cannot be nil value"))
	case nil == t.Err:
		panic(form.ErrorPreCheck("Text Field: " + t.Name + ": Err cannot be nil value"))
	}
}

func (t Text) GetErrorPtr() *error {
	return t.Err
}

func (t Text) PopulateNorm(values form.ValuesInterface) {
	*t.Norm = values.GetOne(t.Name)
}

func (t Text) Transform() {
	*t.Norm = strings.TrimSpace(*t.Model)
}

func (t Text) ReverseTransform() {
	*t.Model = strings.TrimSpace(*t.Norm)
}

func (t Text) ValidateModel() {
	t.validateRequired()
	t.validateMinChar()
	t.validateMaxChar()
	t.validateMustMatch()
}

func (t Text) validateRequired() {
	if !t.Required {
		return
	}

	if "" == *t.Model {
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": t.Label,
			},
		})
	}
}

func (t Text) validateMinChar() {
	if 0 == t.MinChar {
		return
	}

	if t.MinChar > utf8.RuneCountInString(*t.Model) {
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MIN_CHAR,
			Value: map[string]interface{}{
				"Label":   t.Label,
				"MinChar": t.MinChar,
			},
		})
	}
}

func (t Text) validateMaxChar() {
	if 0 == t.MaxChar {
		return
	}

	if t.MaxChar < utf8.RuneCountInString(*t.Model) {
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MAX_CHAR,
			Value: map[string]interface{}{
				"Label":   t.Label,
				"MaxChar": t.MaxChar,
			},
		})
	}
}

func (t Text) validateMustMatch() {
	if nil == t.MustMatchModel && "" == t.MustMatchLabel {
		return
	}

	if *t.MustMatchModel != *t.Model {
		// Error
	}
}
