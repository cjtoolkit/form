package form

import (
	"fmt"
	"time"
)

func (va validateValue) timeInputTime(value time.Time) {
	formatter := func(t time.Time) (str string) {
		switch va._type {
		case InputDatetime:
			str = fmt.Sprint(t.Format(dateTimeFormat))
		case InputDatetimeLocal:
			str = fmt.Sprint(t.In(va.form.loc).Format(dateTimeLocalFormat))
		case InputTime:
			str = fmt.Sprint(t.In(va.form.loc).Format(timeFormat))
		case InputDate:
			str = fmt.Sprint(t.In(va.form.loc).Format(dateFormat))
		case InputMonth:
			str = fmt.Sprint(t.In(va.form.loc).Format(monthFormat))
		case InputWeek:
			year, week := t.In(va.form.loc).ISOWeek()
			str = fmt.Sprintf(weekFormat, year, week)
		}
		return
	}

	min := time.Time{}
	max := time.Time{}
	minErr := ""
	maxErr := ""

	va.fieldsFns.Call("range_time", map[string]interface{}{
		"min":    &min,
		"max":    &max,
		"minErr": &minErr,
		"maxErr": &maxErr,
	})

	if min.IsZero() && max.IsZero() {
		goto check_mandatory
	} else if min.IsZero() {
		goto check_max
	}

	if value.Unix() < min.Unix() {
		if minErr == "" {
			minErr = va.form.T("ErrTimeMin", map[string]interface{}{
				"Time": formatter(min),
			})
			*(va.err) = fmt.Errorf(minErr)
			return
		}
	}

check_max:

	if max.IsZero() {
		goto check_mandatory
	}

	if value.Unix() > max.Unix() {
		if maxErr == "" {
			maxErr = va.form.T("ErrTimeMax", map[string]interface{}{
				"Time": formatter(max),
			})
			*(va.err) = fmt.Errorf(maxErr)
			return
		}
	}

check_mandatory:

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory && value.IsZero() {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}
}
