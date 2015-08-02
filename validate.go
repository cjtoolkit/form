package form

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type validateValue struct {
	form         *form
	err          *error
	warning      *string
	name         string
	preferedName string
	fieldsFns    FieldFuncs
	_type        TypeCode
	t            string
}

func (f *form) validate(formPtr FormPtr) (bool, error) {
	if !isStructPtr(formPtr) {
		return false, fmt.Errorf("form: '%T' is not a pointer", formPtr)
	}

	f.Data[formPtr] = newData()
	data := f.Data[formPtr]

	type _f struct {
		ptr          interface{}
		name         string
		preferedName string
		fieldFns     FieldFuncs
		_type        TypeCode
		err          error
	}

	fieldM := []*_f{}

	fields := &Fields{}

	fields.m = map[string]FieldFuncs{}
	fields.f = []*Field{}
	fields.validating = true
	fields.R = f.req

	formPtr.CJForm(fields)

	for _, field := range fields.f {
		name := field.name
		fieldFns := field.funcs
		var err error

		var suffix func() []interface{}
		fieldFns.Call("suffix", map[string]interface{}{
			"suffix": &suffix,
		})

		if suffix != nil {
			name += fmt.Sprint(suffix()...)
		}

		preferedName := name

		_type := field.typecode

		if _type <= Invalid || _type >= terminate {
			continue
		}

		fieldC := &_f{field.ptr, name, preferedName, fieldFns, _type, nil}

		fieldM = append(fieldM, fieldC)

		if f.Value == nil {
			continue
		}

		switch ptr := field.ptr.(type) {

		case *string:
			*ptr = strings.TrimSpace(f.Value.Shift(preferedName))

		case *[]string:
			clean := []string{}
			unclean := f.Value.All(preferedName)
			for _, value := range unclean {
				clean = append(clean, strings.TrimSpace(value))
			}
			*ptr = clean

		case *int64:
			var _v int64
			_v, err = strconv.ParseInt(strings.TrimSpace(f.Value.Shift(preferedName)), 10, 64)
			if err != nil {
				continue
			}
			*ptr = _v

		case *[]int64:
			vs := []int64{}
			vals := f.Value.All(preferedName)
			for _, val := range vals {
				var v int64
				v, err = strconv.ParseInt(strings.TrimSpace(val), 10, 64)
				if err != nil {
					break
				}
				vs = append(vs, v)
			}
			if err != nil {
				continue
			}
			*ptr = vs

		case *float64:
			var _v float64
			_v, err = strconv.ParseFloat(strings.TrimSpace(f.Value.Shift(preferedName)), 64)
			if err != nil {
				continue
			}
			*ptr = _v

		case *[]float64:
			vs := []float64{}
			vals := f.Value.All(preferedName)
			for _, val := range vals {
				var v float64
				v, err = strconv.ParseFloat(strings.TrimSpace(val), 64)
				if err != nil {
					break
				}
				vs = append(vs, v)
			}
			if err != nil {
				continue
			}
			*ptr = vs

		case *bool:
			b := false
			if strings.TrimSpace(f.Value.Shift(preferedName)) == "1" {
				b = true
			}
			*ptr = b

		case *time.Time:
			_v := strings.TrimSpace(f.Value.Shift(preferedName))

			var _time time.Time

			parser := func(formats ...string) {
				for _, format := range formats {
					_time, err = time.ParseInLocation(format, _v, f.loc)
					if err == nil {
						return
					}
				}
			}

			if _v == "" {
				goto blank
			}

			switch _type {
			case InputDatetime:
				_time, err = time.Parse(dateTimeLocalFormat+".05Z07:00", _v)
				if err != nil {
					_time, err = time.Parse(dateTimeLocalFormat+":05Z07:00", _v)
				}
				if err != nil {
					_time, err = time.Parse(dateTimeFormat, _v)
				}
			case InputDatetimeLocal:
				parser(
					dateTimeLocalFormat+".05",
					dateTimeLocalFormat+":05",
					dateTimeLocalFormat,
				)
			case InputTime:
				parser(
					timeFormat+".05",
					timeFormat+":05",
					timeFormat,
				)
			case InputDate:
				_time, err = time.ParseInLocation(dateFormat, _v, f.loc)
			case InputMonth:
				_time, err = time.ParseInLocation(monthFormat, _v, f.loc)
			case InputWeek:
				_vv := strings.Split(_v, "-W")
				if len(_vv) < 2 {
					err = fmt.Errorf(f.T("ErrOutOfBound"))
					continue
				}
				var year int64
				year, err = strconv.ParseInt(_vv[0], 10, 64)
				if err != nil {
					continue
				}
				var week int64
				week, err = strconv.ParseInt(_vv[1], 10, 64)
				if err != nil {
					continue
				}
				_time = StartingDayOfWeek(int(year), int(week), f.loc)
			}

			if err != nil {
				continue
			}

		blank:

			*ptr = _time

		case **multipart.FileHeader:
			fileHeader := f.Value.FileShift(preferedName)
			if fileHeader == nil {
				continue
			}
			*ptr = fileHeader

		default:
			err = fmt.Errorf(`form: '%T' is not a supported data type for validation,
only *string, *[]string, *int64, *[]int64, *float64, *[]float64, *bool, **multipart.FileHeader
and *time.Time`, ptr)
		}

		fieldC.err = err
	}

	jsonUpdate := func(preferedName string, err error, warning string) {
		m := map[string]interface{}{}
		m["valid"] = err == nil
		if m["valid"].(bool) {
			m["error"] = ""
		} else {
			m["error"] = err.Error()
		}
		m["warning"] = warning
		m["name"] = preferedName
		m["count"] = f.vcount
		f.vcount++
		f.JsonData = append(f.JsonData, m)
	}

	hasError := false

	// Now for full on validation.
	for _, item := range fieldM {
		ptr := item.ptr
		name := item.name
		preferedName := item.preferedName
		fieldFns := item.fieldFns
		_type := item._type
		err := item.err

		warning := ""

		va := validateValue{f, &err, &warning, name, preferedName, fieldFns, _type,
			fmt.Sprintf("%T", ptr)}

		if err != nil {
			hasError = true
			data.addError(name, err)
			data.addWarning(name, warning)
			jsonUpdate(preferedName, err, warning)
			// No point going further with validation, has the field already failed!
			continue
		}

		switch value := ptr.(type) {
		case *string:
			va.str(value)
		case *[]string:
			va.strs(value)
		case *int64:
			va.wnum(value)
		case *[]int64:
			va.wnums(value)
		case *float64:
			va.fnum(value)
		case *[]float64:
			va.fnums(value)
		case *bool:
			va.b(value)
		case *time.Time:
			va.time(value)
		case **multipart.FileHeader:
			va.file(value)
		}

		if err == nil {
			va.fieldsFns.Call("ext", map[string]interface{}{
				"error":   &err,
				"warning": &warning,
			})
		}

		if err != nil {
			hasError = true
		}

		data.addError(name, err)
		data.addWarning(name, warning)

		jsonUpdate(preferedName, err, warning)

	}

	return !hasError, nil
}

