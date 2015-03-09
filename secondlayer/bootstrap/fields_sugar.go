package bootstrap

import (
	"github.com/cjtoolkit/form"
)

type Bootstrap struct {
	BeforeInput, AfterInput  string
	StartOfGroup, EndOfGroup string
	HelpBlock                string
	Disabled, Feedback       bool
}

// Init Sugar for Bootstrap
func Sugar(fns form.FieldFuncs) *Bootstrap {
	boot := &Bootstrap{}

	fns["bootstrap"] = func(m map[string]interface{}) {
		*(m["beforeInput"].(*string)) = boot.BeforeInput
		*(m["afterInput"].(*string)) = boot.AfterInput
		*(m["startOfGroup"].(*string)) = boot.StartOfGroup
		*(m["endOfGroup"].(*string)) = boot.EndOfGroup
		*(m["helpBlock"].(*string)) = boot.HelpBlock
		*(m["disabled"].(*bool)) = boot.Disabled
		*(m["feedback"].(*bool)) = boot.Feedback
	}

	return boot
}
