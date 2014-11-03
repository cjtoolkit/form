package form

import (
	"fmt"
)

func (va validateValue) strInputRadio(value string) {
	var radios []Radio

	va.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		*(va.err) = fmt.Errorf(va.form.T("ErrRadioNotWellFormed"))
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
			return
		}
	}

	if value != "" {
		*(va.err) = fmt.Errorf(va.form.T("ErrOutOfBound"))
		return
	}

	// Mandatory

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) wnumInputRadio(value int64) {
	var radios []RadioInt

	va.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		*(va.err) = fmt.Errorf(va.form.T("ErrRadioNotWellFormed"))
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
			return
		}
	}

	if value != 0 {
		*(va.err) = fmt.Errorf(va.form.T("ErrOutOfBound"))
		return
	}

	// Mandatory

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) fnumInputRadio(value float64) {
	var radios []RadioFloat

	va.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		*(va.err) = fmt.Errorf(va.form.T("ErrRadioNotWellFormed"))
		return
	}

	for _, radio := range radios {
		if value == radio.Value {
			return
		}
	}

	if value != 0 {
		*(va.err) = fmt.Errorf(va.form.T("ErrOutOfBound"))
		return
	}

	// Mandatory

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}
}
