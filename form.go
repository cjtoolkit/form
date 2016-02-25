package form

import (
	"mime/multipart"
	"net/url"
)

type Form struct {
	language Langauge
	values   ValuesInterface
}

func NewForm(language Langauge) *Form {
	return &Form{
		language: language,
	}
}

func NewFormDefaultLanguage() *Form {
	return NewForm(DefaultLanguage())
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
		panic("SetForm or SetMultipartForm has not been called or nil value has been passed to either.")
	}
}

func (f *Form) handleError(errPtr *error) {
	switch r := recover().(type) {
	case TranslatableErrorInterface: // It's either that or a more complex FormFieldInterface (No thanks!)
		r.Translate(f.language)
		*errPtr = r
	case error:
		*errPtr = r
	}
}

func (f *Form) transform(errPtr *error, field FormFieldInterface) {
	defer f.handleError(errPtr)
	field.Transform()
}

func (f *Form) checkErrorInLoop(err error, success *bool) {
	if nil == err {
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
}

func (f *Form) checkErrPtr(errPtr *error) {
	if nil == errPtr {
		panic("'errPtr' cannot be nil")
	}
}

func (f *Form) Transform(form FormBuilderInterface) bool {
	f.checkValues()
	f.checkForm(form)
	success := true

	for _, field := range form.BuildForm() {
		f.checkField(field)
		errPtr := field.GetErrorPtr()
		f.checkErrPtr(errPtr)
		f.transform(errPtr, field)
		f.checkErrorInLoop(*errPtr, &success)
	}

	return success
}

func (f *Form) TransformSingle(field FormFieldInterface) error {
	f.checkValues()
	f.checkField(field)
	errPtr := field.GetErrorPtr()
	f.checkErrPtr(errPtr)
	f.transform(errPtr, field)
	return *errPtr
}

func (f *Form) validate(errPtr *error, field FormFieldInterface) {
	defer f.handleError(errPtr)
	field.PopulateNorm(f.values)
	field.ReverseTransform()
	field.ValidateModel()
}

func (f *Form) Validate(form FormBuilderInterface) bool {
	f.checkValues()
	f.checkForm(form)
	success := true

	for _, field := range form.BuildForm() {
		f.checkField(field)
		errPtr := field.GetErrorPtr()
		f.checkErrPtr(errPtr)
		f.validate(errPtr, field)
		f.checkErrorInLoop(*errPtr, &success)
	}

	return success
}

func (f *Form) ValidateSingle(field FormFieldInterface) error {
	f.checkValues()
	f.checkField(field)
	errPtr := field.GetErrorPtr()
	f.checkErrPtr(errPtr)
	f.validate(errPtr, field)
	return *errPtr
}