func (va validateValue) typeError() {
	*(va.err) = fmt.Errorf(va.form.T("ErrType", map[string]interface{}{
		"DataType": va.t,
	}))
}

func (va validateValue) str(value *string) {
	switch va._type {
	case InputText, InputPassword, InputSearch, InputHidden, InputUrl, InputTel:
		va.strInputText(*value)
	case InputEmail:
		va.strInputEmail(*value)
	case InputRadio:
		va.strInputRadio(*value)
	case InputColor:
		va.strInputColor(*value)
	case Textarea:
		va.strTextarea(*value)
	case Select:
		va.strSelect(*value)
	default:
		va.typeError()
	}
}

func (va validateValue) strs(values *[]string) {
	switch va._type {
	case Select:
		va.strsSelect(*values)
	default:
		va.typeError()
	}
}

func (va validateValue) wnum(value *int64) {
	switch va._type {
	case InputNumber, InputRange, InputHidden:
		va.wnumInputNumber(*value)
	case InputRadio:
		va.wnumInputRadio(*value)
	case Select:
		va.wnumSelect(*value)
	default:
		va.typeError()
	}
}

func (va validateValue) wnums(values *[]int64) {
	switch va._type {
	case Select:
		va.wnumsSelect(*values)
	default:
		va.typeError()
	}
}

func (va validateValue) fnum(value *float64) {
	switch va._type {
	case InputNumber, InputRange, InputHidden:
		va.fnumInputNumber(*value)
	case InputRadio:
		va.fnumInputRadio(*value)
	case Select:
		va.fnumSelect(*value)
	default:
		va.typeError()
	}
}

func (va validateValue) fnums(values *[]float64) {
	switch va._type {
	case Select:
		va.fnumsSelect(*values)
	default:
		va.typeError()
	}
}

func (va validateValue) b(value *bool) {
	switch va._type {
	case InputCheckbox, InputHidden:
		va.bInputCheckbox(*value)
	default:
		va.typeError()
	}
}

func (va validateValue) time(value *time.Time) {
	switch va._type {
	case InputDatetime, InputDatetimeLocal, InputTime, InputDate, InputMonth, InputWeek:
		va.timeInputTime(*value)
	default:
		va.typeError()
	}
}

func (va validateValue) file(value **multipart.FileHeader) {
	switch va._type {
	case InputFile:
		va.fileInputFile(*value)
	default:
		va.typeError()
	}
}
