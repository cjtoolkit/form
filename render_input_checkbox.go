package form

import (
	"fmt"
)

func (r render) bInputCheckbox(hidden bool) {
	_type := "checkbox"
	if hidden {
		_type = "hidden"
	}
	fmt.Fprint(r.w, `<input type="`, _type, `" name="`, es(r.preferedName), `" value="1" `)

	if r.value.Bool() {
		fmt.Fprint(r.w, `checked `)
	}

	r.attr("type", "name", "value", "checked")

	fmt.Fprint(r.w, `/>`)
}
