package fields

import (
	"github.com/cjtoolkit/form"
	"regexp"
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
	Pattern        *regexp.Regexp
	PatternErrKey  string
	InList         []string
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
	t.validatePattern()
	t.validateInList()
}

func (t Text) validateRequired() {
	switch {
	case !t.Required:
		return
	case "" == *t.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": t.Label,
			},
		})
	}
}

func (t Text) validateMinChar() {
	switch {
	case 0 == t.MinChar:
		return
	case t.MinChar > utf8.RuneCountInString(*t.Model):
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
	switch {
	case 0 == t.MaxChar:
		return
	case t.MaxChar < utf8.RuneCountInString(*t.Model):
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
	switch {
	case nil == t.MustMatchModel || "" == t.MustMatchLabel:
		return
	case *t.MustMatchModel != *t.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MUST_MATCH,
			Value: map[string]interface{}{
				"Label":          t.Label,
				"MustMatchLabel": t.MustMatchLabel,
			},
		})
	}
}

func (t Text) validatePattern() {
	switch {
	case nil == t.Pattern:
		return
	case !t.Pattern.MatchString(*t.Model):
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_PATTERN, t.PatternErrKey),
			Value: map[string]interface{}{
				"Label":   t.Label,
				"Pattern": t.Pattern.String(),
			},
		})
	}
}

func (t Text) validateInList() {
	if nil == t.InList {
		return
	}

	model := *t.Model

	for _, value := range t.InList {
		if model == value {
			return
		}
	}

	panic(&form.ErrorValidateModel{
		Key: form.LANG_IN_LIST,
		Value: map[string]interface{}{
			"Label": t.Label,
			"List":  t.InList,
		},
	})
}
