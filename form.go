package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cjtoolkit/i18n"
	"io"
	"net/http"
)

// Form Renderer and Validator interface!
type Form interface {
	// Renders 'structPtr' to 'w', panic if structPtr is not a struct with pointer.
	// Also renders validation errors if 'Validate' or 'MustValidate' was call before hand.
	Render(structPtr StructPtrForm, w io.Writer)
	// As Render but Return string
	RenderStr(structPtr StructPtrForm) string

	// Validate User Input and Populate Field in struct with pointers.
	// Must use struct with pointers otherwise it will return an error.
	// r cannot be 'nil'
	Validate(r *http.Request, structPtrs ...StructPtrForm) (bool, error)
	// Same as validate but panic on error.
	MustValidate(r *http.Request, structPtrs ...StructPtrForm) bool

	// Validate Single Field, won't work with must match.
	ValidateSingle(structPtr StructPtrForm, name string, value []string) (err error)

	// Encode JSON into 'w'
	// {"valid": bool, "data":[{"valid":bool, "error":"", "warning":"", "name":"", "count":int}...]}
	// Must call Validate or MustValidate first, otnilherwise it's print invalid data.
	Json(w io.Writer)
}

// Create new form validator and renderer.
// Panic if unable to verify languageSources.
// To use Default Second Layer specify r as 'nil'.
//
// Note: Stick to one instant per user request,
// do not use it as a global variable, as it's not thread safe.
func New(r RenderSecondLayer, languageSources ...string) Form {
	if r == nil {
		r = DefaultRenderSecondLayer
	}

	return &form{
		T:        i18n.MustTfunc("cjtoolkit-form", languageSources...),
		R:        r,
		Data:     map[StructPtrForm]*formData{},
		JsonData: []map[string]interface{}{},
	}
}

type formData struct {
	Errors  map[string][]error
	Warning map[string][]string
}

func newData() *formData {
	return &formData{
		Errors:  map[string][]error{},
		Warning: map[string][]string{},
	}
}

func (f *formData) addError(name string, err error) {
	f.Errors[name] = append(f.Errors[name], err)
}

func (f *formData) addWarning(name, warning string) {
	f.Warning[name] = append(f.Warning[name], warning)
}

func (f *formData) shiftError(name string) (err error) {
	if len(f.Errors[name]) > 0 {
		if len(f.Errors[name]) <= 1 {
			if len(f.Errors[name]) == 1 {
				err = f.Errors[name][0]
			}
			delete(f.Errors, name)
		} else {
			err, f.Errors[name] = f.Errors[name][0], f.Errors[name][1:]
		}
	}
	return
}

func (f *formData) shiftWarning(name string) (warning string) {
	if len(f.Warning[name]) > 0 {
		if len(f.Warning[name]) <= 1 {
			if len(f.Warning[name]) == 1 {
				warning = f.Warning[name][0]
			}
			delete(f.Warning, name)
		} else {
			warning, f.Warning[name] = f.Warning[name][0], f.Warning[name][1:]
		}
	}
	return
}

type form struct {
	T         i18n.Translator
	R         RenderSecondLayer
	Data      map[StructPtrForm]*formData
	JsonValid bool
	JsonData  []map[string]interface{}
	Value     *value
	vcount    int
	rcount    int
}

func (f *form) Render(structPtr StructPtrForm, w io.Writer) {
	f.render(structPtr, w)
}

func (f *form) RenderStr(structPtr StructPtrForm) string {
	w := &bytes.Buffer{}
	defer w.Reset()
	f.Render(structPtr, w)
	return w.String()
}

func (f *form) Validate(r *http.Request, structPtrs ...StructPtrForm) (bool, error) {
	if r == nil {
		return false, fmt.Errorf("Form: 'r' cannot be 'nil'")
	}
	if f.Value == nil {
		f.Value = newValue(r)
	}
	valid := true
	for _, structPtr := range structPtrs {
		b, err := f.validate(structPtr)
		if err != nil {
			return false, err
		}
		if !b {
			valid = false
		}
	}
	f.JsonValid = valid
	return valid, nil
}

func (f *form) MustValidate(r *http.Request, structPtrs ...StructPtrForm) bool {
	b, err := f.Validate(r, structPtrs...)
	if err != nil {
		panic(err)
	}
	return b
}

func (f *form) ValidateSingle(structPtr StructPtrForm, name string, value []string) (err error) {
	err = f.validateSingle(structPtr, name, value)
	return
}

func (f *form) Json(w io.Writer) {
	v := map[string]interface{}{
		"valid": f.JsonValid,
		"data":  f.JsonData,
	}
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
