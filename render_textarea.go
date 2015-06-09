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
		delete(attr, "maxlength")
		delete(attr, "required")
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

	_s := ""
	var size *Size

	r.fieldsFns.Call("size", map[string]interface{}{
		"size": &size,
	})

	if size != nil && size.Max > 0 {
		textarea.Attr["maxlength"] = fmt.Sprintf("%d", size.Max)
	}

	mandatory := false

	r.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &_s,
	})

	if mandatory {
		textarea.Attr["required"] = " "
	}
}
