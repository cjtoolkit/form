package form

import (
	"fmt"
)

func (r renderValue) strTextarea(value string) {
	textarea := &FirstLayerTextarea{}
	r.fls.append(textarea)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "rows")
		delete(attr, "cols")
		textarea.Attr = attr
	} else {
		textarea.Attr = map[string]string{}
	}

	textarea.Attr["name"] = r.preferedName
	textarea.Content = value

	rows := int(4)
	cols := int(25)

	r.fieldsFns.Call("textarea", map[string]interface{}{
		"rows": &rows,
		"cols": &cols,
	})

	textarea.Attr["rows"] = fmt.Sprintf("%d", rows)
	textarea.Attr["cols"] = fmt.Sprintf("%d", cols)
}
