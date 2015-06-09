package form

import (
	"fmt"
	"regexp"
)

func (r renderValue) strInputText(value string) {
	_type := func() (str string) {
		switch r._type {
		case InputText:
			str = "text"
		case InputSearch:
			str = "search"
		case InputPassword:
			str = "password"
		case InputHidden:
			str = "hidden"
		case InputUrl:
			str = "url"
		case InputTel:
			str = "tel"
		}
		return
	}

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
	input.Attr["type"] = _type()
	input.Attr["value"] = value

	var pattern *regexp.Regexp
	_s := ""

	r.fieldsFns.Call("pattern", map[string]interface{}{
		"pattern": &pattern,
		"err":     &_s,
	})

	if pattern != nil {
		input.Attr["pattern"] = pattern.String()
	}

	var size *Size

	r.fieldsFns.Call("size", map[string]interface{}{
		"size": &size,
	})

	if size != nil && size.Max > 0 {
		input.Attr["maxlength"] = fmt.Sprintf("%d", size.Max)
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
