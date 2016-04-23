package fields

import (
	"github.com/cjtoolkit/form"
	"math"
	"strconv"
	"strings"
	"encoding/json"
)

/*
Implement:
	FormFieldInterface in "github.com/cjtoolkit/form"
*/
type Float struct {
	Name     string   // Mandatory
	Label    string   // Mandatory
	Norm     *string  // Mandatory
	Model    *float64 // Mandatory
	Err      *error   // Mandatory
	Required bool
	Min      float64
	MinZero  bool
	Max      float64
	MaxZero  bool
	Step     float64
	InList   []float64
	Extra    func()
}

type floatJson struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Required bool      `json:"required"`
	Success  bool      `json:"success"`
	Error    string    `json:"error,omitempty"`
	Min      float64   `json:"min,omitempty"`
	MinZero  bool      `json:"minZero,omitempty"`
	Max      float64   `json:"max,omitempty"`
	MaxZero  bool      `json:"maxZero,omitempty"`
	Step     float64   `json:"step,omitempty"`
	List     []float64 `json:"list,omitempty"`
}

const (
	FLOAT_BIT                  = 64
	FLOAT_FMT_NO_EXPONENT byte = 'f'
	FLOAT_PRECISION            = -1
)

func (f Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(floatJson{
		Type:     "float",
		Name:     f.Name,
		Required: f.Required,
		Success:  nil == *f.Err,
		Error:    getMessageFromError(*f.Err),
		Min:      f.Min,
		MinZero:  f.MinZero,
		Max:      f.Max,
		MaxZero:  f.MaxZero,
		Step:     f.Step,
		List:     f.InList,
	})
}

func (f Float) PreCheck() {
	switch {
	case "" == strings.TrimSpace(f.Name):
		panic(form.ErrorPreCheck("Float Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(f.Label):
		panic(form.ErrorPreCheck("Float Field: " + f.Name + ": Label cannot be empty string"))
	case nil == f.Norm:
		panic(form.ErrorPreCheck("Float Field: " + f.Name + ": Norm cannot be nil value"))
	case nil == f.Model:
		panic(form.ErrorPreCheck("Float Field: " + f.Name + ": Model cannot be nil value"))
	case nil == f.Err:
		panic(form.ErrorPreCheck("Float Field: " + f.Name + ": Err cannot be nil value"))
	}
}

func (f Float) GetErrorPtr() *error {
	return f.Err
}

func (f Float) PopulateNorm(values form.ValuesInterface) {
	*f.Norm = strings.TrimSpace(values.GetOne(f.Name))
}

func (f Float) Transform() {
	*f.Norm = strconv.FormatFloat(*f.Model, FLOAT_FMT_NO_EXPONENT, FLOAT_PRECISION, FLOAT_BIT)
}

func (f Float) ReverseTransform() {
	num, err := strconv.ParseFloat(strings.TrimSpace(*f.Norm), FLOAT_BIT)
	ExecFuncIfErrIsNotNil(err, func() {
		panic(&form.ErrorReverseTransform{
			Key: form.LANG_NOT_FLOAT,
			Value: map[string]interface{}{
				"Label": f.Label,
			},
		})
	})
	*f.Model = num
}

func (f Float) ValidateModel() {
	f.validateRequired()
	execFnIfNotNil(f.Extra)
}

func (f Float) validateRequired() {
	switch {
	case !f.Required:
		return
	case 0 == *f.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": f.Label,
			},
		})
	}
}

func (f Float) validateMin() {
	switch {
	case 0 == f.Min && !f.MinZero:
		return
	case f.Min > *f.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_MIN,
			Value: map[string]interface{}{
				"Label": f.Label,
				"Min":   f.Min,
			},
		})
	}
}

func (f Float) validateMax() {
	switch {
	case 0 == f.Max && !f.MaxZero:
		return
	case f.Max < *f.Model:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_MAX,
			Value: map[string]interface{}{
				"Label": f.Label,
				"Max":   f.Max,
			},
		})
	}
}

func (f Float) validateStep() {
	num := math.Mod(*f.Model, f.Step)
	switch {
	case 0 == f.Step:
		return
	case 0 != num || math.NaN() == num:
		panic(&form.ErrorValidateModel{
			Key: form.LANG_NUMBER_STEP,
			Value: map[string]interface{}{
				"Label": f.Label,
				"Step":  f.Step,
			},
		})
	}
}

func (f Float) validateInList() {
	if nil == f.InList {
		return
	}

	model := *f.Model

	for _, value := range f.InList {
		if model == value {
			return
		}
	}

	panic(&form.ErrorValidateModel{
		Key: form.LANG_IN_LIST,
		Value: map[string]interface{}{
			"Label": f.Label,
			"List":  f.InList,
		},
	})
}