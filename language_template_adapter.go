package form

import (
	"io"
)

/*
Implement:
	LangaugeTemplateInterface
*/
type LangaugeTemplateAdapter func(wr io.Writer, data interface{}) (err error)

func (lTA LangaugeTemplateAdapter) Execute(wr io.Writer, data interface{}) (err error) {
	return lTA(wr, data)
}
