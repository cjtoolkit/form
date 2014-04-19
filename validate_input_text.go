package form

import (
	"fmt"
	"regexp"
	"strings"
)

func (va validate) strInputText() {
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

	value = strings.TrimSpace(value)

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

	var truth bool
	var err error
	var regExpStr string

	regExp, ok := va.getRegExp("RegExp")
	if !ok {
		regExp, ok = va.getRegExp("Pattern")
	}
	if ok {
		regExpStr = regExp.String()
		goto rule_check
	}

	regExpStr, ok = va.getStr("RegExp")
	if !ok {
		regExpStr, ok = va.getStr("Pattern")
		if !ok {
			goto skiprule
		}
	}

	regExp, err = regexp.Compile(regExpStr)
	if err != nil {
		va.setErr(err)
		return
	}

rule_check:

	truth = regExp.MatchString(value)

	if !truth {
		regExpErr, ok := va.getStr("RegExpErr")
		if !ok {
			regExpErr, ok = va.getStr("PatternErr")
			if !ok {
				regExpErr = fmt.Sprintf(va.i18n.Key(ErrPatternMismatch), regExpStr)
			}
		}
		va.setErr(FormError(regExpErr))
		return
	}

skiprule:

	va.callExt()
}
