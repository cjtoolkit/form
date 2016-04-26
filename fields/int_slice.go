package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"strings"
	"strconv"
	"sort"
)

type IntSlice struct {
	Name     string    // Mandatory
	Label    string    // Mandatory
	Norm     *[]string // Mandatory
	Model    *[]int64  // Mandatory
	Err      *error    // Mandatory
	Required bool
	Extra    func()
}

type intSliceJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
}

const (
	INT_SLICE_DECIMAL = 10
	INT_SLICE_BIT = 64
)

func (i IntSlice) MarshalJSON() ([]byte, error) {
	return json.Marshal(intSliceJson{
		Type:     "int_slice",
		Name:     i.Name,
		Required: i.Required,
		Success:  nil == *i.Err,
		Error:    getMessageFromError(*i.Err),
	})
}

func (i IntSlice) PreCheck() {
	switch {
	case "" == strings.TrimSpace(i.Name):
		panic(form.ErrorPreCheck("IntSlice Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(i.Label):
		panic(form.ErrorPreCheck("IntSlice Field: " + i.Name + ": Label cannot be empty string"))
	case nil == i.Norm:
		panic(form.ErrorPreCheck("IntSlice Field: " + i.Name + ": Norm cannot be nil value"))
	case nil == i.Model:
		panic(form.ErrorPreCheck("IntSlice Field: " + i.Name + ": Model cannot be nil value"))
	case nil == i.Err:
		panic(form.ErrorPreCheck("IntSlice Field: " + i.Name + ": Err cannot be nil value"))
	}
}

func (i IntSlice) GetErrorPtr() *error {
	return i.Err
}

func (i IntSlice) PopulateNorm(values form.ValuesInterface) {
	*i.Norm = values.GetAll(i.Name)
}

func (i IntSlice) Transform() {
	*i.Norm = nil
	for _, num := range *i.Model {
		*i.Norm = append(*i.Norm, strconv.FormatInt(num, INT_SLICE_DECIMAL))
	}
	sort.Strings(*i.Norm)
}

func (i IntSlice) ReverseTransform() {
	*i.Model = nil
	for _, str := range *i.Norm {
		num, err := strconv.ParseInt(strings.TrimSpace(str), INT_SLICE_DECIMAL, INT_SLICE_BIT)
		ExecFuncIfErrIsNotNil(err, func() {
			panic(&form.ErrorReverseTransform{
				Key: form.LANG_NOT_INT,
				Value: map[string]interface{}{
					"Label": i.Label,
				},
			})
		})
		*i.Model = append(*i.Model, num)
	}
	sort.Sort(Int64Sort(*i.Model))
}

func (i IntSlice) ValidateModel() {
	i.validateRequired()
	execFnIfNotNil(i.Extra)
}

func (i IntSlice) validateRequired() {
	switch {
	case !i.Required:
		return
	case 0 == len(*i.Model):
		panic(&form.ErrorValidateModel{
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": i.Label,
			},
		})
	}
}