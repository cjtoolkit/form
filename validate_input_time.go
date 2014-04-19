package form

import (
	"fmt"
	"time"
)

func (va validate) timeInputTime(_type string) {
	_time := va.value.Interface().(time.Time)

	formatter := func(t time.Time) (str string) {
		switch _type {
		case "datetime":
			str = fmt.Sprint(t.Format(dateTimeFormat))
		case "datetime-local":
			str = fmt.Sprint(t.Format(dateTimeLocalFormat))
		case "time":
			str = fmt.Sprint(t.Format(timeFormat))
		case "date":
			str = fmt.Sprint(t.Format(dateFormat))
		case "month":
			str = fmt.Sprint(t.Format(monthFormat))
		case "week":
			year, week := t.ISOWeek()
			str = fmt.Sprintf(weekFormat, year, week)
		}
		return
	}

	if min, ok := va.getTime("Min"); ok {
		if _time.Unix() < min.Unix() {
			if minErr, ok := va.getStr("MinErr"); ok {
				va.setErr(FormError(minErr))
			} else {
				va.setErr(FormError(fmt.Sprintf(va.i18n.Key(ErrTimeMin), formatter(min))))
			}
			return
		}
	}

	if max, ok := va.getTime("Max"); ok {
		if _time.Unix() > max.Unix() {
			if maxErr, ok := va.getStr("MaxErr"); ok {
				va.setErr(FormError(maxErr))
			} else {
				va.setErr(FormError(fmt.Sprintf(va.i18n.Key(ErrTimeMax), formatter(max))))
			}
			return
		}
	}

	if mandatory, ok := va.getBool("Mandatory"); ok {
		if _time.Unix() == -62135596800 && mandatory {
			if manErr, ok := va.getStr("MandatoryErr"); ok {
				va.setErr(FormError(manErr))
			} else {
				va.setErr(FormError(va.i18n.Key(ErrMandatory)))
			}
			return
		}
	}

	va.callExt()
}
