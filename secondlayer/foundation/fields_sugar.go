package foundation

import (
	"github.com/cjtoolkit/form"
)

type Foundation struct {
	BeforeInput, AfterInput  string
	StartOfGroup, EndOfGroup string
}

// Init Sugar for Bootstrap
func Sugar(fns form.FieldFuncs) *Foundation {
	f := &Foundation{}

	fns["bootstrap"] = func(m map[string]interface{}) {
		*(m["beforeInput"].(*string)) = f.BeforeInput
		*(m["afterInput"].(*string)) = f.AfterInput
		*(m["startOfGroup"].(*string)) = f.StartOfGroup
		*(m["endOfGroup"].(*string)) = f.EndOfGroup
	}

	return f
}
