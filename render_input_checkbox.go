package form

import (
	"fmt"
)

func (r renderValue) bInputCheckbox(value bool) {
	_type := func() (str string) {
		switch r._type {
		case InputCheckbox:
			str = "checkbox"
		case InputHidden:
			str = "hidden"
		}
		return
	}

	w := r.w

	fmt.Fprintf(w, `<input name="%s" type="%s" value="1" `, es(r.preferedName), _type())

	if value {
		fmt.Fprint(w, `checked`)
	}

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "value")
		delete(attr, "checked")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `/>`)
}
