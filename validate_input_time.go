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

	var rangeTime *RangeTime

	va.fieldsFns.Call("range_time", map[string]interface{}{
		"range": &rangeTime,
	})

	if rangeTime == nil || (rangeTime.Min.IsZero() && rangeTime.Max.IsZero()) {
		goto check_mandatory
	} else if rangeTime.Min.IsZero() {
		goto check_max
	}

	if value.Unix() < rangeTime.Min.Unix() {
		if rangeTime.MinErr == "" {
			rangeTime.MinErr = va.form.T("ErrTimeMin", map[string]interface{}{
				"Time": formatter(rangeTime.Min),
			})
			*(va.err) = fmt.Errorf(rangeTime.MinErr)
			return
		}
	}

check_max:

	if rangeTime.Max.IsZero() {
		goto check_mandatory
	}

	if value.Unix() > rangeTime.Max.Unix() {
		if rangeTime.MaxErr == "" {
			rangeTime.MaxErr = va.form.T("ErrTimeMax", map[string]interface{}{
				"Time": formatter(rangeTime.Max),
			})
			*(va.err) = fmt.Errorf(rangeTime.MaxErr)
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
