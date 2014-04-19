package form

import (
	"fmt"
)

func (r render) strInputText(_type string) {
	fmt.Fprint(r.w, `<input name="`, es(r.preferedName), `" type="`, _type, `" `)

	fmt.Fprint(r.w, `value="`, es(r.value.String()), `" `)

	var pattern string

	regExp, ok := r.getRegExp("RegExp")
	if !ok {
		regExp, ok = r.getRegExp("Pattern")
	}
	if ok {
		pattern = regExp.String()
		fmt.Fprint(r.w, `pattern="`, es(pattern), `" `)
		goto max_char_check
	}

	pattern, ok = r.getStr("RegExp")
	if !ok {
		pattern, ok = r.getStr("Pattern")
	}
	if ok {
		fmt.Fprint(r.w, `pattern="`, es(pattern), `" `)
	}

max_char_check:

	maxChar, ok := r.getInt("MaxChar")
	if !ok {
		maxChar, ok = r.getInt("MaxLenght")
	}
	if ok {
		fmt.Fprint(r.w, `maxlength="`, maxChar, `" `)
	}

	r.attr("type", "value", "pattern", "maxlength")

	fmt.Fprint(r.w, `/>`)
}
