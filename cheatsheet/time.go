package cheatsheet

import (
	"time"
)

/*
InputDatetime, InputDatetimeLocal, InputTime, InputDate, InputMonth and InputWeek
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"range":
		"min": *time.Time
		"max": *time.Time
		"minErr": *string
		"maxErr": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Time time.Time
