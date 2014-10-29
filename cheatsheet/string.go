package cheatsheet

/*
InputText, InputPassword, InputSearch, InputHidden, InputUrl and InputTel
	"form" (m):
		"type" (m):*form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"mustmatch":
		"name" (m): *string
		"value" (m): *string
		"err": *string
	"size":
		"min": *int
		"max": *int
		"minErr": *string
		"maxErr": *string
	"pattern":
		"pattern" (m): **regexp.Regexp
		"err": *string
	"attr":
		"attr" (m): *map[string]string

InputEmail
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"mustmatch":
		"name" (m): *string
		"value" (m): *string
		"err": *string
	"size":
		"min": *int
		"max": *int
		"minErr": *string
		"maxErr": *string
	"email":
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
		"radio" (m): *[]form.Radio
	"mandatory":
		"mandatory" (m): *bool
		"err": *string

Textarea
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"size":
		"min": *int
		"max": *int
		"minErr": *string
		"maxErr": *string
	"textarea":
		"rows": *int
		"cols": *int
	"attr":
		"attr" (m): *map[string]string

Select (Single)
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"option" (m):
		"option" (m): *[]form.Option
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type String string

/*
Select (Multiple)
	"form" (m):
		"type" (m): *form.TypeCode
		"name": *string
	"ext":
		"error": *error
		"warning": *string
	"option" (m):
		"option" (m): *[]form.Option
	"mandatory":
		"mandatory" (m): *bool
		"err": *string
	"attr":
		"attr" (m): *map[string]string
*/
type Strings []string
