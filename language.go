package form

import (
	"bytes"
	text "text/template"
)

const (
	LANG_FIELD_REQUIRED = "field_required"
)

type Langauge map[string]*text.Template

func DefaultLanguage() Langauge {
	return Langauge{
		LANG_FIELD_REQUIRED: BuildLanguageTemplate("'{{.Label}}' is required."),
	}
}

func (l Langauge) Translate(name string, value interface{}) (msg string) {
	if nil != l[name] {
		buf := &bytes.Buffer{}
		defer buf.Reset()
		l[name].Execute(buf, value)
		msg = buf.String()
	}
	return
}
