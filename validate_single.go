package form

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (f *form) validateSingle(formPtr FormPtr, name string, value []string) (err error) {
	if len(value) <= 0 {
		err = fmt.Errorf("form: value cannot be nil or empty")
		return
	}

	if !isStructPtr(formPtr) {
		err = fmt.Errorf("form: '%T' is not a pointer", formPtr)
		return
	}

	fields := &Fields{}

	fields.m = map[string]FieldFuncs{}
	fields.n = map[string]*Field{}
	fields.validating = true
	fields.R = f.req

	formPtr.CJForm(fields)

	field := fields.n[name]
	if field == nil {
		err = fmt.Errorf("form: field '%s' does not exist", name)
		return
	}

	fieldName := field.name
	fieldFns := field.funcs

	preferedName := fieldName

	_type := field.typecode

	if _type <= Invalid || _type >= terminate {
		err = fmt.Errorf("form: Invalid form type")
		return
	}

	validator := &validateValue{}
	validator.form = f

	warning := ""
	validator.err = &err
	validator.warning = &warning

	validator.name = fieldName
	validator.preferedName = preferedName

	validator.fieldsFns = fieldFns
	validator._type = _type
	validator.t = fmt.Sprintf("%T", field.ptr)

	switch field.ptr.(type) {

	case *string:
		_value := strings.TrimSpace(value[0])
		(*validator).str(&_value)

	case *[]string:
		values := []string{}
		for _, _value := range value {
			values = append(values, strings.TrimSpace(_value))
		}
		(*validator).strs(&values)

	case *int64:
		var _value int64
		_value, err = strconv.ParseInt(strings.TrimSpace(value[0]), 10, 64)
		if err != nil {
			return
		}
		(*validator).wnum(&_value)

	case *[]int64:
		values := []int64{}
		for _, _value := range value {
			var v int64
			v, err = strconv.ParseInt(strings.TrimSpace(_value), 10, 64)
			if err != nil {
				return
			}
			values = append(values, v)
		}
		(*validator).wnums(&values)

	case *float64:
		var _value float64
		_value, err = strconv.ParseFloat(strings.TrimSpace(value[0]), 64)
		if err != nil {
			return
		}
		(*validator).fnum(&_value)

	case *[]float64:
		values := []float64{}
		for _, _value := range value {
			var v float64
			v, err = strconv.ParseFloat(strings.TrimSpace(_value), 64)
			if err != nil {
				return
			}
			values = append(values, v)
		}
		(*validator).fnums(&values)

	case *bool:
		_value := false
		if strings.TrimSpace(value[0]) == "1" {
			_value = true
		}
		(*validator).b(&_value)

	case *time.Time:
		_valueStr := strings.TrimSpace(value[0])

		var _value time.Time

		parser := func(formats ...string) {
			for _, format := range formats {
				_value, err = time.ParseInLocation(format, _valueStr, f.loc)
				if err == nil {
					return
				}
			}
		}

		if _valueStr == "" {
			goto blank
		}

		switch _type {
		case InputDatetime:
			_value, err = time.Parse(dateTimeLocalFormat+".05Z07:00", _valueStr)
			if err != nil {
				_value, err = time.Parse(dateTimeLocalFormat+":05Z07:00", _valueStr)
			}
			if err != nil {
				_value, err = time.Parse(dateTimeFormat, _valueStr)
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
			_value, err = time.ParseInLocation(dateFormat, _valueStr, f.loc)
		case InputMonth:
			_value, err = time.ParseInLocation(monthFormat, _valueStr, f.loc)
		case InputWeek:
			_vv := strings.Split(_valueStr, "-W")
			if len(_vv) < 2 {
				err = fmt.Errorf(f.T("ErrOutOfBound"))
				return
			}
			var year int64
			year, err = strconv.ParseInt(_vv[0], 10, 64)
			if err != nil {
				return
			}
			var week int64
			week, err = strconv.ParseInt(_vv[1], 10, 64)
			if err != nil {
				return
			}
			_value = StartingDayOfWeek(int(year), int(week), f.loc)
		}

		if err != nil {
			return
		}

	blank:

		(*validator).time(&_value)

	default:
		err = fmt.Errorf(`form: '%T' is not a supported data type for single validation,
only *string, *[]string, *int64, *[]int64, *float64, *[]float64, *bool and *time.Time`, field.ptr)
	}

	return
}
