package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"regexp"
	"strings"
	"unicode/utf8"
)

/*
Implement:
	FormFieldInterface in "github.com/cjtoolkit/form"
*/
type String struct {
	Name           string  // Mandatory
	Label          string  // Mandatory
	Norm           *string // Mandatory
	Model          *string // Mandatory
	Err            *error  // Mandatory
	Required       bool
	MinRune        int
	MaxRune        int
	MustMatchName  string
	MustMatchLabel string
	MustMatchModel *string
	Pattern        *regexp.Regexp
	PatternErrKey  string
	InList         []string
	Extra          func()
}

type stringJson struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Required  bool     `json:"required"`
	Success   bool     `json:"success"`
	Error     string   `json:"error,omitempty"`
	Min       int      `json:"min,omitempty"`
	Max       int      `json:"max,omitempty"`
	MustMatch string   `json:"mustMatch,omitempty"`
	Pattern   string   `json:"pattern,omitempty"`
	List      []string `json:"list,omitempty"`
}

func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(stringJson{
		Type:      "string",
		Name:      s.Name,
		Required:  s.Required,
		Success:   nil == *s.Err,
		Error:     getMessageFromError(*s.Err),
		Min:       s.MinRune,
		Max:       s.MaxRune,
		MustMatch: s.MustMatchName,
		Pattern:   getPatternFromRegExp(s.Pattern),
		List:      s.InList,
	})
}

func (s String) PreCheck() {
	switch {
	case "" == strings.TrimSpace(s.Name):
		panic(form.ErrorPreCheck("Text Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(s.Label):
		panic(form.ErrorPreCheck("Text Field: " + s.Name + ": Label cannot be empty string"))
	case nil == s.Norm:
		panic(form.ErrorPreCheck("Text Field: " + s.Name + ": Norm cannot be nil value"))
	case nil == s.Model:
		panic(form.ErrorPreCheck("Text Field: " + s.Name + ": Model cannot be nil value"))
	case nil == s.Err:
		panic(form.ErrorPreCheck("Text Field: " + s.Name + ": Err cannot be nil value"))
	}
}

func (s String) GetErrorPtr() *error {
	return s.Err
}

func (s String) PopulateNorm(values form.ValuesInterface) {
	*s.Norm = values.GetOne(s.Name)
}

func (s String) Transform() {
	*s.Norm = strings.TrimSpace(*s.Model)
}

func (s String) ReverseTransform() {
	*s.Model = strings.TrimSpace(*s.Norm)
}

func (s String) ValidateModel() {
	s.validateRequired()
	s.validateMinRune()
	s.validateMaxRune()
	s.validateMustMatch()
	s.validatePattern()
	s.validateInList()
	execFnIfNotNil(s.Extra)
}

func (s String) validateRequired() {
	switch {
	case !s.Required:
		return
	case "" == *s.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": s.Label,
			},
		})
	}
}

func (s String) validateMinRune() {
	switch {
	case 0 == s.MinRune:
		return
	case s.MinRune > utf8.RuneCountInString(*s.Model):
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MIN_CHAR,
			Value: map[string]interface{}{
				"Label":   s.Label,
				"MinRune": s.MinRune,
			},
		})
	}
}

func (s String) validateMaxRune() {
	switch {
	case 0 == s.MaxRune:
		return
	case s.MaxRune < utf8.RuneCountInString(*s.Model):
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MAX_CHAR,
			Value: map[string]interface{}{
				"Label":   s.Label,
				"MaxRune": s.MaxRune,
			},
		})
	}
}

func (s String) validateMustMatch() {
	switch {
	case nil == s.MustMatchModel || "" == s.MustMatchLabel:
		return
	case *s.MustMatchModel != *s.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_MUST_MATCH,
			Value: map[string]interface{}{
				"Label":          s.Label,
				"MustMatchLabel": s.MustMatchLabel,
			},
		})
	}
}

func (s String) validatePattern() {
	switch {
	case nil == s.Pattern:
		return
	case !s.Pattern.MatchString(*s.Model):
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_PATTERN, s.PatternErrKey),
			Value: map[string]interface{}{
				"Label":   s.Label,
				"Pattern": s.Pattern.String(),
			},
		})
	}
}

func (s String) validateInList() {
	if nil == s.InList {
		return
	}

	model := *s.Model

	for _, value := range s.InList {
		if model == value {
			return
		}
	}

	panic(&form.ErrorValidateModel{
		Key: form.LANG_IN_LIST,
		Value: map[string]interface{}{
			"Label": s.Label,
			"List":  s.InList,
		},
	})
}
