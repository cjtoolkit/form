package form

import (
	"strings"
)

func (r renderValue) fileInputFile() {
	input := &FirstLayerInput{}
	r.fls.append(input)

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "accept")
		delete(attr, "required")
		input.Attr = attr
	} else {
		input.Attr = map[string]string{}
	}

	input.Attr["name"] = r.preferedName
	input.Attr["type"] = "file"

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
		input.Attr["accept"] = strings.Join(mimes, "|")
	}

	mandatory := false

	r.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &_s,
	})

	if mandatory {
		input.Attr["required"] = " "
	}
}
