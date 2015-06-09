package form

import (
	"fmt"
	"math"
)

func (va validateValue) wnumInputNumber(value int64) {
	rangeMin := int64(-9223372036854775808)
	rangeMax := int64(9223372036854775807)
	rangeMinErr := ""
	rangeMaxErr := ""

	va.fieldsFns.Call("range_int", map[string]interface{}{
		"min":    &rangeMin,
		"max":    &rangeMax,
		"minErr": &rangeMinErr,
		"maxErr": &rangeMaxErr,
	})

	if rangeMin == -9223372036854775808 && rangeMax == 9223372036854775807 {
		goto dostep
	} else if rangeMin == -9223372036854775808 {
		goto doMax
	}

	if value < rangeMin {
		if rangeMinErr == "" {
			rangeMinErr = va.form.T("ErrNumberMin", map[string]interface{}{
				"Count": rangeMin,
			})
		}
		*(va.err) = fmt.Errorf(rangeMinErr)
		return
	}

doMax:

	if rangeMax == 9223372036854775807 {
		goto dostep
	}

	if value > rangeMax {
		if rangeMaxErr == "" {
			rangeMaxErr = va.form.T("ErrNumberMax", map[string]interface{}{
				"Count": rangeMax,
			})
		}
		*(va.err) = fmt.Errorf(rangeMaxErr)
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
	rangeMin := math.NaN()
	rangeMax := math.NaN()
	rangeMinErr := ""
	rangeMaxErr := ""

	va.fieldsFns.Call("range_float", map[string]interface{}{
		"min":    &rangeMin,
		"max":    &rangeMax,
		"minErr": &rangeMinErr,
		"maxErr": &rangeMaxErr,
	})

	if rangeMin == math.NaN() && rangeMax == math.NaN() {
		goto dostep
	} else if rangeMin == math.NaN() {
		goto doMax
	}

	if value < rangeMin {
		if rangeMinErr == "" {
			rangeMinErr = va.form.T("ErrNumberMin", map[string]interface{}{
				"Count": rangeMin,
			})
		}
		*(va.err) = fmt.Errorf(rangeMinErr)
		return
	}

doMax:

	if rangeMax == math.NaN() {
		goto dostep
	}

	if value > rangeMax {
		if rangeMaxErr == "" {
			rangeMaxErr = va.form.T("ErrNumberMax", map[string]interface{}{
				"Count": rangeMax,
			})
		}
		*(va.err) = fmt.Errorf(rangeMaxErr)
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
