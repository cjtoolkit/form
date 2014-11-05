package form

import (
	"fmt"
)

func (r renderValue) strTextarea(value string) {
	w := r.w

	fmt.Fprintf(w, `<textarea name="%s" `, es(r.preferedName))

	rows := int(4)
	cols := int(25)

	r.fieldsFns.Call("textarea", map[string]interface{}{
		"rows": &rows,
		"cols": &cols,
	})

	fmt.Fprintf(w, `rows="%d" cols="%d" `, rows, cols)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "rows")
		delete(attr, "cols")
		fmt.Fprint(w, RenderAttr(attr))
	}

	fmt.Fprintf(w, `>%s</textarea>`, es(value))
}
