package form

import (
	"math"
	"regexp"
	"time"
)

// Input Name
func (fns FieldFuncs) Name(name string) {
	fns.Call("set_name", map[string]interface{}{
		"set_name": name,
	})

	fns["set_name"] = nil

	fns["name"] = func(m map[string]interface{}) {
		*(m["name"].(*string)) = name
	}
}

// Make Field Mandatory
func (fns FieldFuncs) Mandatory() *string {
	err := ""

	fns["mandatory"] = func(m map[string]interface{}) {
		*(m["mandatory"].(*bool)) = true
		*(m["err"].(*string)) = err
	}

	return &err
}

// Size of Field
type Size struct {
	Min, Max       int
	MinErr, MaxErr string
}

// Size of Field
func (fns FieldFuncs) Size() *Size {
	size := &Size{-1, -1, "", ""}

	fns["size"] = func(m map[string]interface{}) {
		*(m["min"].(*int)) = size.Min
		*(m["max"].(*int)) = size.Max
		*(m["minErr"].(*string)) = size.MinErr
		*(m["maxErr"].(*string)) = size.MaxErr
	}

	return size
}

// Range of Integer Field
type RangeInt struct {
	Min, Max       int64
	MinErr, MaxErr string
}

// Range of Integer Field
func (fns FieldFuncs) RangeInt() *RangeInt {
	rangeInt := &RangeInt{-9223372036854775808, 9223372036854775807, "", ""}

	fns["range"] = func(m map[string]interface{}) {
		*(m["min"].(*int64)) = rangeInt.Min
		*(m["max"].(*int64)) = rangeInt.Max
		*(m["minErr"].(*string)) = rangeInt.MinErr
		*(m["maxErr"].(*string)) = rangeInt.MaxErr
	}

	return rangeInt
}

// Range of Float Field
type RangeFloat struct {
	Min, Max       float64
	MinErr, MaxErr string
}

// Range of Float Field
func (fns FieldFuncs) RangeFloat() *RangeFloat {
	rangeInt := &RangeFloat{math.NaN(), math.NaN(), "", ""}

	fns["range"] = func(m map[string]interface{}) {
		*(m["min"].(*float64)) = rangeInt.Min
		*(m["max"].(*float64)) = rangeInt.Max
		*(m["minErr"].(*string)) = rangeInt.MinErr
		*(m["maxErr"].(*string)) = rangeInt.MaxErr
	}

	return rangeInt
}

// Range of Time Field
type RangeTime struct {
	Min, Max       time.Time
	MinErr, MaxErr string
}

// Range of Time Field
func (fns FieldFuncs) RangeTime() *RangeTime {
	rangeInt := &RangeTime{time.Time{}, time.Time{}, "", ""}

	fns["range"] = func(m map[string]interface{}) {
		*(m["min"].(*time.Time)) = rangeInt.Min
		*(m["max"].(*time.Time)) = rangeInt.Max
		*(m["minErr"].(*string)) = rangeInt.MinErr
		*(m["maxErr"].(*string)) = rangeInt.MaxErr
	}

	return rangeInt
}

// Number of Step (int64)
func (fns FieldFuncs) StepInt(i int64) *string {
	err := ""

	fns["step"] = func(m map[string]interface{}) {
		*(m["step"].(*int64)) = i
		*(m["err"].(*string)) = err
	}

	return &err
}

// Number of Step Float64
func (fns FieldFuncs) StepFloat(f float64) *string {
	err := ""

	fns["step"] = func(m map[string]interface{}) {
		*(m["step"].(*float64)) = f
		*(m["err"].(*string)) = err
	}

	return &err
}

// Must Match Field
type MustMatch struct {
	Name  string
	Value *string
	Err   string
}

// Must Match Field
func (fns FieldFuncs) MustMatch() *MustMatch {
	mustMatch := &MustMatch{}

	fns["mustmatch"] = func(m map[string]interface{}) {
		*(m["name"].(*string)) = mustMatch.Name
		*(m["value"].(*string)) = *mustMatch.Value
		*(m["err"].(*string)) = mustMatch.Err
	}

	return mustMatch
}

// Regular Expression
func (fns FieldFuncs) Pattern(re *regexp.Regexp) *string {
	err := ""

	fns["pattern"] = func(m map[string]interface{}) {
		*(m["pattern"].(**regexp.Regexp)) = re
		*(m["err"].(*string)) = err
	}

	return &err
}

// Rows and Columns
func (fns FieldFuncs) Textarea(rows, cols int) {
	fns["textarea"] = func(m map[string]interface{}) {
		*(m["rows"].(*int)) = rows
		*(m["cols"].(*int)) = cols
	}
}

// Custom Email Error
func (fns FieldFuncs) EmailError(err string) {
	fns["email"] = func(m map[string]interface{}) {
		*(m["email"].(*string)) = err
	}
}

// Attribute
func (fns FieldFuncs) Attr(attr map[string]string) {
	fns["attr"] = func(m map[string]interface{}) {
		*(m["attr"].(*map[string]string)) = attr
	}
}

// Append Options, accept []Option, []OptionInt, []OptionFloat
func (fns FieldFuncs) Options(v interface{}) {
	fns["option"] = func(m map[string]interface{}) {
		switch v := v.(type) {
		case []Option:
			*(m["option"].(*[]Option)) = v
		case []OptionInt:
			*(m["option"].(*[]OptionInt)) = v
		case []OptionFloat:
			*(m["option"].(*[]OptionFloat)) = v
		}
	}
}

// Append Radios, accept []Radio, []RadioInt, []RadioFloat
func (fns FieldFuncs) Radios(v interface{}) {
	fns["radio"] = func(m map[string]interface{}) {
		switch v := v.(type) {
		case []Radio:
			*(m["radio"].(*[]Radio)) = v
		case []RadioInt:
			*(m["radio"].(*[]RadioInt)) = v
		case []RadioFloat:
			*(m["radio"].(*[]RadioFloat)) = v
		}
	}
}

// File
type File struct {
	Size      int64
	SizeErr   string
	Accept    []string
	AcceptErr string
}

// File
func (fns FieldFuncs) File() *File {
	file := &File{Size: -1}

	fns["file"] = func(m map[string]interface{}) {
		*(m["size"].(*int64)) = file.Size
		*(m["sizeErr"].(*string)) = file.SizeErr
		*(m["accept"].(*[]string)) = file.Accept
		*(m["acceptErr"].(*string)) = file.AcceptErr
	}

	return file
}

// HTML
type HTML struct {
	Before string
	After  string
}

// HTML
func (fns FieldFuncs) HTML() *HTML {
	html := &HTML{}

	fns["html"] = func(m map[string]interface{}) {
		*(m["before"].(*string)) = html.Before
		*(m["after"].(*string)) = html.After
	}

	return html
}

// Label
type Label struct {
	Content, For string
	Attr         map[string]string
}

// Label
func (fns FieldFuncs) Label() *Label {
	label := &Label{}

	fns["html"] = func(m map[string]interface{}) {
		*(m["content"].(*string)) = label.Content
		*(m["for"].(*string)) = label.For
		*(m["attr"].(*map[string]string)) = label.Attr
	}

	return label
}

// Custom Rules
func (fns FieldFuncs) Custom(fn func(*error, *string)) {
	fns["ext"] = func(m map[string]interface{}) {
		fn(m["error"].(*error), m["warning"].(*string))
	}
}
