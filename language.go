package form

import (
	"bytes"
)

/*
Implement:
	LangaugeInterface
*/
type Langauge map[string]LangaugeTemplateInterface

const (
	LANG_FIELD_REQUIRED = "field_required"
	LANG_MIN_CHAR       = "min_char"
	LANG_MAX_CHAR       = "max_char"
	LANG_MUST_MATCH     = "must_match"
	LANG_PATTERN        = "pattern"
	LANG_IN_LIST        = "in_list"
	LANG_NOT_INT        = "not_int"
)

func DefaultLanguage() Langauge {
	bLT := BuildLanguageTemplate
	return Langauge{
		LANG_FIELD_REQUIRED: bLT(`'{{.Label}}' is required.`),
		LANG_MIN_CHAR:       bLT(`'{{.Label}}' should be greater than '{{.MinChar}}' character{{.MinChar|pluralise "s"}}.`),
		LANG_MAX_CHAR:       bLT(`'{{.Label}}' should be less than '{{.MaxChar}}' character{{.MaxChar|pluralise "s"}}.`),
		LANG_MUST_MATCH:     bLT(`'{{.Label}}' should match '{{.MustMatchLabel}}.'`),
		LANG_PATTERN:        bLT(`'{{.Label}}' should match '{{.Pattern}}.'`),
		LANG_IN_LIST:        bLT(`Value of '{{.Label}}' is not in the list '{{.List|list "and"}}'.`),
		LANG_NOT_INT:        bLT(`Value of '{{.Label}}' is not a whole number.`),
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
