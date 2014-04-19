package form

import (
	"encoding/gob"
	"reflect"
	"regexp"
	"time"
)

type FormData struct {
	Errors  map[string]error  `json:"-" xml:"-" form:"-"`
	Warning map[string]string `json"-" xml:"-" form: "-"`
	Checked bool              `json"-" xml:"-" form: "-"`
}

// Use that as an amonynous field for creating form.
type Form struct {
	Data *FormData
}

// Get Error
func (f *Form) Err(name string) error {
	if f.Data == nil {
		return nil
	}
	if f.Data.Errors == nil {
		return nil
	}
	return f.Data.Errors[name]
}

// Set Error
func (f *Form) SetErr(name string, err error) {
	if f.Data == nil {
		f.Data = &FormData{}
	}
	if f.Data.Errors == nil {
		f.Data.Errors = map[string]error{}
	}
	f.Data.Errors[name] = err
}

// Has at least one Error
func (f *Form) HasErr() bool {
	if f.Data == nil {
		return false
	}
	return f.Data.Errors != nil
}

// Returns Error Format.
func (f *Form) ErrFormat() string {
	return `<p>%v</p>`
}

// Get Warning
func (f *Form) GetWarning(name string) string {
	if f.Data == nil {
		return ""
	}
	if f.Data.Warning == nil {
		return ""
	}
	return f.Data.Warning[name]
}

// Set Warning
func (f *Form) SetWarning(name, warning string) {
	if f.Data == nil {
		f.Data = &FormData{}
	}
	if f.Data.Warning == nil {
		f.Data.Warning = map[string]string{}
	}
	f.Data.Warning[name] = warning
}

// Returns Warning Format
func (f *Form) WarningFormat() string {
	return `<p>%v</p>`
}

// Return Group Format
func (f *Form) Group() string {
	return `%v`
}

// Return Group Format Success
func (f *Form) GroupSuccess() string {
	return `%v`
}

// Return Group Error Format
func (f *Form) GroupError() string {
	return `%v`
}

// Return Group Warning Format
func (f *Form) GroupWarning() string {
	return `%v`
}

// Wrap around Form Input
func (f *Form) Wrap() string {
	return `%v`
}

// Been Checked
func (f *Form) BeenChecked() bool {
	if f.Data == nil {
		return false
	}
	return f.Data.Checked
}

// Mark as Checked
func (f *Form) Check() {
	if f.Data == nil {
		f.Data = &FormData{}
	}
	f.Data.Checked = true
}

// Form Interface
type FormInterface interface {
	Err(string) error
	SetErr(string, error)
	HasErr() bool
	ErrFormat() string
	GetWarning(string) string
	SetWarning(string, string)
	WarningFormat() string
	Group() string
	GroupSuccess() string
	GroupError() string
	GroupWarning() string
	Wrap() string
	BeenChecked() bool
	Check()
}

// A String that implement the error interface
type FormError string

func (f FormError) Error() string {
	return string(f)
}

func init() {
	gob.Register(&Form{})
	gob.Register(FormError(""))
	gob.Register(&FormData{})
}

type form struct {
	m                         reflect.Value
	t                         reflect.Type
	v                         reflect.Value
	field                     reflect.StructField
	value                     reflect.Value
	name, preferedName, ftype string
}

func (f form) get(suffix string) interface{} {
	m := f.m.MethodByName(f.name + suffix)
	if !m.IsValid() {
		return nil
	}
	in := make([]reflect.Value, 0)
	values := m.Call(in)
	if len(values) == 0 {
		return nil
	}
	return values[0].Interface()
}

func (f form) getStr(suffix string) (string, bool) {
	str, ok := f.get(suffix).(string)
	return str, ok
}

func (f form) getStrs(suffix string) ([]string, bool) {
	strs, ok := f.get(suffix).([]string)
	return strs, ok
}

func (f form) getStrMap(suffix string) (map[string]string, bool) {
	mstr, ok := f.get(suffix).(map[string]string)
	return mstr, ok
}

func (f form) getInt(suffix string) (int64, bool) {
	num, ok := f.get(suffix).(int64)
	return num, ok
}

func (f form) getFloat(suffix string) (float64, bool) {
	num, ok := f.get(suffix).(float64)
	return num, ok
}

func (f form) getBool(suffix string) (bool, bool) {
	b, ok := f.get(suffix).(bool)
	return b, ok
}

func (f form) getRegExp(suffix string) (*regexp.Regexp, bool) {
	re, ok := f.get(suffix).(*regexp.Regexp)
	return re, ok
}

func (f form) getTime(suffix string) (time.Time, bool) {
	_time, ok := f.get(suffix).(time.Time)
	return _time, ok
}
