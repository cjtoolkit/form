package form

import (
	"fmt"
)

func (r renderValue) strInputEmail(value string) {
	w := r.w

	fmt.Fprintf(w, `<input name="%s" type="email" value="%s" `, es(r.preferedName), es(value))

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
