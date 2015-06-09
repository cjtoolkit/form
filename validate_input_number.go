package form

import (
	"fmt"
	"math"
)

func (va validateValue) wnumInputNumber(value int64) {
	var rangeInt *RangeInt

	va.fieldsFns.Call("range_int", map[string]interface{}{
		"range": &rangeInt,
	})

	if rangeInt == nil || (rangeInt.Min == minInt64 && rangeInt.Max == maxInt64) {
		goto dostep
	} else if rangeInt.Min == minInt64 {
		goto doMax
	}

	if value < rangeInt.Min {
		if rangeInt.MinErr == "" {
			rangeInt.MinErr = va.form.T("ErrNumberMin", map[string]interface{}{
				"Count": rangeInt.Min,
			})
		}
		*(va.err) = fmt.Errorf(rangeInt.MinErr)
		return
	}

doMax:

	if rangeInt.Max == maxInt64 {
		goto dostep
	}

	if value > rangeInt.Max {
		if rangeInt.MaxErr == "" {
			rangeInt.MaxErr = va.form.T("ErrNumberMax", map[string]interface{}{
				"Count": rangeInt.Max,
			})
		}
		*(va.err) = fmt.Errorf(rangeInt.MaxErr)
		return
	}

dostep:

	step := int64(1)
	stepErr := ""

	va.fieldsFns.Call("step_int", map[string]interface{}{
		"step": &step,
		"err":  &stepErr,
	})

	if value%step != 0 {
		if stepErr == "" {
			stepErr = va.form.T("ErrNumberStep", map[string]interface{}{
				"Count": step,
			})
		}
		*(va.err) = fmt.Errorf(stepErr)
		return
	}
}

func (va validateValue) fnumInputNumber(value float64) {
	var rangeFloat *RangeFloat

	va.fieldsFns.Call("range_float", map[string]interface{}{
		"range": &rangeFloat,
	})

	if rangeFloat == nil || (rangeFloat.Min == math.NaN() && rangeFloat.Max == math.NaN()) {
		goto dostep
	} else if rangeFloat.Min == math.NaN() {
		goto doMax
	}

	if value < rangeFloat.Min {
		if rangeFloat.MinErr == "" {
			rangeFloat.MinErr = va.form.T("ErrNumberMin", map[string]interface{}{
				"Count": rangeFloat.Min,
			})
		}
		*(va.err) = fmt.Errorf(rangeFloat.MinErr)
		return
	}

doMax:

	if rangeFloat.Max == math.NaN() {
		goto dostep
	}

	if value > rangeFloat.Max {
		if rangeFloat.MaxErr == "" {
			rangeFloat.MaxErr = va.form.T("ErrNumberMax", map[string]interface{}{
				"Count": rangeFloat.Max,
			})
		}
		*(va.err) = fmt.Errorf(rangeFloat.MaxErr)
		return
	}

dostep:

	step := float64(1)
	stepErr := ""

	va.fieldsFns.Call("step_float", map[string]interface{}{
		"step": &step,
		"err":  &stepErr,
	})

	// Anything below 0.5 will not work, for some reason.
	if step < 0.5 {
		step = 0.5
	}

	if math.Mod(value, step) != 0 {
		if stepErr == "" {
			stepErr = va.form.T("ErrNumberStep", map[string]interface{}{
				"Count": step,
			})
		}
		*(va.err) = fmt.Errorf(stepErr)
		return
	}
}
