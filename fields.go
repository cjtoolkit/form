package form

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
	name  string
	funcs FieldFuncs
}

// Fields
type Fields struct {
	f []*Field
}

// Init Field
func (f *Fields) Init(fieldname string, typeCode TypeCode) FieldFuncs {
	fns := FieldFuncs{
		"init": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = typeCode
		},
	}
	f.f = append(f.f, &Field{fieldname, fns})
	return fns
}
