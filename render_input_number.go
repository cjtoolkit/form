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
		var rangeInt *RangeInt
		_s := ""

		r.fieldsFns.Call("range_int", map[string]interface{}{
			"range": &rangeInt,
		})

		if rangeInt != nil {

			if rangeInt.Min != minInt64 {
				input.Attr["min"] = fmt.Sprintf("%d", rangeInt.Min)
			}

			if rangeInt.Max != maxInt64 {
				input.Attr["max"] = fmt.Sprintf("%d", rangeInt.Max)
			}

		}

		step := int64(1)

		r.fieldsFns.Call("step_int", map[string]interface{}{
			"step": &step,
			"err":  &_s,
		})

		input.Attr["step"] = fmt.Sprintf("%d", step)
	case float64:
		var rangeFloat *RangeFloat
		_s := ""

		r.fieldsFns.Call("range_float", map[string]interface{}{
			"range": &rangeFloat,
		})

		if rangeFloat != nil {

			if rangeFloat.Min != math.NaN() {
				input.Attr["min"] = fmt.Sprintf("%f", rangeFloat.Min)
			}

			if rangeFloat.Max != math.NaN() {
				input.Attr["max"] = fmt.Sprintf("%f", rangeFloat.Max)
			}

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
