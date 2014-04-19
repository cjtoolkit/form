package form

import (
	"fmt"
)

func (r render) strInputColor() {
	fmt.Fprint(r.w, `<input name="`, es(r.preferedName), `" type="color" `)

	fmt.Fprint(r.w, `value="`, es(r.value.String()), `" `)

	r.attr("type", "value", "pattern", "maxlength")

	fmt.Fprint(r.w, `/>`)
}
