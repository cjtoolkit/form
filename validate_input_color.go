package form

import (
	"fmt"
	"regexp"
	"strings"
)

var color_rule = regexp.MustCompile(`^#([a-fA-F0-9]{6})$`)

func (va validateValue) strInputColor(value string) {
	value = strings.TrimSpace(value)

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
		va.data.Errors[va.name] = fmt.Errorf(manErr)
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
		va.data.Errors[va.name] = fmt.Errorf(colorErr)
		return
	}
}
