package form

import (
	"fmt"
	"strings"
)

func (r renderValue) fileInputFile() {
	w := r.w
	fmt.Fprintf(w, `<input name="%s" type="file" `, es(r.preferedName))

	// Size and Mime

	_n := int64(-1)
	_s := ""
	var mimes []string

	r.fieldsFns.Call("file", map[string]interface{}{
		"size":      &_n,
		"sizeErr":   &_s,
		"accept":    &mimes,
		"acceptErr": &_s,
	})

	if mimes != nil {
		fmt.Fprintf(w, `accept="%s" `, es(strings.Join(mimes, "|")))
	}

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "accept")
		fmt.Fprint(w, RenderAttr(attr))
	}

	fmt.Fprint(w, `/>`)
}
