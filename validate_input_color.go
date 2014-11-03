package form

import (
	"fmt"
	"regexp"
)

var color_rule = regexp.MustCompile(`^#([a-fA-F0-9]{6})$`)

func (va validateValue) strInputColor(value string) {
	// Mandatory

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory && len(value) == 0 {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}

	// Colour Check
	colorErr := ""

	va.fieldsFns.Call("color", map[string]interface{}{
		"err": &colorErr,
	})

	if !color_rule.MatchString(value) {
		if colorErr == "" {
			colorErr = va.form.T("ErrInvalidColorCode")
		}
		*(va.err) = fmt.Errorf(colorErr)
		return
	}
}
