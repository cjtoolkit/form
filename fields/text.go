package fields

import (
	"github.com/cjtoolkit/form"
	"strings"
)

/*
Implement:
	FormFieldInterface in "github.com/cjtoolkit/form"
*/
type Text struct {
	Name  string
	Label string
	Norm  *string
	Model *string
	Err   *error
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

}
