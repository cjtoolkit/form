package cheatsheet

/*
InputNumber, InputRange and InputHidden
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"range":
		"min": *int64
		"max": *int64
		"minErr": *string
		"maxErr": *string
	"step":
		"step": *int64
		"err": *string
	"attr":
		"attr" (m): *map[string]string

InputRadio
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"radio" (m):
		"radio" (m): *[]form.RadioInt
	"mandatory":
		"mandatory" (m): *bool
		"err": *string

Select (Single)
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"option" (m):
		"option" (m): *[]form.OptionInt
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Int64 int64

/*
Select (Multiple)
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"option" (m):
		"option" (m): *[]form.OptionInt
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Int64s []int64
