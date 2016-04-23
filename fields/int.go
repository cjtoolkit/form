package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"strconv"
	"strings"
)

/*
Implement:
	FormFieldInterface in "github.com/cjtoolkit/form"
*/
type Int struct {
	Name     string  // Mandatory
	Label    string  // Mandatory
	Norm     *string // Mandatory
	Model    *int64  // Mandatory
	Err      *error  // Mandatory
	Required bool
	Min      int64
	MinZero  bool
	Max      int64
	MaxZero  bool
	Step     int64
	InList   []int64
	Extra    func()
}

type intJson struct {
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Required bool    `json:"required"`
	Success  bool    `json:"success"`
	Error    string  `json:"error,omitempty"`
	Min      int64   `json:"min,omitempty"`
	MinZero  bool    `json:"minZero,omitempty"`
	Max      int64   `json:"max,omitempty"`
	MaxZero  bool    `json:"maxZero,omitempty"`
	Step     int64   `json:"step,omitempty"`
	List     []int64 `json:"list,omitempty"`
}

const (
	INT_DECIMAL = 10
	INT_BIT     = 64
)

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(intJson{
		Type:     "int",
		Name:     i.Name,
		Required: i.Required,
		Success:  nil == *i.Err,
		Min:      i.Min,
		MinZero:  i.MinZero,
		Max:      i.Max,
		MaxZero:  i.MaxZero,
		Step:     i.Step,
		List:     i.InList,
	})
}

func (i Int) PreCheck() {
	switch {
	case "" == strings.TrimSpace(i.Name):
		panic(form.ErrorPreCheck("Int Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(i.Label):
		panic(form.ErrorPreCheck("Int Field: " + i.Name + ": Label cannot be empty string"))
	case nil == i.Norm:
		panic(form.ErrorPreCheck("Int Field: " + i.Name + ": Norm cannot be nil value"))
	case nil == i.Model:
		panic(form.ErrorPreCheck("Int Field: " + i.Name + ": Model cannot be nil value"))
	case nil == i.Err:
		panic(form.ErrorPreCheck("Int Field: " + i.Name + ": Err cannot be nil value"))
	}
}

func (i Int) GetErrorPtr() *error {
	return i.Err
}

func (i Int) PopulateNorm(values form.ValuesInterface) {
	*i.Norm = values.GetOne(i.Name)
}

func (i Int) Transform() {
	*i.Norm = strconv.FormatInt(*i.Model, INT_DECIMAL)
}

func (i Int) ReverseTransform() {
	num, err := strconv.ParseInt(strings.TrimSpace(*i.Norm), INT_BIT, INT_DECIMAL)
	ExecFuncIfErrIsNotNil(err, func() {
		panic(&form.ErrorReverseTransform{
			Key: form.LANG_NOT_INT,
			Value: map[string]interface{}{
				"Label": i.Label,
			},
		})
	})
	*i.Model = num
}

func (i Int) ValidateModel() {
	i.validateRequired()
	i.validateMin()
	i.validateMax()
	i.validateStep()
	i.validateInList()
	execFnIfNotNil(i.Extra)
}

func (i Int) validateRequired() {
	switch {
	case !i.Required:
		return
	case 0 == *i.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": i.Label,
			},
		})
	}
}

func (i Int) validateMin() {
	switch {
	case 0 == i.Min && !i.MinZero:
		return
	case i.Min > *i.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_MIN,
			Value: map[string]interface{}{
				"Label": i.Label,
				"Min":   i.Min,
			},
		})
	}
}

func (i Int) validateMax() {
	switch {
	case 0 == i.Max && !i.MaxZero:
		return
	case i.Max < *i.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_MAX,
			Value: map[string]interface{}{
				"Label": i.Label,
				"Max":   i.Max,
			},
		})
	}
}

func (i Int) validateStep() {
	switch {
	case 0 == i.Step:
		return
	case 0 != *i.Model%i.Step:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_STEP,
			Value: map[string]interface{}{
				"Label": i.Label,
				"Step":  i.Step,
			},
		})
	}
}

func (i Int) validateInList() {
	if nil == i.InList {
		return
	}

	model := *i.Model

	for _, value := range i.InList {
		if model == value {
			return
		}
	}

	panic(&form.ErrorValidateModel{
		Key: form.LANG_IN_LIST,
		Value: map[string]interface{}{
			"Label": i.Label,
			"List":  i.InList,
		},
	})
}
