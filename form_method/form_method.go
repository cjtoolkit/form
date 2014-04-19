/*
Generate Method Stub for Form Helper.

Usage
First two argument for method pre reference, the rest are for field names.
*/
package main

import (
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"text/template"
)

type Text struct {
	PreRef string
	Field  string
}

var temp, _ = template.New("Text").Parse(`
// Mandatory, specify the type.
// "input:*", "textarea" and "select"
func ({{.PreRef}}) {{.Field}}Type() string {
	return "input:text"
}

// Specify custom attributes.
func ({{.PreRef}}) {{.Field}}Attr() map[string]string {
	return map[string]string{"class": "form-control"}
}

// Add Label
func ({{.PreRef}}) {{.Field}}Label() form.Label {
	return form.Label{"content", "for", map[string]string{"class": "control-label"}}
}

var _{{.Field}}Pattern = regexp.MustCompile(^([a-zA-Z]*)$")

// For "input:text", "input:password" and "input:search" only
// Specify Pattern in RegExp 
func ({{.PreRef}}) {{.Field}}Pattern() *regexp.Regexp {
	return _{{.Field}}Pattern
}

// For "input:text", "input:password" and "input:search" only
// Specify Error Message for Pattern
func ({{.PreRef}}) {{.Field}}PatternErr() string {
	return "Letters Only"
}

// For "input:text", "input:password" and "input:search" only
// Specify other Field to Match Against
func ({{.PreRef}}) {{.Field}}MustMatch() string {
	return "Field"
}

// For "input:text", "input:password" and "input:search" only
// Specify Error Message for MustMatch
func ({{.PreRef}}) {{.Field}}MustMatchErr() string {
	return "Does not match 'field'"
}

// For "input:text", "input:password", "input:search" and "textarea" only
// Specify Min Length
func ({{.PreRef}}) {{.Field}}MinLength() int64 {
	return 1
}

// For "input:text", "input:password", "input:search" and "textarea" only
// Specify Max Length
func ({{.PreRef}}) {{.Field}}MaxLength() int64 {
	return 10
}

// For "input:text", "input:password", "input:search" and "textarea" only
// Specify Min Length Error Message
func ({{.PreRef}}) {{.Field}}MinLengthErr() string {
	return "Must be 1 or above"
}

// For "input:text", "input:password", "input:search" and "textarea" only
// Specify Max Length Error Message
func ({{.PreRef}}) {{.Field}}MaxLengthErr() string {
	return "Must be 10 or below"
}

// For "textarea" only
// Specify number of Rows
func ({{.PreRef}}) {{.Field}}Rows() int64 {
	return 4
}

// For "textarea" only
// Specify number of Columns
func ({{.PreRef}}) {{.Field}}Cols() int64 {
	return 4
}

// For "input:number" and "input:range" only
// For "int64" you can only use "int64".
// For "float64" you can use either "float64" or "int64"
func ({{.PreRef}}) {{.Field}}Min() int64 {
	return 1
}

// For "input:number" and "input:range" only
// For "int64" you can only use "int64".
// For "float64" you can use either "float64" or "int64"
func ({{.PreRef}}) {{.Field}}Max() int64 {
	return 10
}

// For "input:number" and "input:range" only
// For "int64" you can only use "int64".
// For "float64" you can use either "float64" or "int64"
func ({{.PreRef}}) {{.Field}}Step() int64 {
	return 1
}

// Mandatory for 'input:radio'
// For "string" use "[]Radio"
// For "int64" use "[]RadioInt"
// For "float64" use "[]RadioFloat"
func ({{.PreRef}}) {{.Field}}Radio() []form.Radio {
	return []form.Radio{
		{"car", "Car", false, map[string]string{"class": "car"}},
		{"motorbike", "Motorbike", true, nil},
	}
}

// Mandatory for 'select'
// For "string" use "[]Option"
// For "int64" use "[]OptionInt"
// For "float64" use "[]OptionFloat"
func ({{.PreRef}}) {{.Field}}Options() []form.Option {
	return []form.Option{
		{
			Content: "Car",
			Value:   "car",
			Label:   "car",
		},
		{
			Content:  "Motorcycle",
			Value:    "motorcycle",
			Label:    "motorcycle",
			Selected: true,
			Attr:     map[string]string{"class": "motorcycle"},
		},
	}
}

// For "input:file" only
// Specify File Size Limit in bytes.
func ({{.PreRef}}) {{.Field}}Size() int64 {
	return 10 * 1024 * 1024
}

// For "input:file" only
// Specify File Size Limit Error Message!
func ({{.PreRef}}) {{.Field}}SizeErr() string {
	return "File Size exceed 10mb"
}

// For "input:file" only
// Specify File MIME.
func ({{.PreRef}}) {{.Field}}Accept() []string {
	return []string{"image/jpeg", "image/png", "image/*"}
}

// For adding additional rules, such as database lookups.
func ({{.PreRef}}) {{.Field}}Ext() {

}
`)

func main() {
	if len(os.Args) <= 3 {
		return
	}

	preRef := os.Args[1:3]

	fields := os.Args[3:]

	buf := &bytes.Buffer{}

	text := &Text{
		PreRef: preRef[0] + " " + preRef[1],
	}

	for _, text.Field = range fields {
		temp.Execute(buf, text)
	}

	clipboard.WriteAll(`// Just modify and delete as necessary!
` + buf.String())
	buf.Reset()

	fmt.Println("Method stub are now in your computer clipboard! Just paste it into your document!")
}
