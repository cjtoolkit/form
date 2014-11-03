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
			str = fmt.Sprint(t.Format(dateTimeLocalFormat))
		case InputTime:
			str = fmt.Sprint(t.Format(timeFormat))
		case InputDate:
			str = fmt.Sprint(t.Format(dateFormat))
		case InputMonth:
			str = fmt.Sprint(t.Format(monthFormat))
		case InputWeek:
			year, week := t.ISOWeek()
			str = fmt.Sprintf(weekFormat, year, week)
		}
		return
	}

	min := time.Time{}
	max := time.Time{}
	minErr := ""
	maxErr := ""

	va.fieldsFns.Call("range", map[string]interface{}{
		"min":    &min,
		"max":    &max,
		"minErr": &minErr,
		"maxErr": &maxErr,
	})

	if min.Unix() == -62135596800 && max.Unix() == -62135596800 {
		goto check_mandatory
	} else if min.Unix() == -62135596800 {
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

	if max.Unix() == -62135596800 {
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

	if mandatory && value.Unix() == -62135596800 {
		if manErr == "" {
			manErr = va.form.T("ErrMandatory")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}
}
