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

	input := &FirstLayerInput{}
	r.fls.append(input)

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
		delete(attr, "required")
		input.Attr = attr
	} else {
		input.Attr = map[string]string{}
	}

	input.Attr["name"] = r.preferedName
	input.Attr["type"] = _type()
	input.Attr["value"] = ovalue()

	switch value.(type) {
	case int64:
		rangeMin := int64(-9223372036854775808)
		rangeMax := int64(9223372036854775807)
		_s := ""

		r.fieldsFns.Call("range_int", map[string]interface{}{
			"min":    &rangeMin,
			"max":    &rangeMax,
			"minErr": &_s,
			"maxErr": &_s,
		})

		if rangeMin != int64(-9223372036854775808) {
			input.Attr["min"] = fmt.Sprintf("%d", rangeMin)
		}

		if rangeMax != int64(9223372036854775807) {
			input.Attr["max"] = fmt.Sprintf("%d", rangeMax)
		}

		step := int64(1)

		r.fieldsFns.Call("step_int", map[string]interface{}{
			"step": &step,
			"err":  &_s,
		})

		input.Attr["step"] = fmt.Sprintf("%d", step)
	case float64:
		rangeMin := math.NaN()
		rangeMax := math.NaN()
		_s := ""

		r.fieldsFns.Call("range_float", map[string]interface{}{
			"min":    &rangeMin,
			"max":    &rangeMax,
			"minErr": &_s,
			"maxErr": &_s,
		})

		if rangeMin != math.NaN() {
			input.Attr["min"] = fmt.Sprintf("%f", rangeMin)
		}

		if rangeMax != math.NaN() {
			input.Attr["max"] = fmt.Sprintf("%f", rangeMax)
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

		input.Attr["step_float"] = fmt.Sprintf("%f", step)
	}

	_s := ""
	mandatory := false

	r.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &_s,
	})

	if mandatory {
		input.Attr["required"] = " "
	}
}
