package form

import (
	"fmt"
	"regexp"
)

func (r renderValue) strInputText(value string) {
	_type := func() (str string) {
		switch r._type {
		case InputText:
			str = "text"
		case InputSearch:
			str = "search"
		case InputPassword:
			str = "password"
		case InputHidden:
			str = "hidden"
		case InputUrl:
			str = "url"
		case InputTel:
			str = "tel"
		}
		return
	}

	w := r.w

	fmt.Fprintf(w, `<input name="%s" type="%s" value="%s" `, es(r.preferedName), _type(), es(value))

	var pattern *regexp.Regexp
	_s := ""

	r.fieldsFns.Call("pattern", map[string]interface{}{
		"pattern": &pattern,
		"err":     &_s,
	})

	if pattern != nil {
		fmt.Fprintf(w, `pattern="%s" `, es(pattern.String()))
	}

	_n, max := int(-1), int(-1)

	r.fieldsFns.Call("size", map[string]interface{}{
		"min":    &_n,
		"max":    &max,
		"minErr": &_s,
		"maxErr": &_s,
	})

	if max > 0 {
		fmt.Fprintf(w, `maxlength="%d" `, max)
	}

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "value")
		delete(attr, "pattern")
		delete(attr, "mexlength")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `/>`)
}
