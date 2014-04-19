package form

import (
	"fmt"
	"math"
)

func (va validate) numInputNumber(_float bool) {
	var min, max, step interface{}
	var bmin, bmax, bstep bool

	min, bmin = va.getInt("Min")
	if !bmin && _float {
		min, bmin = va.getFloat("Min")
	}

	max, bmax = va.getInt("Max")
	if !bmax && _float {
		max, bmax = va.getFloat("Max")
	}

	step, bstep = va.getInt("Step")
	if !bstep && _float {
		step, bstep = va.getFloat("Step")
	}

	minErr, ok := va.getStr("MinErr")
	if !ok {
		minErr = va.i18n.Key(ErrNumberMin)
	}

	maxErr, ok := va.getStr("MaxErr")
	if !ok {
		maxErr = va.i18n.Key(ErrNumberMax)
	}

	stepErr, ok := va.getStr("StepErr")
	if !ok {
		stepErr = va.i18n.Key(ErrNumberStep)
	}

	switch {
	case _float:
		value := va.value.Float()

		if bmin {
			_min := float64(0)
			switch t := min.(type) {
			case float64:
				_min = t
			case int64:
				_min = float64(t)
			}

			if value < _min {
				va.setErr(FormError(fmt.Sprintf(minErr, _min)))
				return
			}
		}

		if bmax {
			_max := float64(0)
			switch t := max.(type) {
			case float64:
				_max = t
			case int64:
				_max = float64(t)
			}

			if value > _max {
				va.setErr(FormError(fmt.Sprintf(maxErr, _max)))
				return
			}
		}

		if bstep {
			_step := float64(0)
			switch t := step.(type) {
			case float64:
				_step = t
			case int64:
				_step = float64(t)
			}

			if math.Remainder(value, _step) != 0 {
				va.setErr(FormError(fmt.Sprintf(stepErr, _step)))
				return
			}
		}

	default:
		value := va.value.Int()

		if bmin {
			_min := min.(int64)

			if value < _min {
				va.setErr(FormError(fmt.Sprintf(minErr, _min)))
				return
			}
		}

		if bmax {
			_max := max.(int64)

			if value > _max {
				va.setErr(FormError(fmt.Sprintf(maxErr, _max)))
				return
			}
		}

		if bstep {
			_step := step.(int64)

			if value%_step != 0 {
				va.setErr(FormError(fmt.Sprintf(stepErr, _step)))
				return
			}
		}
	}

	va.callExt()
}
