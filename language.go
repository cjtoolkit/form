package form

import (
	"bytes"
)

/*
Implement:
	LangaugeInterface
*/
type Langauge map[string]LangaugeTemplateInterface

func (l Langauge) Translate(name string, value interface{}) (msg string) {
	if nil != l[name] {
		buf := &bytes.Buffer{}
		defer buf.Reset()
		l[name].Execute(buf, value)
		msg = buf.String()
	}
	return
}

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
	LANG_FILE_MIME      = "file_mime"
	LANG_FILE_SIZE      = "file_size"
)

func englishLanguage() Langauge {
	T := BuildLanguageTemplate
	return Langauge{
		LANG_FIELD_REQUIRED: T(`'{{.Label}}' is required.`),
		LANG_MIN_CHAR:       T(`'{{.Label}}' should be more than or equal to '{{.MinRune}}' character{{.MinRune|pluralise "s"}}.`),
		LANG_MAX_CHAR:       T(`'{{.Label}}' should be less than or equal '{{.MaxRune}}' character{{.MaxRune|pluralise "s"}}.`),
		LANG_MUST_MATCH:     T(`'{{.Label}}' should match '{{.MustMatchLabel}}.'`),
		LANG_PATTERN:        T(`'{{.Label}}' should match '{{.Pattern}}.'`),
		LANG_IN_LIST:        T(`Value of '{{.Label}}' is not in the list '{{.List|list "and"}}'.`),
		LANG_NOT_INT:        T(`'{{.Label}}' is not a whole number.`),
		LANG_NOT_FLOAT:      T(`'{{.Label}}' is not a decimal.`),
		LANG_NUMBER_MIN:     T(`'{{.Label}}' should be more than or equal to '{{.Min}}'.`),
		LANG_NUMBER_MAX:     T(`'{{.Label}}' should be less than or equal to '{{.Max}}'.`),
		LANG_NUMBER_STEP:    T(`'{{.Label}}' should be in step of '{{.Step}}'.`),
		LANG_TIME_FORMAT:    T(`'{{.Label}}' is not a valid time format.`),
		LANG_TIME_MIN:       T(`'{{.Label}}' should be more than '{{.Min}}'.`),
		LANG_TIME_MAX:       T(`'{{.Label}}' should be less than '{{.Max}}'.`),
		LANG_FILE_MIME:      T(`'{{.Label}}' should be '{{.Mime|list "and"}}'.`),
		LANG_FILE_SIZE:      T(`'{{.Label}}' should be less '{{.Size}}' in bytes.`),
	}
}

var englishLanguageMap = englishLanguage()

func AddToEnglishLanguageMap(key string, tmp LangaugeTemplateInterface) {
	englishLanguageMap[key] = tmp
}
