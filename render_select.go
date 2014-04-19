package form

import (
	"fmt"
	"reflect"
)

func (r render) strSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]Option)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	value := r.value.String()

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != "" {
			fmt.Fprint(r.w, `value="`, es(option.Value), `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if value == "" {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}

func (r render) strsSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]Option)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" multiple `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	_values := r.value.Interface().([]string)
	values_len := len(_values)

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != "" {
			fmt.Fprint(r.w, `value="`, es(option.Value), `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if values_len == 0 {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			for _, value := range _values {
				if value == option.Value {
					fmt.Fprint(r.w, `selected `)
				}
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}

func (r render) wnumSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]OptionInt)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	value := r.value.Int()

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != 0 {
			fmt.Fprint(r.w, `value="`, option.Value, `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if value == 0 {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}

func (r render) wnumsSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]OptionInt)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" multiple `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	_values := r.value.Interface().([]int64)
	values_len := len(_values)

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != 0 {
			fmt.Fprint(r.w, `value="`, option.Value, `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if values_len == 0 {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			for _, value := range _values {
				if value == option.Value {
					fmt.Fprint(r.w, `selected `)
				}
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}

func (r render) fnumSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]OptionFloat)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	value := r.value.Float()

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != 0 {
			fmt.Fprint(r.w, `value="`, option.Value, `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if value == 0 {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == option.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}

func (r render) fnumsSelect() {
	m := r.m.MethodByName(r.name + "Options")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	options, ok := values[0].Interface().([]OptionFloat)
	if !ok {
		return
	}

	fmt.Fprint(r.w, `<select name="`, es(r.preferedName), `" multiple `)

	r.attr("name", "multiple")

	fmt.Fprint(r.w, `>`)

	fmt.Fprintln(r.w)

	attrOptionExc := []string{"value", "label", "selected"}

	_values := r.value.Interface().([]float64)
	values_len := len(_values)

	for _, option := range options {
		fmt.Fprint(r.w, `<option `)

		if option.Value != 0 {
			fmt.Fprint(r.w, `value="`, option.Value, `" `)
		}

		if option.Label != "" {
			fmt.Fprint(r.w, `label="`, es(option.Label), `" `)
		}

		if values_len == 0 {
			if option.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			for _, value := range _values {
				if value == option.Value {
					fmt.Fprint(r.w, `selected `)
				}
			}
		}

		if option.Attr == nil {
			goto do_content
		}

		for _, exc := range attrOptionExc {
			delete(option.Attr, exc)
		}

		for key, value := range option.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	do_content:

		if option.Content != "" {
			fmt.Fprint(r.w, `>`, es(option.Content), `</option>`)
		} else {
			fmt.Fprint(r.w, `/>`)
		}

		fmt.Fprintln(r.w)
	}

	fmt.Fprint(r.w, `</select>`)
}
