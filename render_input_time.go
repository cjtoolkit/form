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

	input := &FirstLayerInput{}
	r.fls.append(input)

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
		input.Attr = attr
	} else {
		input.Attr = map[string]string{}
	}

	formatter := func(t time.Time) (str string) {
		if t.IsZero() {
			return
		}

		switch r._type {
		case InputDatetime:
			str = fmt.Sprint(t.Format(dateTimeFormat))
		case InputDatetimeLocal:
			str = fmt.Sprint(t.In(r.form.loc).Format(dateTimeLocalFormat))
		case InputTime:
			str = fmt.Sprint(t.In(r.form.loc).Format(timeFormat))
		case InputDate:
			str = fmt.Sprint(t.In(r.form.loc).Format(dateFormat))
		case InputMonth:
			str = fmt.Sprint(t.In(r.form.loc).Format(monthFormat))
		case InputWeek:
			year, week := t.In(r.form.loc).ISOWeek()
			str = fmt.Sprintf(weekFormat, year, week)
		}

		return
	}

	input.Attr["name"] = r.preferedName
	input.Attr["type"] = _type()
	input.Attr["value"] = formatter(value)

	var rangeTime *RangeTime

	r.fieldsFns.Call("range_time", map[string]interface{}{
		"range": &rangeTime,
	})

	if rangeTime != nil {
		if !rangeTime.Min.IsZero() {
			input.Attr["min"] = formatter(rangeTime.Min)
		}

		if !rangeTime.Max.IsZero() {
			input.Attr["max"] = formatter(rangeTime.Max)
		}
	}

	// Todo: add support for step.

}
