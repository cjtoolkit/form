package form

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type validate struct {
	form
	_form FormInterface
	res   http.ResponseWriter
	req   *http.Request
	i18n  I18n
}

type formTypeError struct {
	DataType string
	FormType string
}

func (va validate) typeError() {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	temp, err := template.New("TypeError").Parse(va.i18n.Key(ErrType))
	if err != nil {
		va.setErr(err)
		return
	}
	temp.Execute(buf, formTypeError{va.t.Name(), va.ftype})
	va.setErr(FormError(buf.String()))
}

func (va validate) str() {
	switch va.ftype {
	case "input:text", "input:password", "input:search", "input:hidden",
		"input:url", "input:tel":
		va.strInputText()
	case "input:email":
		va.strInputEmail()
	case "input:radio":
		va.strInputRadio()
	case "input:color":
		va.strInputColor()
	case "textarea":
		va.strTextarea()
	case "select":
		va.strSelect()
	default:
		va.typeError()
	}
}

func (va validate) strs() {
	switch va.ftype {
	case "select":
		va.strsSelect()
	default:
		va.typeError()
	}
}

func (va validate) wnum() {
	switch va.ftype {
	case "input:number", "input:range", "input:hidden":
		va.numInputNumber(false)
	case "input:radio":
		va.wnumInputRadio()
	case "select":
		va.wnumSelect()
	default:
		va.typeError()
	}
}

func (va validate) wnums() {
	switch va.ftype {
	case "select":
		va.wnumsSelect()
	default:
		va.typeError()
	}
}

func (va validate) fnum() {
	switch va.ftype {
	case "input:number", "input:range", "input:hidden":
		va.numInputNumber(true)
	case "input:radio":
		va.fnumInputRadio()
	case "select":
		va.fnumSelect()
	default:
		va.typeError()
	}
}

func (va validate) fnums() {
	switch va.ftype {
	case "select":
		va.fnumsSelect()
	default:
		va.typeError()
	}
}

func (va validate) b() {
	switch va.ftype {
	case "input:checkbox", "input:hidden":
		va.bInputCheckbox()
	default:
		va.typeError()
	}
}

func (va validate) time() {
	switch va.ftype {
	case "input:datetime":
		va.timeInputTime("datetime")
	case "input:datetime-local":
		va.timeInputTime("datetime-local")
	case "input:time":
		va.timeInputTime("time")
	case "input:date":
		va.timeInputTime("date")
	case "input:month":
		va.timeInputTime("month")
	case "input:week":
		va.timeInputTime("week")
	default:
		va.typeError()
	}
}

func (va validate) file() {
	switch va.ftype {
	case "input:file":
		va.fileInputFile()
	default:
		va.typeError()
	}
}

func (va validate) setErr(err error) {
	va._form.SetErr(va.name, err)
}

func (va validate) callExt() {
	m := va.m.MethodByName(va.name + "Ext")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	m.Call(in)
}

// Validate by Struct Field. (Must be struct with pointer)
func ValidateItself(v FormInterface, res http.ResponseWriter, req *http.Request) bool {
	t := reflect.TypeOf(v)
	vc := reflect.ValueOf(v)
	vcc := vc

	switch {
	case isStructPtr(t):
		t = t.Elem()
		vc = vc.Elem()
	default:
		return false
	}

	lang := DefaultI18n

	for i := 0; i < t.NumField(); i++ {
		field := vc.Field(i)
		if !field.CanSet() {
			continue
		}
		name := t.Field(i).Name
		preferedName := name
		tag := t.Field(i).Tag.Get("form")
		if tag == "-" {
			continue
		}
		if tag != "" {
			preferedName = tag
		}

		va := validate{form{vcc, t, vc, t.Field(i), vc.Field(i), name, preferedName, ""}, v, res, req, lang}

		ftype, ok := va.getStr("Type")
		if !ok {
			continue
		}

		va.form.ftype = ftype

		if v.Err(name) != nil {
			// No point going further with validation, has it already failed!
			continue
		}

		switch vc.Field(i).Interface().(type) {
		case string:
			va.str()
		case []string:
			va.strs()
		case int64:
			va.wnum()
		case []int64:
			va.wnums()
		case float64:
			va.fnum()
		case []float64:
			va.fnums()
		case bool:
			va.b()
		case time.Time:
			va.time()
		case *multipart.FileHeader:
			va.file()
		}
	}

	// Tick it has checked
	v.Check()

	return !v.HasErr()
}

func ValidateItselfMulti(res http.ResponseWriter, req *http.Request, forms ...FormInterface) bool {
	b := true
	for _, form := range forms {
		if !ValidateItself(form, res, req) {
			b = false
		}
	}
	return b
}

