package form

import (
	"fmt"
	"reflect"
)

func (r render) strInputRadio() {
	m := r.m.MethodByName(r.name + "Radio")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	radios, ok := values[0].Interface().([]Radio)
	if !ok {
		return
	}

	attrExclude := []string{"type", "name", "selected", "value"}

	value := r.value.String()

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(r.w, `<label>`, radio.Label, `: `)
		}

		fmt.Fprint(r.w, `<input type="radio" name"`, es(r.preferedName), `" `)

		fmt.Fprint(r.w, `value="`, es(radio.Value), `" `)

		if value == "" {
			if radio.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if radio.Attr == nil {
			goto end
		}

		for _, exc := range attrExclude {
			delete(radio.Attr, exc)
		}

		for key, value := range radio.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	end:

		fmt.Fprint(r.w, `/>`)
		if radio.Label != "" {
			fmt.Fprint(r.w, `</label>`)
		}
		fmt.Fprintln(r.w)
	}
}

func (r render) wnumInputRadio() {
	m := r.m.MethodByName(r.name + "Radio")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	radios, ok := values[0].Interface().([]RadioInt)
	if !ok {
		return
	}

	attrExclude := []string{"type", "name", "selected", "value"}

	value := r.value.Int()

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(r.w, `<label>`, radio.Label, `: `)
		}

		fmt.Fprint(r.w, `<input type="radio" name"`, es(r.preferedName), `" `)

		fmt.Fprint(r.w, `value="`, radio.Value, `" `)

		if value == 0 {
			if radio.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if radio.Attr == nil {
			goto end
		}

		for _, exc := range attrExclude {
			delete(radio.Attr, exc)
		}

		for key, value := range radio.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	end:

		fmt.Fprint(r.w, `/>`)
		if radio.Label != "" {
			fmt.Fprint(r.w, `</label>`)
		}
		fmt.Fprintln(r.w)
	}
}

func (r render) fnumInputRadio() {
	m := r.m.MethodByName(r.name + "Radio")
	if !m.IsValid() {
		return
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return
	}

	radios, ok := values[0].Interface().([]RadioFloat)
	if !ok {
		return
	}

	attrExclude := []string{"type", "name", "selected", "value"}

	value := r.value.Float()

	for _, radio := range radios {
		if radio.Label != "" {
			fmt.Fprint(r.w, `<label>`, radio.Label, `: `)
		}

		fmt.Fprint(r.w, `<input type="radio" name"`, es(r.preferedName), `" `)

		fmt.Fprint(r.w, `value="`, radio.Value, `" `)

		if value == 0 {
			if radio.Selected {
				fmt.Fprint(r.w, `selected `)
			}
		} else {
			if value == radio.Value {
				fmt.Fprint(r.w, `selected `)
			}
		}

		if radio.Attr == nil {
			goto end
		}

		for _, exc := range attrExclude {
			delete(radio.Attr, exc)
		}

		for key, value := range radio.Attr {
			fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
		}

	end:

		fmt.Fprint(r.w, `/>`)
		if radio.Label != "" {
			fmt.Fprint(r.w, `</label>`)
		}
		fmt.Fprintln(r.w)
	}
}
