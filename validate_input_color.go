package form

import (
	"regexp"
)

var color_rule = regexp.MustCompile(`^#([a-fA-F0-9]{6})$`)

func (va validate) strInputColor() {
	value := va.value.String()

	mandatory, ok := va.getBool("Mandatory")
	if ok {
		if !mandatory && value == "" {
			va.callExt()
			return
		}
	}

	if !color_rule.MatchString(value) {
		if mandatory && value == "" {
			manErr, ok := va.getStr("MandatoryErr")
			if !ok {
				manErr = va.i18n.Key(ErrMandatory)
			}
			va.setErr(FormError(manErr))
			return
		}
		errStr, ok := va.getStr("Err")
		if ok {
			va.setErr(FormError(errStr))
		} else {
			va.setErr(FormError(va.i18n.Key(ErrInvalidColorCode)))
		}
		return
	}

	va.callExt()
}
