package form

import (
	"fmt"
)

func (r renderValue) strInputEmail(value string) {
	input := &FirstLayerInput{}
	r.fls.append(input)

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
		delete(attr, "required")
		input.Attr = attr
	} else {
		input.Attr = map[string]string{}
	}

	input.Attr["name"] = r.preferedName
	input.Attr["type"] = "email"
	input.Attr["value"] = value

	var size *Size

	r.fieldsFns.Call("size", map[string]interface{}{
		"size": &size,
	})

	if size != nil && size.Max > 0 {
		input.Attr["maxlength"] = fmt.Sprintf("%d", size.Max)
	}

	_s := ""
	mandatory := false

	r.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &_s,
	})

	if mandatory {
		input.Attr["required"] = " "
	}
}
