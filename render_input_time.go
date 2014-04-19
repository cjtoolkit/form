package form

import (
	"fmt"
	"time"
)

func (r render) timeInputTime(_type string) {
	_time := r.value.Interface().(time.Time)

	formatter := func(t time.Time) {
		if t.Unix() != -62135596800 {
			switch _type {
			case "datetime":
				fmt.Fprint(r.w, t.Format(dateTimeFormat))
			case "datetime-local":
				fmt.Fprint(r.w, t.Format(dateTimeLocalFormat))
			case "time":
				fmt.Fprint(r.w, t.Format(timeFormat))
			case "date":
				fmt.Fprint(r.w, t.Format(dateFormat))
			case "month":
				fmt.Fprint(r.w, t.Format(monthFormat))
			case "week":
				year, week := t.ISOWeek()
				fmt.Fprintf(r.w, weekFormat, year, week)
			}
		}
	}

	fmt.Fprint(r.w, `<input name="`, es(r.preferedName), `" type="`, _type, `" value="`)

	formatter(_time)

	fmt.Fprint(r.w, `" `)

	if min, ok := r.getTime("Min"); ok {
		fmt.Fprint(r.w, `min="`)
		formatter(min)
		fmt.Fprint(r.w, `" `)
	}

	if max, ok := r.getTime("Max"); ok {
		fmt.Fprint(r.w, `max="`)
		formatter(max)
		fmt.Fprint(r.w, `" `)
	}

	// Todo: add support for step.

	r.attr("type", "name", "max", "min", "step")

	fmt.Fprint(r.w, `/>`)
}
