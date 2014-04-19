package form

import (
	"fmt"
)

func (va validate) strTextarea() {
	value := va.value.String()
	manErr, ok := va.getStr("MandatoryErr")
	if !ok {
		manErr = va.i18n.Key(ErrMandatory)
	}
	if len(value) == 0 {
		minChar, ok := va.getInt("MinChar")
		if !ok {
			minChar, ok = va.getInt("MinLength")
		}
		if ok {
			if minChar >= 1 {
				va.setErr(FormError(manErr))
				return
			}
		}
	}

	minChar, ok := va.getInt("MinChar")
	if !ok {
		minChar, ok = va.getInt("MinLength")
		if !ok {
			mandatory, ok := va.getBool("Mandatory")
			if ok && mandatory {
				minChar = 1
			} else {
				goto skipmin
			}
		}
	}
	if minChar <= 0 {
		goto skipmin
	}

	if int64(len(value)) < minChar {
		if value == "" {
			va.setErr(FormError(manErr))
			return
		}

		if minChar == 1 || value[0] == ' ' || value[0] == '\r' || value[0] == '\n' {
			va.setErr(FormError(manErr))
			return
		} else {
			minCharErr, ok := va.getStr("MinCharErr")
			if ok {
				va.setErr(FormError(minCharErr))
			} else {
				minCharErr, ok = va.getStr("MinLengthErr")
				if ok {
					va.setErr(FormError(minCharErr))
				} else {
					va.setErr(FormError(fmt.Sprintf(va.i18n.Key(ErrMinChar), minChar)))
				}
			}
			return
		}
	}

skipmin:

	maxChar, ok := va.getInt("MaxChar")
	if !ok {
		maxChar, ok = va.getInt("MaxLength")
		if !ok {
			goto skipmax
		}
	}
	if maxChar <= 0 {
		goto skipmax
	}

	if int64(len(value)) > maxChar {
		maxCharErr, ok := va.getStr("MaxCharErr")
		if ok {
			va.setErr(FormError(maxCharErr))
		} else {
			maxCharErr, ok = va.getStr("MaxLengthErr")
			if ok {
				va.setErr(FormError(maxCharErr))
			} else {
				va.setErr(FormError(fmt.Sprintf(va.i18n.Key(ErrMaxChar), maxChar)))
			}
		}
		return
	}

skipmax:

	va.callExt()
}
