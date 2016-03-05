package form

import (
	"bytes"
	text "text/template"
)

const (
	LANG_FIELD_REQUIRED = "field_required"
	LANG_MIN_CHAR       = "min_char"
	LANG_MAX_CHAR       = "max_char"
	LANG_MUST_MATCH     = "must_match"
	LANG_PATTERN        = "pattern"
)

/*
Implement:
	LangaugeInterface
*/
type Langauge map[string]*text.Template

func DefaultLanguage() Langauge {
	return Langauge{
		LANG_FIELD_REQUIRED: BuildLanguageTemplate("'{{.Label}}' is required."),
		LANG_MIN_CHAR:       BuildLanguageTemplate("'{{.Label}}' should be greater than '{{.MinChar}}'"),
		LANG_MAX_CHAR:       BuildLanguageTemplate("'{{.Label}}' should be less than '{{.MaxChar}}'"),
		LANG_MUST_MATCH:     BuildLanguageTemplate("'{{.Label}}' should match '{{.MustMatchLabel}}'"),
		LANG_PATTERN:        BuildLanguageTemplate("'{{.Label}}' should match '{{.Pattern}}'"),
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
