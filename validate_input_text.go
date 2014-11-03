package form

import (
	"fmt"
	"regexp"
)

func (va validateValue) strInputText(value string) {
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

	// MustMatch

	mustMatchFieldName := ""
	mustMatchFieldValue := ""
	mustMatchErr := ""

	va.fieldsFns.Call("mustmatch", map[string]interface{}{
		"name":  &mustMatchFieldName,
		"value": &mustMatchFieldValue,
		"err":   &mustMatchErr,
	})

	if mustMatchFieldName == "" {
		goto skipmatch
	}

	if value != mustMatchFieldValue {
		if mustMatchErr == "" {
			mustMatchErr = va.form.T("ErrMustMatchMismatch", map[string]interface{}{
				"Name": mustMatchFieldName,
			})
		}
		va.data.Errors[va.name] = fmt.Errorf(mustMatchErr)
		return
	}

skipmatch:

	// Size

	min, max := int(-1), int(-1)
	minErr, maxErr := "", ""

	va.fieldsFns.Call("size", map[string]interface{}{
		"min":    &min,
		"max":    &max,
		"minErr": &minErr,
		"maxErr": &maxErr,
	})

	if min <= -1 && max <= -1 {
		goto skipmax
	} else if min <= -1 {
		goto skipmin
	}

	// Min Size

	if len(value) < min {
		if minErr == "" {
			minErr = va.form.T("ErrMinChar", map[string]interface{}{
				"Count": min,
			})
		}
		va.data.Errors[va.name] = fmt.Errorf(minErr)
		return
	}

skipmin:

	if max <= -1 {
		goto skipmax
	}

	if len(value) > max {
		if maxErr == "" {
			maxErr = va.form.T("ErrMaxChar", map[string]interface{}{
				"Count": max,
			})
		}
		va.data.Errors[va.name] = fmt.Errorf(maxErr)
		return
	}

skipmax:

	// Check Pattern

	var pattern *regexp.Regexp
	patternErr := ""

	va.fieldsFns.Call("pattern", map[string]interface{}{
		"pattern": &pattern,
		"err":     &patternErr,
	})

	if pattern == nil {
		goto skiprule
	}

	if !pattern.MatchString(value) {
		if patternErr == "" {
			patternErr = va.form.T("ErrPatternMismatch", map[string]interface{}{
				"Pattern": pattern,
			})
		}
		va.data.Errors[va.name] = fmt.Errorf(patternErr)
		return
	}

skiprule:
}
