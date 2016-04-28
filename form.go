package form

import (
	"fmt"
	"mime/multipart"
	"net/url"
)

type Form struct {
	language        LanguageInterface
	values          ValuesInterface
	disablePreCheck bool
}

func NewForm(language LanguageInterface) (f *Form) {
	f = &Form{
		language: language,
	}
	f.checkLanguage()
	return
}

func NewFormEnglishLanguage() *Form {
	return NewForm(englishLanguageMap)
}

func (f *Form) DisablePreCheck() *Form {
	f.disablePreCheck = true
	return f
}

func (f *Form) checkLanguage() {
	if nil == f.language {
		panic("Form language cannot be set to nil")
	}
}

func (f *Form) SetForm(form url.Values) {
	f.values = newValues(form)
	f.checkValues()
}

func (f *Form) SetMultipartForm(form *multipart.Form) {
	f.values = newValuesFile(form)
	f.checkValues()
}

func (f *Form) checkValues() {
	if nil == f.values {
		panic("SetForm or SetMultipartForm cannot be nil value.")
	}
}

func (f *Form) handleError(errPtr *error) {
	switch r := recover().(type) {
	case nil:
		// Not an error do nothing.
	case ErrorTransform: // Should be 500 error, not the client fault.  Usually the developers fault.
		panic(r)
	case TranslatableErrorInterface: // It's either that or a more complex FormFieldInterface (No thanks!)
		r.Translate(f.language)
		*errPtr = r
	case error:
		*errPtr = r
	default:
		*errPtr = ErrorUnknown(fmt.Sprint(r))
	}
}

func (f *Form) checkErrorInLoop(err error, success *bool) {
	if nil != err {
		*success = false
	}
}

func (f *Form) checkForm(form FormBuilderInterface) {
	if nil == form {
		panic("'form' cannot be nil")
	}
}

func (f *Form) checkField(field FormFieldInterface) {
	if nil == field {
		panic("'field' cannot be nil")
	}
	if !f.disablePreCheck {
		field.PreCheck()
	}
}

func (f *Form) checkErrPtr(errPtr *error) {
	if nil == errPtr {
		panic("'errPtr' cannot be nil")
	}
}

func (f *Form) transform(errPtr *error, field FormFieldInterface) {
	defer f.handleError(errPtr)
	field.Transform()
}

// Panic if form and field are 'nil' or if field fails pre check.
func (f *Form) Transform(form FormBuilderInterface) bool {
	f.checkForm(form)
	success := true

	for _, field := range form.Fields() {
		f.checkField(field)
		errPtr := field.GetErrorPtr()
		f.checkErrPtr(errPtr)
		f.transform(errPtr, field)
		f.checkErrorInLoop(*errPtr, &success)
	}

	return success
}

// Panic if form and field are 'nil' or if field fails pre check.
func (f *Form) TransformSingle(field FormFieldInterface) error {
	f.checkField(field)
	errPtr := field.GetErrorPtr()
	f.checkErrPtr(errPtr)
	f.transform(errPtr, field)
	return *errPtr
}

func (f *Form) validate(errPtr *error, field FormFieldInterface) {
	defer f.handleError(errPtr)
	*errPtr = nil
	if nil != f.values {
		field.PopulateNorm(f.values)
		field.ReverseTransform()
	}
	field.ValidateModel()
}

// Panic if values, form and field are 'nil' or if field fails pre check.
func (f *Form) Validate(form FormBuilderInterface) bool {
	f.checkForm(form)
	success := true

	for _, field := range form.Fields() {
		f.checkField(field)
		errPtr := field.GetErrorPtr()
		f.checkErrPtr(errPtr)
		f.validate(errPtr, field)
		f.checkErrorInLoop(*errPtr, &success)
	}

	return success
}

// Panic if values, form and field are 'nil' or if field fails pre check.
func (f *Form) ValidateSingle(field FormFieldInterface) error {
	f.checkField(field)
	errPtr := field.GetErrorPtr()
	f.checkErrPtr(errPtr)
	f.validate(errPtr, field)
	return *errPtr
}
