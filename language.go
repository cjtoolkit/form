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
	LANG_NOT_FLOAT      = "not_float"
	LANG_NUMBER_MIN     = "number_min"
	LANG_NUMBER_MAX     = "number_max"
	LANG_NUMBER_STEP    = "number_step"
	LANG_TIME_FORMAT    = "time_format"
	LANG_TIME_MIN       = "time_min"
	LANG_TIME_MAX       = "time_max"
)

func defaultLanguage() Langauge {
	bLT := BuildLanguageTemplate
	return Langauge{
		LANG_FIELD_REQUIRED: bLT(`'{{.Label}}' is required.`),
		LANG_MIN_CHAR:       bLT(`'{{.Label}}' should be more than or equal to '{{.MinRune}}' character{{.MinRune|pluralise "s"}}.`),
		LANG_MAX_CHAR:       bLT(`'{{.Label}}' should be less than or equal '{{.MaxRune}}' character{{.MaxRune|pluralise "s"}}.`),
		LANG_MUST_MATCH:     bLT(`'{{.Label}}' should match '{{.MustMatchLabel}}.'`),
		LANG_PATTERN:        bLT(`'{{.Label}}' should match '{{.Pattern}}.'`),
		LANG_IN_LIST:        bLT(`Value of '{{.Label}}' is not in the list '{{.List|list "and"}}'.`),
		LANG_NOT_INT:        bLT(`'{{.Label}}' is not a whole number.`),
		LANG_NOT_FLOAT:      bLT(`'{{.Label}}' is not a decimal.`),
		LANG_NUMBER_MIN:     bLT(`'{{.Label}}' should be more than or equal to '{{.Min}}'.`),
		LANG_NUMBER_MAX:     bLT(`'{{.Label}}' should be less than or equal to '{{.Max}}'.`),
		LANG_NUMBER_STEP:    bLT(`'{{.Label}}' should be in step of '{{.Step}}'.`),
		LANG_TIME_FORMAT:    bLT(`'{{.Label}}' is not a valid time format.`),
		LANG_TIME_MIN:       bLT(`'{{.Label}}' should be more than '{{.Min}}'`),
		LANG_TIME_MAX:       bLT(`'{{.Label}}' should be less than '{{.Max}}'`),
	}
}

var defaultLanguageMap = defaultLanguage()

func (l Langauge) Translate(name string, value interface{}) (msg string) {
	if nil != l[name] {
		buf := &bytes.Buffer{}
		defer buf.Reset()
		l[name].Execute(buf, value)
		msg = buf.String()
	}
	return
}
