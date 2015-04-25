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
	m  map[string]FieldFuncs
	n  map[string]*Field
	nm map[string]*Field
	f  []*Field
}

// Init Field
func (f *Fields) Init(fieldname string, typeCode TypeCode) FieldFuncs {
	if f.m[fieldname] != nil {
		return f.m[fieldname]
	}

	fns := FieldFuncs{
		"init": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = typeCode
		},
	}
	afield := &Field{fieldname, fns}
	if f.nm != nil {
		f.nm[fieldname] = afield
	}
	if f.n != nil {
		fns["set_name"] = func(m map[string]interface{}) {
			name := m["set_name"].(string)
			f.n[name] = afield
		}
	}
	f.m[fieldname] = fns
	if f.f != nil {
		f.f = append(f.f, afield)
	}
	return fns
}
