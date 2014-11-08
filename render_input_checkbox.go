package form

func (r renderValue) bInputCheckbox(value bool) {
	_type := func() (str string) {
		switch r._type {
		case InputCheckbox:
			str = "checkbox"
		case InputHidden:
			str = "hidden"
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
		delete(attr, "checked")
		input.Attr = attr
	} else {
		input.Attr = map[string]string{}
	}

	input.Attr["name"] = r.preferedName
	input.Attr["type"] = _type()
	input.Attr["value"] = "1"

	if value {
		input.Attr["checked"] = " "
	}
}
