package form

import (
	"fmt"
	"strings"
)

func (r renderValue) strInputRadio(value string) {
	value = strings.TrimSpace(value)
	w := r.w

	var radios []Radio

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(w, `<label>`)
		}

		fmt.Fprintf(w, `<input name="%s" type="radio" value="%s" `, es(r.preferedName), es(radio.Value))

		if value == "" {
			if radio.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			fmt.Fprint(w, ParseAttr(radio.Attr))
		}

		fmt.Fprint(w, `/>`)

		if radio.Label != "" {
			fmt.Fprint(w, ` %s</label>`, es(radio.Label))
		}
	}
}

func (r renderValue) wnumInputRadio(value int64) {
	w := r.w

	var radios []RadioInt

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(w, `<label>`)
		}

		fmt.Fprintf(w, `<input name="%s" type="radio" value="%d" `, es(r.preferedName), radio.Value)

		if value == 0 {
			if radio.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			fmt.Fprint(w, ParseAttr(radio.Attr))
		}

		fmt.Fprint(w, `/>`)

		if radio.Label != "" {
			fmt.Fprint(w, ` %s</label>`, es(radio.Label))
		}
	}
}

func (r renderValue) fnumInputRadio(value float64) {
	w := r.w

	var radios []RadioFloat

	r.fieldsFns.Call("radio", map[string]interface{}{
		"radio": &radios,
	})

	if radios == nil {
		panic(fmt.Errorf(r.form.T("ErrRadioNotWellFormed")))
		return
	}

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(w, `<label>`)
		}

		fmt.Fprintf(w, `<input name="%s" type="radio" value="%f" `, es(r.preferedName), radio.Value)

		if value == 0 {
			if radio.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if radio.Attr != nil {
			delete(radio.Attr, "type")
			delete(radio.Attr, "name")
			delete(radio.Attr, "selected")
			delete(radio.Attr, "value")
			fmt.Fprint(w, ParseAttr(radio.Attr))
		}

		fmt.Fprint(w, `/>`)

		if radio.Label != "" {
			fmt.Fprint(w, ` %s</label>`, es(radio.Label))
		}
	}
}
