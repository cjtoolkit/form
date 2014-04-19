package form

import (
	"reflect"
)

func (va validate) strInputRadio() {
	value := va.value.String()

	errStr := FormError(va.i18n.Key(ErrRadioNotWellFormed))

	m := va.m.MethodByName(va.name + "Radio")
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

	radios, ok := values[0].Interface().([]Radio)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
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

func (va validate) wnumInputRadio() {
	value := va.value.Int()

	errStr := FormError(va.i18n.Key(ErrRadioNotWellFormed))

	m := va.m.MethodByName(va.name + "Radio")
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

	radios, ok := values[0].Interface().([]RadioInt)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
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

func (va validate) fnumInputRadio() {
	value := va.value.Float()

	errStr := FormError(va.i18n.Key(ErrRadioNotWellFormed))

	m := va.m.MethodByName(va.name + "Radio")
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

	radios, ok := values[0].Interface().([]RadioFloat)
	if !ok {
		va.setErr(errStr)
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
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
