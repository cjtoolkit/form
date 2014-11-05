package form

import (
	"fmt"
	"math"
)

func (r renderValue) numInputNumber(value interface{}) {
	_type := func() (str string) {
		switch r._type {
		case InputNumber:
			str = "number"
		case InputRange:
			str = "range"
		case InputHidden:
			str = "hidden"
		}
		return
	}

	ovalue := func() (str string) {
		switch value := value.(type) {
		case int64:
			str = fmt.Sprintf(`%d`, value)
		case float64:
			str = fmt.Sprintf(`%f`, value)
		}
		return
	}

	w := r.w

	fmt.Fprintf(w, `<input name="%s" type="%s" value="%s" `, es(r.preferedName), _type(), ovalue())

	switch value.(type) {
	case int64:
		rangeMin := int64(-9223372036854775808)
		rangeMax := int64(9223372036854775807)
		_s := ""

		r.fieldsFns.Call("range", map[string]interface{}{
			"min":    &rangeMin,
			"max":    &rangeMax,
			"minErr": &_s,
			"maxErr": &_s,
		})

		if rangeMin != int64(-9223372036854775808) {
			fmt.Fprintf(w, `min="%d" `, rangeMin)
		}

		if rangeMax != int64(9223372036854775807) {
			fmt.Fprintf(w, `max="%d" `, rangeMax)
		}

		step := int64(1)

		r.fieldsFns.Call("step", map[string]interface{}{
			"step": &step,
			"err":  &_s,
		})

		fmt.Fprintf(w, `step="%d" `, step)
	case float64:
		rangeMin := math.NaN()
		rangeMax := math.NaN()
		_s := ""

		r.fieldsFns.Call("range", map[string]interface{}{
			"min":    &rangeMin,
			"max":    &rangeMax,
			"minErr": &_s,
			"maxErr": &_s,
		})

		if rangeMin != math.NaN() {
			fmt.Fprintf(w, `min="%f" `, rangeMin)
		}

		if rangeMax != math.NaN() {
			fmt.Fprintf(w, `max="%f" `, rangeMax)
		}

		step := float64(1)

		r.fieldsFns.Call("step", map[string]interface{}{
			"step": &step,
			"err":  &_s,
		})

		// Anything below 0.5 will not work, for some reason.
		if step < 0.5 {
			step = 0.5
		}

		fmt.Fprintf(w, `step="%f" `, step)
	}

	var attr map[string]string

	r.fieldsFns.Call("attr", map[string]interface{}{
		"attr": &attr,
	})

	if attr != nil {
		delete(attr, "name")
		delete(attr, "type")
		delete(attr, "value")
		delete(attr, "min")
		delete(attr, "max")
		delete(attr, "step")
		fmt.Fprint(w, RenderAttr(attr))
	}

	fmt.Fprint(w, `/>`)
}
