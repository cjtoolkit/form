package form

import (
	"fmt"
	"strings"
)

// For keeping a list of enclosed functions for struct pointer fields.
type FieldFuncs map[string]func(m map[string]interface{})

// Attemp to call a function in FieldFuncs. Does not call if function does not exist.
func (fns FieldFuncs) Call(name string, m map[string]interface{}) {
	if fns[name] == nil {
		return
	}
	fns[name](m)
}

type Field struct {
	ptr      interface{}
	name     string
	funcs    FieldFuncs
	typecode TypeCode
}

// Fields
type Fields struct {
	m          map[string]FieldFuncs
	n          map[string]*Field
	f          []*Field
	validating bool
}

// Init Field
func (f *Fields) Init(ptr interface{}, inputName string, typeCode TypeCode) FieldFuncs {
	if ptr == nil {
		panic(fmt.Errorf("form: field init: ptr cannot be 'nil'!"))
	}

	if !isStructPtr(ptr) {
		panic(fmt.Errorf("form: field init: ptr: '%T' is not a pointer!", ptr))
	}

	inputName = strings.TrimSpace(inputName)

	if inputName == "" {
		panic(fmt.Errorf("form: field init: inputName cannot be blank"))
	}

	if f.m[inputName] != nil {
		return f.m[inputName]
	}

	fns := FieldFuncs{}
	field := &Field{}

	field.ptr = ptr
	field.name = inputName
	field.funcs = fns
	field.typecode = typeCode

	f.m[inputName] = fns
	if f.f != nil {
		f.f = append(f.f, field)
	}

	if f.n != nil {
		f.n[inputName] = field
	}

	return fns
}

// Is it in progress of validation?
func (f *Fields) Validating() bool {
	return f.validating
}
