package form

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"time"
)

type renderValue struct {
	form         *form
	name         string
	preferedName string
	fieldsFns    FieldFuncs
	_type        TypeCode
	w            io.Writer
}

func (f *form) render(structPtr interface{}, w io.Writer) {
	t := reflect.TypeOf(structPtr)
	vc := reflect.ValueOf(structPtr)
	vcc := vc

	switch {
	case isStructPtr(t):
		t = t.Elem()
		vc = vc.Elem()
	default:
		panic(fmt.Errorf("form: '%p' is not a struct pointer", structPtr))
	}

	data := f.Data[structPtr]

	buf := &bytes.Buffer{}

	for fieldNo := 0; fieldNo < t.NumField(); fieldNo++ {
		field := vc.Field(fieldNo)
		if !field.CanSet() {
			continue
		}

		name := t.Field(fieldNo).Name
		preferedName := name

		opsFunc := vcc.MethodByName(name + "Field")
		if !opsFunc.IsValid() {
			continue
		}

		val := opsFunc.Call(make([]reflect.Value, 0))
		var fieldFns FieldFuncs
		var ok bool
		if fieldFns, ok = val[0].Interface().(FieldFuncs); !ok {
			continue
		}

		_type := Invalid

		fieldFns.Call("form", map[string]interface{}{
			"type": &_type,
			"name": &preferedName,
		})

		if _type <= Invalid || _type >= terminate {
			continue
		}

		// First Layer

		r := renderValue{f, name, preferedName, fieldFns, _type, buf}

		switch value := field.Interface().(type) {
		case string:
			r.str(value)
		case []string:
			r.strs(value)
		case int64:
			r.wnum(value)
		case []int64:
			r.wnums(value)
		case float64:
			r.fnum(value)
		case []float64:
			r.fnums(value)
		case bool:
			r.b(value)
		case time.Time:
			r.time(value)
		case *multipart.FileHeader:
			r.file()
		default:
			continue
		}

		// Second Layer

		var err error
		warning := ""

		if data != nil {
			err = data.Errors[r.name]
			warning = data.Warning[r.name]
		}

		secondLayerData := RenderData{f.rcount, _type, err, warning, buf.String(), fieldFns, data != nil}
		f.rcount++
		buf.Reset()

		f.R.Render(w, secondLayerData)
	}
}

func (r renderValue) str(value string) {
	switch r._type {
	case InputText, InputSearch, InputPassword, InputHidden, InputUrl, InputTel:
		r.strInputText(value)
	case InputEmail:
		r.strInputEmail(value)
	case InputRadio:
		r.strInputRadio(value)
	case InputColor:
		r.strInputColor(value)
	case Textarea:
		r.strTextarea(value)
	case Select:
		r.strSelect(value)
	}
}

func (r renderValue) strs(values []string) {
	switch r._type {
	case Select:
		r.strsSelect(values)
	}
}

func (r renderValue) wnum(value int64) {
	switch r._type {
	case InputNumber, InputRange, InputHidden:
		r.numInputNumber(value)
	case InputRadio:
		r.wnumInputRadio(value)
	case Select:
		r.wnumSelect(value)
	}
}

func (r renderValue) wnums(values []int64) {
	switch r._type {
	case Select:
		r.wnumsSelect(values)
	}
}

func (r renderValue) fnum(value float64) {
	switch r._type {
	case InputNumber, InputRange, InputHidden:
		r.numInputNumber(value)
	case InputRadio:
		r.fnumInputRadio(value)
	case Select:
		r.fnumSelect(value)
	}
}

func (r renderValue) fnums(values []float64) {
	switch r._type {
	case Select:
		r.fnumsSelect(values)
	}
}

func (r renderValue) b(value bool) {
	switch r._type {
	case InputCheckbox, InputHidden:
		r.bInputCheckbox(value)
	}
}

func (r renderValue) time(value time.Time) {
	switch r._type {
	case InputDatetime, InputDatetimeLocal, InputTime, InputDate, InputMonth, InputWeek:
		r.timeInputTime(value)
	}
}

func (r renderValue) file() {
	switch r._type {
	case InputFile:
		r.fileInputFile()
	}
}
