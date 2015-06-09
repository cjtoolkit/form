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

	var size *Size

	va.fieldsFns.Call("size", map[string]interface{}{
		"size": &size,
	})

	if size == nil || (size.Min <= -1 && size.Max <= -1) {
		goto skipmax
	} else if size.Min <= -1 {
		goto skipmin
	}

	// Min Size

	if len(value) < size.Min {
		if size.MinErr == "" {
			size.MinErr = va.form.T("ErrMinChar", map[string]interface{}{
				"Count": size.Min,
			})
		}
		*(va.err) = fmt.Errorf(size.MinErr)
		return
	}

skipmin:

	if size.Max <= -1 {
		goto skipmax
	}

	if len(value) > size.Max {
		if size.MaxErr == "" {
			size.MaxErr = va.form.T("ErrMaxChar", map[string]interface{}{
				"Count": size.Max,
			})
		}
		*(va.err) = fmt.Errorf(size.MaxErr)
		return
	}

skipmax:
}
