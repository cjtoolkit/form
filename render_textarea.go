package form

import (
	"fmt"
)

func (r render) strTextarea() {
	fmt.Fprint(r.w, `<textarea name="`+es(r.preferedName), `" `)

	rows, ok := r.getInt("Rows")
	if ok {
		fmt.Fprint(r.w, `rows="`, rows, `" `)
	} else {
		fmt.Fprint(r.w, `rows="`, 4, `" `)
	}

	cols, ok := r.getInt("Cols")
	if ok {
		fmt.Fprint(r.w, `cols="`, cols, `" `)
	} else {
		fmt.Fprint(r.w, `cols="`, 25, `" `)
	}

	r.attr("rows", "cols")

	fmt.Fprint(r.w, `>`)

	fmt.Fprint(r.w, es(r.value.String()))

	fmt.Fprint(r.w, `</textarea>`)
}
