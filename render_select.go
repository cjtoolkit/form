package form

import (
	"fmt"
	"strings"
)

func (r renderValue) strSelect(value string) {
	value = strings.TrimSpace(value)
	w := r.w

	var options []Option

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != "" {
			fmt.Fprintf(w, `value="%s" `, es(option.Value))
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if value == "" {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}

func (r renderValue) strsSelect(values []string) {
	w := r.w

	var options []Option

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" multiple `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != "" {
			fmt.Fprintf(w, `value="%s" `, es(option.Value))
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if len(values) == 0 {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			for _, value := range values {
				value = strings.TrimSpace(value)
				if value == option.Value {
					fmt.Fprint(w, `selected `)
					break
				}
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}

func (r renderValue) wnumSelect(value int64) {
	w := r.w

	var options []OptionInt

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != 0 {
			fmt.Fprintf(w, `value="%d" `, option.Value)
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if value == 0 {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}

func (r renderValue) wnumsSelect(values []int64) {
	w := r.w

	var options []OptionInt

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" multiple `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != 0 {
			fmt.Fprintf(w, `value="%d" `, option.Value)
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if len(values) == 0 {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			for _, value := range values {
				if value == option.Value {
					fmt.Fprint(w, `selected `)
					break
				}
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}

func (r renderValue) fnumSelect(value float64) {
	w := r.w

	var options []OptionFloat

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != 0 {
			fmt.Fprintf(w, `value="%f" `, option.Value)
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if value == 0 {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(w, `selected `)
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}

func (r renderValue) fnumsSelect(values []float64) {
	w := r.w

	var options []OptionFloat

	r.fieldsFns.Call("option", map[string]interface{}{
		"option": &options,
	})

	if options == nil {
		panic(fmt.Errorf(r.form.T("ErrSelectNotWellFormed")))
		return
	}

	fmt.Fprintf(w, `<select name="%s" multiple `, es(r.preferedName))

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "multiple")
		fmt.Fprint(w, ParseAttr(attr))
	}

	fmt.Fprint(w, `>`)

	for _, option := range options {
		fmt.Fprint(w, `<option `)

		if option.Value != 0 {
			fmt.Fprintf(w, `value="%f" `, option.Value)
		}

		if option.Label != "" {
			fmt.Fprintf(w, `label="%s" `, es(option.Label))
		}

		if len(values) == 0 {
			if option.Selected {
				fmt.Fprint(w, `selected `)
			}
		} else {
			for _, value := range values {
				if value == option.Value {
					fmt.Fprint(w, `selected `)
					break
				}
			}
		}

		if option.Attr != nil {
			delete(option.Attr, "value")
			delete(option.Attr, "label")
			delete(option.Attr, "selected")
			fmt.Fprint(w, ParseAttr(option.Attr))
		}

		if option.Content != "" {
			fmt.Fprintf(w, `>%s</option>`, es(option.Content))
		} else {
			fmt.Fprint(w, `/>`)
		}
	}

	fmt.Fprint(w, `</select>`)
}
