package form

import (
	"fmt"
)

func (r render) numInputNumber(_float bool, _type string) {
	fmt.Fprint(r.w, `<input type="`, _type, `" `)

	fmt.Fprint(r.w, `name="`, es(r.preferedName), `" `)

	fmt.Fprint(r.w, ` value="`, r.value.Interface(), `" `)

	var min, max, step interface{}
	var ok bool

	min, ok = r.getInt("Min")
	if !ok {
		if _float {
			min, ok = r.getFloat("Min")
		}
	}
	if ok {
		fmt.Fprint(r.w, `min="`, min, `" `)
	}

	max, ok = r.getInt("Max")
	if !ok {
		if _float {
			max, ok = r.getFloat("Max")
		}
	}
	if ok {
		fmt.Fprint(r.w, `max="`, max, `" `)
	}

	step, ok = r.getInt("Step")
	if !ok {
		if _float {
			step, ok = r.getFloat("Step")
			if !ok {
				fmt.Fprint(r.w, `step="any" `)
			}
		}
	}
	if ok {
		fmt.Fprint(r.w, `step="`, step, `" `)
	}

	r.attr("type", "name", "max", "min", "step")

	fmt.Fprint(r.w, `/>`)
}
