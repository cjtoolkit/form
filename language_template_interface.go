package form

import (
	"io"
)

/*
For:
	*Template in "text/template"
	*Template in "html/template"
*/
type LangaugeTemplateInterface interface {
	Execute(wr io.Writer, data interface{}) (err error)
}
