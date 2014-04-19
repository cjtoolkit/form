package form

import (
	"reflect"
)

func (va validate) strSelect() {
	value := va.value.String()

	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]Option)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		if value == option.Value {
			va.callExt()
			return
		}
	}

	if value != "" {
		va.setErr(FormError(va.i18n.Key(ErrOutOfBound)))
		return
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}

func (va validate) strsSelect() {
	_values := va.value.Interface().([]string)

	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]Option)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		selected := false
		for _, value := range _values {
			if value == option.Value {
				selected = true
			}
		}
		if selected {
			va.callExt()
			return
		}
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}

func (va validate) wnumSelect() {
	value := va.value.Int()

	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]OptionInt)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		if value == option.Value {
			va.callExt()
			return
		}
	}

	if value != 0 {
		va.setErr(FormError(va.i18n.Key(ErrOutOfBound)))
		return
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}

func (va validate) wnumsSelect() {
	_values := va.value.Interface().([]int64)
	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]OptionInt)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		selected := false
		for _, value := range _values {
			if value == option.Value {
				selected = true
			}
		}
		if selected {
			va.callExt()
			return
		}
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}

func (va validate) fnumSelect() {
	value := va.value.Float()

	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]OptionFloat)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		if value == option.Value {
			va.callExt()
			return
		}
	}

	if value != 0 {
		va.setErr(FormError(va.i18n.Key(ErrOutOfBound)))
		return
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}

func (va validate) fnumsSelect() {
	_values := va.value.Interface().([]float64)

	errStr := FormError(va.i18n.Key(ErrSelectNotWellFormed))

	m := va.m.MethodByName(va.name + "Options")
	if !m.IsValid() {
		va.setErr(errStr)
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		va.setErr(errStr)
		return
	}

	options, ok := values[0].Interface().([]OptionFloat)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, option := range options {
		selected := false
		for _, value := range _values {
			if value == option.Value {
				selected = true
			}
		}
		if selected {
			va.callExt()
			return
		}
	}

	b, ok := va.getBool("Mandatory")
	if ok {
		if b {
			manErr, ok := va.getStr("MandatoryErr")
			if ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrNotSelect)))
			}
			return
		}
	}

	va.callExt()
}
