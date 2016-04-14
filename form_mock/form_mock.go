package form_mock

import (
	"github.com/cjtoolkit/form"
	. "github.com/smartystreets/goconvey/convey"
	"mime/multipart"
	"net/url"
)

type FormMock struct {
	C C

	setFormParamForm chan url.Values

	setFormMultipartFormParamForm chan *multipart.Form

	transformParamForm  chan form.FormBuilderInterface
	transformWillReturn chan bool

	transformSingleParamField chan form.FormFieldInterface
	transformSingleWillReturn chan error

	validateParamForm  chan form.FormBuilderInterface
	validateWillReturn chan bool

	validateSingleParamFiled chan form.FormFieldInterface
	validateSingleWillReturn chan error
}

func NewFormMock() *FormMock {
	return &FormMock{
		setFormParamForm: make(chan url.Values),

		setFormMultipartFormParamForm: make(chan *multipart.Form),

		transformParamForm:  make(chan form.FormBuilderInterface),
		transformWillReturn: make(chan bool),

		transformSingleParamField: make(chan form.FormFieldInterface),
		transformSingleWillReturn: make(chan error),

		validateParamForm:  make(chan form.FormBuilderInterface),
		validateWillReturn: make(chan bool),

		validateSingleParamFiled: make(chan form.FormFieldInterface),
		validateSingleWillReturn: make(chan error),
	}
}

func (f *FormMock) ExpectSetForm(expectForm url.Values) {
	f.setFormParamForm <- expectForm
}

func (f *FormMock) SetForm(form url.Values) {
	f.C.So(form, ShouldResemble, <-f.setFormParamForm)
}

func (f *FormMock) ExpectSetMultipartForm(expectForm *multipart.Form) {
	f.setFormMultipartFormParamForm <- expectForm
}

func (f *FormMock) SetMultipartForm(form *multipart.Form) {
	f.C.So(form, ShouldResemble, <-f.setFormMultipartFormParamForm)
}

func (f *FormMock) ExpectTransform(expectForm form.FormBuilderInterface, willReturn bool) {
	f.transformParamForm <- expectForm
	f.transformWillReturn <- willReturn
}

func (f *FormMock) Transform(form form.FormBuilderInterface) bool {
	f.C.So(form, ShouldResemble, <-f.transformParamForm)

	return <-f.transformWillReturn
}

func (f *FormMock) ExpectTransformSingle(expectField form.FormFieldInterface, willReturn error) {
	f.transformSingleParamField <- expectField
	f.transformSingleWillReturn <- willReturn
}

func (f *FormMock) TransformSingle(field form.FormFieldInterface) error {
	f.C.So(field, ShouldResemble, <-f.transformSingleParamField)

	return <-f.transformSingleWillReturn
}

func (f *FormMock) ExpectValidate(expectForm form.FormBuilderInterface, willReturn bool) {
	f.validateParamForm <- expectForm
	f.validateWillReturn <- willReturn
}

func (f *FormMock) Validate(form form.FormBuilderInterface) bool {
	f.C.So(form, ShouldResemble, <-f.validateParamForm)

	return <-f.validateWillReturn
}

func (f *FormMock) ExpectValidateSingle(expectField form.FormFieldInterface, willReturn error) {
	f.validateSingleParamFiled <- expectField
	f.validateSingleWillReturn <- willReturn
}

func (f *FormMock) ValidateSingle(field form.FormFieldInterface) error {
	f.C.So(field, ShouldResemble, <-f.validateSingleParamFiled)

	return <-f.validateSingleWillReturn
}
