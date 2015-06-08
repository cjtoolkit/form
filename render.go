package form

import (
	"fmt"
	"io"
	"mime/multipart"
	"time"
)

type renderValue struct {
	form         *form
	name         string
	preferedName string
	fieldsFns    FieldFuncs
	_type        TypeCode
	fls          *FirstLayerStack
}

func (f *form) render(formPtr FormPtr, w io.Writer) {
	if !isStructPtr(formPtr) {
		panic(fmt.Errorf("form: '%T' is not a pointer", formPtr))
	}

	data := f.Data[formPtr]

	fields := &Fields{}

	fields.m = map[string]FieldFuncs{}
	fields.f = []*Field{}
	fields.R = f.req

	formPtr.CJForm(fields)

	for _, field := range fields.f {
		name := field.name
		fieldFns := field.funcs

		preferedName := name
		_type := field.typecode

		if _type <= Invalid || _type >= terminate {
			continue
		}

		// First Layer

		r := renderValue{f, name, preferedName, fieldFns, _type, &FirstLayerStack{}}

		switch value := field.ptr.(type) {
		case *string:
			r.str(value)
		case *[]string:
			r.strs(value)
		case *int64:
			r.wnum(value)
		case *[]int64:
			r.wnums(value)
		case *float64:
			r.fnum(value)
		case *[]float64:
			r.fnums(value)
		case *bool:
			r.b(value)
		case *time.Time:
			r.time(value)
		case **multipart.FileHeader:
			r.file()
		default:
			continue
		}

		// Second Layer

		var err error
		warning := ""

		if data != nil {
			err = data.shiftError(r.name)
			warning = data.shiftWarning(r.name)
		}

		secondLayerData := RenderData{preferedName, f.rcount, _type, err, warning,
			fieldFns, data != nil, *(r.fls)}
		f.rcount++

		f.R.Render(w, secondLayerData)
	}
}

func (r renderValue) str(value *string) {
	switch r._type {
	case InputText, InputSearch, InputPassword, InputHidden, InputUrl, InputTel:
		r.strInputText(*value)
	case InputEmail:
		r.strInputEmail(*value)
	case InputRadio:
		r.strInputRadio(*value)
	case InputColor:
		r.strInputColor(*value)
	case Textarea:
		r.strTextarea(*value)
	case Select:
		r.strSelect(*value)
	}
}

func (r renderValue) strs(values *[]string) {
	switch r._type {
	case Select:
		r.strsSelect(*values)
	}
}

func (r renderValue) wnum(value *int64) {
	switch r._type {
	case InputNumber, InputRange, InputHidden:
		r.numInputNumber(*value)
	case InputRadio:
		r.wnumInputRadio(*value)
	case Select:
		r.wnumSelect(*value)
	}
}

func (r renderValue) wnums(values *[]int64) {
	switch r._type {
	case Select:
		r.wnumsSelect(*values)
	}
}

func (r renderValue) fnum(value *float64) {
	switch r._type {
	case InputNumber, InputRange, InputHidden:
		r.numInputNumber(*value)
	case InputRadio:
		r.fnumInputRadio(*value)
	case Select:
		r.fnumSelect(*value)
	}
}

func (r renderValue) fnums(values *[]float64) {
	switch r._type {
	case Select:
		r.fnumsSelect(*values)
	}
}

func (r renderValue) b(value *bool) {
	switch r._type {
	case InputCheckbox, InputHidden:
		r.bInputCheckbox(*value)
	}
}

func (r renderValue) time(value *time.Time) {
	switch r._type {
	case InputDatetime, InputDatetimeLocal, InputTime, InputDate, InputMonth, InputWeek:
		r.timeInputTime(*value)
	}
}

func (r renderValue) file() {
	switch r._type {
	case InputFile:
		r.fileInputFile()
	}
}
