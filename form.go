package form

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cjtoolkit/i18n"
	"io"
	"net/http"
	"time"
)

// Form Renderer and Validator interface!
type Form interface {
	// Renders 'formPtr' to 'w', panic if formPtr is not a struct with pointer.
	// Also renders validation errors if 'Validate' or 'MustValidate' was call before hand.
	Render(w io.Writer, formPtr FormPtr)
	// As Render but Return string
	RenderStr(formPtr FormPtr) string

	// Validate User Input and Populate Field in struct with pointers.
	// Must use struct with pointers otherwise it will return an error.
	// r cannot be 'nil'
	Validate(formPtrs ...FormPtr) (bool, error)
	// Same as validate but panic on error.
	MustValidate(formPtrs ...FormPtr) bool

	// Validate Single Field, won't work with must match.
	ValidateSingle(formPtr FormPtr, name string, value []string) (err error)

	// Encode JSON into 'w'
	// {"valid": bool, "data":[{"valid":bool, "error":"", "warning":"", "name":"", "count":int}...]}
	// Must call Validate or MustValidate first, otnilherwise it's print invalid data.
	Json(w io.Writer)

	// Set Location
	Location(loc *time.Location)
}

// Create new form validator and renderer.
// Panic if unable to verify languageSources or r is nil
// To use Default Second Layer specify rsl as 'nil'.
//
// Note: Stick to one instant per user request,
// do not use it as a global variable, as it's not thread safe.
func New(r *http.Request, rsl RenderSecondLayer, languageSources ...string) Form {
	if r == nil {
		panic(fmt.Errorf("Form: 'r' cannot be 'nil'"))
	}

	if rsl == nil {
		rsl = DefaultRenderSecondLayer
	}

	return &form{
		T:        i18n.MustTfunc("cjtoolkit-form", languageSources...),
		R:        rsl,
		Data:     map[FormPtr]*formData{},
		JsonData: []map[string]interface{}{},
		loc:      time.Local,
		req:      r,
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
	Data      map[FormPtr]*formData
	JsonValid bool
	JsonData  []map[string]interface{}
	Value     *value
	vcount    int
	rcount    int
	loc       *time.Location
	req       *http.Request
}

func (f *form) Render(w io.Writer, formPtr FormPtr) {
	f.render(formPtr, w)
}

func (f *form) RenderStr(formPtr FormPtr) string {
	w := &bytes.Buffer{}
	defer w.Reset()
	f.Render(w, formPtr)
	return w.String()
}

func (f *form) Validate(formPtrs ...FormPtr) (bool, error) {
	if f.Value == nil {
		f.Value = newValue(f.req)
	}
	valid := true
	for _, formPtr := range formPtrs {
		b, err := f.validate(formPtr)
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

func (f *form) MustValidate(formPtrs ...FormPtr) bool {
	b, err := f.Validate(formPtrs...)
	if err != nil {
		panic(err)
	}
	return b
}

func (f *form) ValidateSingle(formPtr FormPtr, name string, value []string) (err error) {
	err = f.validateSingle(formPtr, name, value)
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

func (f *form) Location(loc *time.Location) {
	if loc == nil {
		return
	}
	f.loc = loc
}