// Read user request, populate struct field, than call ValidateItself (Must be struct with pointer)
func Validate(v FormInterface, res http.ResponseWriter, req *http.Request) bool {
	r := req
	r.ParseMultipartForm(10 * 1024 * 1024)
	valShift := func(name string) string {
		str := ""
		if r.MultipartForm != nil {
			if len(r.MultipartForm.Value[name]) <= 1 {
				if len(r.MultipartForm.Value[name]) == 1 {
					str = r.MultipartForm.Value[name][0]
				}
				delete(r.MultipartForm.Value, name)
			} else {
				str, r.MultipartForm.Value[name] =
					r.MultipartForm.Value[name][0], r.MultipartForm.Value[name][1:]
			}
		} else if r.PostForm != nil {
			if len(r.PostForm[name]) <= 1 {
				if len(r.PostForm[name]) == 1 {
					str = r.PostForm[name][0]
				}
				delete(r.PostForm, name)
			} else {
				str, r.PostForm[name] =
					r.PostForm[name][0], r.PostForm[name][1:]
			}
		} else if r.Form != nil {
			if len(r.Form[name]) <= 1 {
				if len(r.Form[name]) == 1 {
					str = r.Form[name][0]
				}
				delete(r.Form, name)
			} else {
				str, r.Form[name] =
					r.Form[name][0], r.Form[name][1:]
			}
		}
		return str
	}
	valAll := func(name string) []string {
		strs := []string{}
		if r.MultipartForm != nil {
			if len(r.MultipartForm.Value[name]) > 0 {
				strs = r.MultipartForm.Value[name]
				delete(r.MultipartForm.Value, name)
			}
		} else if r.PostForm != nil {
			if len(r.PostForm[name]) > 0 {
				strs = r.PostForm[name]
				delete(r.PostForm, name)
			}
		} else if r.Form != nil {
			if len(r.Form[name]) > 0 {
				strs = r.Form[name]
				delete(r.Form, name)
			}
		}
		return strs
	}
	fileShift := func(name string) *multipart.FileHeader {
		var fileHeader *multipart.FileHeader
		if r.MultipartForm != nil {
			if len(r.MultipartForm.File[name]) <= 1 {
				if len(r.MultipartForm.File[name]) == 1 {
					fileHeader = r.MultipartForm.File[name][0]
				}
				delete(r.MultipartForm.File, name)
			} else {
				fileHeader, r.MultipartForm.File[name] =
					r.MultipartForm.File[name][0], r.MultipartForm.File[name][1:]
			}
		}
		return fileHeader
	}

	t := reflect.TypeOf(v)
	vc := reflect.ValueOf(v)
	vcc := vc

	switch {
	case isStructPtr(t):
		t = t.Elem()
		vc = vc.Elem()
	default:
		return false
	}

	for i := 0; i < t.NumField(); i++ {
		field := vc.Field(i)
		if !field.CanSet() {
			continue
		}
		name := t.Field(i).Name
		preferedName := name
		tag := t.Field(i).Tag.Get("form")
		if tag == "-" {
			continue
		}
		if tag != "" {
			preferedName = tag
		}

		typeFunc := vcc.MethodByName(name + "Type")
		if !typeFunc.IsValid() {
			continue
		}

		switch vc.Field(i).Interface().(type) {
		case string:
			vc.Field(i).Set(reflect.ValueOf(valShift(preferedName)))
		case []string:
			vc.Field(i).Set(reflect.ValueOf(valAll(preferedName)))
		case int64:
			_v, err := strconv.ParseInt(valShift(preferedName), 10, 64)
			if err != nil {
				v.SetErr(name, err)
				continue
			}
			vc.Field(i).Set(reflect.ValueOf(_v))
		case []int64:
			vs := []int64{}
			vals := valAll(preferedName)
			var err error
			for _, val := range vals {
				var v int64
				v, err = strconv.ParseInt(val, 10, 64)
				if err != nil {
					break
				}
				vs = append(vs, v)
			}
			if err != nil {
				v.SetErr(name, err)
				continue
			}
			vc.Field(i).Set(reflect.ValueOf(vs))
		case float64:
			_v, err := strconv.ParseFloat(valShift(preferedName), 64)
			if err != nil {
				v.SetErr(name, err)
				continue
			}
			vc.Field(i).Set(reflect.ValueOf(_v))
		case []float64:
			vs := []float64{}
			vals := valAll(preferedName)
			var err error
			for _, val := range vals {
				var v float64
				v, err = strconv.ParseFloat(val, 64)
				if err != nil {
					break
				}
				vs = append(vs, v)
			}
			if err != nil {
				v.SetErr(name, err)
				continue
			}
			vc.Field(i).Set(reflect.ValueOf(vs))
		case bool:
			b := false
			if valShift(preferedName) == "1" {
				b = true
			}
			vc.Field(i).Set(reflect.ValueOf(b))
		case time.Time:
			in := make([]reflect.Value, 0)
			values := typeFunc.Call(in)
			if len(values) == 0 {
				continue
			}
			_type, ok := values[0].Interface().(string)
			if !ok {
				continue
			}

			_v := valShift(preferedName)

			var _time time.Time
			var err error

			if _v == "" {
				goto blank
			}

			switch _type {
			case "input:datetime":
				_time, err = time.Parse(dateTimeFormat, _v)
			case "input:datetime-local":
				_time, err = time.Parse(dateTimeLocalFormat, _v)
			case "input:time":
				_time, err = time.Parse(timeFormat, _v)
			case "input:date":
				_time, err = time.Parse(dateFormat, _v)
			case "input:month":
				_time, err = time.Parse(monthFormat, _v)
			case "input:week":
				_vv := strings.Split(_v, "-W")
				if len(_vv) < 2 {
					v.SetErr(name, FormError(DefaultI18n.Key(ErrOutOfBound)))
					continue
				}
				year, err := strconv.ParseInt(_vv[0], 10, 64)
				if err != nil {
					v.SetErr(name, err)
					continue
				}
				week, err := strconv.ParseInt(_vv[1], 10, 64)
				if err != nil {
					v.SetErr(name, err)
					continue
				}
				_time = StartingDayOfWeek(int(year), int(week))
			}

			if err != nil {
				v.SetErr(name, err)
				continue
			}

		blank:

			vc.Field(i).Set(reflect.ValueOf(_time))

		case *multipart.FileHeader:
			fileHeader := fileShift(preferedName)
			if fileHeader == nil {
				continue
			}
			vc.Field(i).Set(reflect.ValueOf(fileHeader))
		}
	}

	return ValidateItself(v, res, req)
}

func ValidateMulti(res http.ResponseWriter, req *http.Request, forms ...FormInterface) bool {
	b := true
	for _, form := range forms {
		if !Validate(form, res, req) {
			b = false
		}
	}
	return b
}
