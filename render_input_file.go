package form

import (
	"fmt"
	"strings"
)

func (r render) fileInputFile() {
	fmt.Fprint(r.w, `<input type="file" name="`, es(r.preferedName), `" `)

	mimes, ok := r.getStrs("Accept")
	if ok {
		fmt.Fprint(r.w, `accept="`)

		fmt.Fprint(r.w, es(strings.Join(mimes, "|")))

		fmt.Fprint(r.w, `" `)
	}

	r.attr("type", "name", "accept")

	fmt.Fprint(r.w, `/>`)
}
