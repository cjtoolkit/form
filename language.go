package form

import (
	"bytes"
	text "text/template"
)

type Langauge map[string]*text.Template

func DefaultLanguage() Langauge {
	return Langauge{}
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
