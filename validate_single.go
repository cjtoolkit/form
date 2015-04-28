package form

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (f *form) validateSingle(structPtr Interface, name string, value []string) (err error) {
	if len(value) <= 0 {
		err = fmt.Errorf("form: value cannot be nil or empty")
		return
	}

	var preStructPtr interface{} = structPtr

	if v, ok := structPtr.(Hijacker); ok {
		preStructPtr = v.CJStructPtr()
	}

	t := reflect.TypeOf(preStructPtr)
	vc := reflect.ValueOf(preStructPtr)

	switch {
	case isStructPtr(t):
		t = t.Elem()
		vc = vc.Elem()
	default:
		err = fmt.Errorf("form: '%p' is not a struct pointer", preStructPtr)
		return
	}

	fields := Fields{
		map[string]FieldFuncs{},
		map[string]*Field{},
		map[string]*Field{},
		nil,
	}

	structPtr.CJForm(&fields)

	afield := fields.n[name]
	if afield == nil {
		afield = fields.nm[name]
	}
	if afield == nil {
		err = fmt.Errorf("form: field '%s' does not exist", name)
		return
	}

	fieldName := afield.name
	fieldFns := afield.funcs

	_, exist := t.FieldByName(fieldName)
	if !exist {
		err = fmt.Errorf("form: '%s' field does not exist", fieldName)
		return
	}

	preferedName := name

	field := vc.FieldByName(fieldName)
	if !field.CanSet() {
		err = fmt.Errorf("form: '%s' field cannot be set", fieldName)
		return
	}

	_type := Invalid

	fieldFns.Call("init", map[string]interface{}{
		"type": &_type,
	})

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
	validator.t = t

	switch field.Interface().(type) {

	case string:
		_value := strings.TrimSpace(value[0])
		(*validator).str(_value)

	case []string:
		values := []string{}
		for _, _value := range value {
			values = append(values, strings.TrimSpace(_value))
		}
		(*validator).strs(values)

	case int64:
		var _value int64
		_value, err = strconv.ParseInt(strings.TrimSpace(value[0]), 10, 64)
		if err != nil {
			return
		}
		(*validator).wnum(_value)

	case []int64:
		values := []int64{}
		for _, _value := range value {
			var v int64
			v, err = strconv.ParseInt(strings.TrimSpace(_value), 10, 64)
			if err != nil {
				return
			}
			values = append(values, v)
		}
		(*validator).wnums(values)

	case float64:
		var _value float64
		_value, err = strconv.ParseFloat(strings.TrimSpace(value[0]), 64)
		if err != nil {
			return
		}
		(*validator).fnum(_value)

	case []float64:
		values := []float64{}
		for _, _value := range value {
			var v float64
			v, err = strconv.ParseFloat(strings.TrimSpace(_value), 64)
			if err != nil {
				return
			}
			values = append(values, v)
		}
		(*validator).fnums(values)

	case bool:
		_value := false
		if strings.TrimSpace(value[0]) == "1" {
			_value = true
		}
		(*validator).b(_value)

	case time.Time:
		_valueStr := strings.TrimSpace(value[0])

		var _value time.Time

		if _valueStr == "" {
			goto blank
		}

		switch _type {
		case InputDatetime:
			_value, err = time.Parse(dateTimeFormat, _valueStr)
		case InputDatetimeLocal:
			_value, err = time.Parse(dateTimeLocalFormat, _valueStr)
		case InputTime:
			_value, err = time.Parse(timeFormat, _valueStr)
		case InputDate:
			_value, err = time.Parse(dateFormat, _valueStr)
		case InputMonth:
			_value, err = time.Parse(monthFormat, _valueStr)
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
			_value = StartingDayOfWeek(int(year), int(week))
		}

		if err != nil {
			return
		}

	blank:

		(*validator).time(_value)

	default:
		err = fmt.Errorf(`form: '%v' is not a supported data type for single validation,
only string, []string, int64, []int64, float64, []float64, bool and time.Time`, t)
	}

	return
}
