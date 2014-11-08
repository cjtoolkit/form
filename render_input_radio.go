package form

import (
	"fmt"
)

func (r renderValue) strInputRadio(value string) {
	var radios []Radio

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		input := &FirstLayerInput{}
		r.fls.append(input)

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			input.Attr = radio.Attr
		} else {
			input.Attr = map[string]string{}
		}

		input.Attr["name"] = r.preferedName
		input.Attr["type"] = "radio"
		input.Attr["value"] = radio.Value

		if value == "" {
			if radio.Selected {
				input.Attr["selected"] = " "
			}
		} else {
			if value == radio.Value {
				input.Attr["selected"] = " "
			}
		}

		if radio.Label != "" {
			input.Label = radio.Label
		}
	}
}

func (r renderValue) wnumInputRadio(value int64) {
	var radios []RadioInt

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		input := &FirstLayerInput{}
		r.fls.append(input)

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			input.Attr = radio.Attr
		} else {
			input.Attr = map[string]string{}
		}

		input.Attr["name"] = r.preferedName
		input.Attr["type"] = "radio"
		input.Attr["value"] = fmt.Sprintf("%d", radio.Value)

		if value == 0 {
			if radio.Selected {
				input.Attr["selected"] = " "
			}
		} else {
			if value == radio.Value {
				input.Attr["selected"] = " "
			}
		}

		if radio.Label != "" {
			input.Label = radio.Label
		}
	}
}

func (r renderValue) fnumInputRadio(value float64) {
	var radios []RadioFloat

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		input := &FirstLayerInput{}
		r.fls.append(input)

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			input.Attr = radio.Attr
		} else {
			input.Attr = map[string]string{}
		}

		input.Attr["name"] = r.preferedName
		input.Attr["type"] = "radio"
		input.Attr["value"] = fmt.Sprintf("%f", radio.Value)

		if value == 0 {
			if radio.Selected {
				input.Attr["selected"] = " "
			}
		} else {
			if value == radio.Value {
				input.Attr["selected"] = " "
			}
		}

		if radio.Label != "" {
			input.Label = radio.Label
		}
	}
}
