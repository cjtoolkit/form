package form

import (
	"fmt"
	"strings"
)

func (va validateValue) strSelect(value string) {
	value = strings.TrimSpace(value)

	var options []Option

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		if value == option.Value {
			return
		}
	}

	if value != "" {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrOutOfBound"))
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) strsSelect(values []string) {
	var options []Option

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		for _, value := range values {
			value = strings.TrimSpace(value)
			if value == option.Value {
				return
			}
		}
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) wnumSelect(value int64) {
	var options []OptionInt

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		if value == option.Value {
			return
		}
	}

	if value != 0 {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrOutOfBound"))
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) wnumsSelect(values []int64) {
	var options []OptionInt

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		for _, value := range values {
			if value == option.Value {
				return
			}
		}
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) fnumSelect(value float64) {
	var options []OptionFloat

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		if value == option.Value {
			return
		}
	}

	if value != 0 {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrOutOfBound"))
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}

func (va validateValue) fnumsSelect(values []float64) {
	var options []OptionFloat

	va.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		va.data.Errors[va.name] = fmt.Errorf(va.form.T("ErrSelectNotWellFormed"))
		return
	}

	for _, option := range options {
		for _, value := range values {
			if value == option.Value {
				return
			}
		}
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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}
