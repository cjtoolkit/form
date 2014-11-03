package form

import (
	"fmt"
)

func (va validateValue) strTextarea(value string) {
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
		*(va.err) = fmt.Errorf(minErr)
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
		*(va.err) = fmt.Errorf(maxErr)
		return
	}

skipmax:
}
