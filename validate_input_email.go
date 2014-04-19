package form

import (
	"fmt"
	"regexp"
)

// From http://www.w3.org/TR/html5/states-of-the-type-attribute.html#valid-e-mail-address
var email_rule = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func (va validate) strInputEmail() {
	value := va.value.String()

	matchValueStr := ""

	MustMatch, ok := va.getStr("MustMatch")
	if ok {
		matchValue := va.v.FieldByName(MustMatch)
		if !matchValue.IsValid() {
			va.setErr(FormError(va.i18n.Key(ErrMustMatchMissing)))
			return
		}
		matchValueStr = matchValue.String()
	} else {
		goto skipmatch
	}

	if matchValueStr != value {
		errMsg, ok := va.getStr("MustMatchErr")
		if !ok {
			errMsg = fmt.Sprintf(va.i18n.Key(ErrMustMatchMismatch), MustMatch)
		}
		va.setErr(FormError(errMsg))
		return
	}

skipmatch:

	mandatory, ok := va.getBool("Mandatory")
	if ok {
		if !mandatory && value == "" {
			va.callExt()
			return
		}
	}

	if !email_rule.MatchString(value) {
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
			va.setErr(FormError(va.i18n.Key(ErrInvalidEmailAddress)))
		}
		return
	}

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
