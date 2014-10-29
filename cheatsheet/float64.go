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
		"min": *float64
		"max": *float64
		"minErr": *string
		"maxErr": *string
	"step":
		"step": *float64
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
		"radio" (m): *[]form.RadiFloat
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
		"option" (m): *[]form.OptionFloat
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Float64 float64

/*
Select (Multiple)
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"option" (m):
		"option" (m): *[]form.OptionFloat
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Float64s []float64
