package form

import (
	"fmt"
)

func (va validateValue) bInputCheckbox(value bool) {
	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory && !value {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		va.data.Errors[va.name] = fmt.Errorf(manErr)
		return
	}
}
