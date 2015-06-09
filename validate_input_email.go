package form

import (
	"fmt"
	"regexp"
)

// From http://www.w3.org/TR/html5/states-of-the-type-attribute.html#valid-e-mail-address
var email_rule = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func (va validateValue) strInputEmail(value string) {
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

	// MustMatch

	var mustMatch *MustMatch

	va.fieldsFns.Call("mustmatch", map[string]interface{}{
		"mustmatch": &mustMatch,
	})

	if mustMatch == nil || mustMatch.Name == "" {
		goto skipmatch
	}

	if value != *mustMatch.Value {
		if mustMatch.Err == "" {
			mustMatch.Err = va.form.T("ErrMustMatchMismatch", map[string]interface{}{
				"Name": mustMatch.Name,
			})
		}
		*(va.err) = fmt.Errorf(mustMatch.Err)
		return
	}

skipmatch:

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

	emailErr := ""

	va.fieldsFns.Call("email", map[string]interface{}{
		"err": &emailErr,
	})

	truth := email_rule.MatchString(value)

	if !truth {
		if emailErr == "" {
			emailErr = va.form.T("ErrInvalidEmailAddress")
		}
		*(va.err) = fmt.Errorf(emailErr)
		return
	}
}
