package form

import (
	"fmt"
	"time"
)

func (r renderValue) timeInputTime(value time.Time) {
	_type := func() (str string) {
		switch r._type {
		case InputDatetime:
			str = "datetime"
		case InputDatetimeLocal:
			str = "datetime-local"
		case InputTime:
			str = "time"
		case InputDate:
			str = "date"
		case InputMonth:
			str = "month"
		case InputWeek:
			str = "week"
		}
		return
	}

	w := r.w

	formatter := func(t time.Time) {
		if t.Unix() != -62135596800 {
			switch r._type {
			case InputDatetime:
				fmt.Fprint(w, es(t.Format(dateTimeFormat)))
			case InputDatetimeLocal:
				fmt.Fprint(w, es(t.Format(dateTimeLocalFormat)))
			case InputTime:
				fmt.Fprint(w, es(t.Format(timeFormat)))
			case InputDate:
				fmt.Fprint(w, es(t.Format(dateFormat)))
			case InputMonth:
				fmt.Fprint(w, es(t.Format(monthFormat)))
			case InputWeek:
				year, week := t.ISOWeek()
				fmt.Fprintf(w, weekFormat, year, week)
			}
		}
	}

	fmt.Fprintf(w, `<input name="%s" type="%s" value="`, es(r.preferedName), _type())

	formatter(value)

	fmt.Fprint(w, `" `)

	min := time.Time{}
	max := time.Time{}
	_s := ""

	r.fieldsFns.Call("range", map[string]interface{}{
		"min":    &min,
		"max":    &max,
		"minErr": &_s,
		"maxErr": &_s,
	})

	if min.Unix() != -62135596800 {
		fmt.Fprint(w, `min="`)
		formatter(min)
		fmt.Fprint(w, `" `)
	}

	if max.Unix() != -62135596800 {
		fmt.Fprint(w, `max="`)
		formatter(max)
		fmt.Fprint(w, `" `)
	}

	// Todo: add support for step.

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "max")
		delete(attr, "min")
		delete(attr, "step")
		fmt.Fprint(w, RenderAttr(attr))
	}

	fmt.Fprint(w, `/>`)
}
