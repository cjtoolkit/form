package form

import (
	"fmt"
)

func (r render) strInputEmail() {
	fmt.Fprint(r.w, `<input type="email" name="`, es(r.preferedName), `" `)

	fmt.Fprint(r.w, `value="`, es(r.value.String()), `" `)

	r.attr("type", "value", "pattern", "maxlength")

	fmt.Fprint(r.w, `/>`)
}
