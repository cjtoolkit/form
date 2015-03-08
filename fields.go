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

type Fields map[string]FieldFuncs

func (f Fields) Init(fieldname string, typeCode TypeCode) FieldFuncs {
	f[fieldname] = FieldFuncs{
		"init": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = typeCode
		},
	}
	return f[fieldname]
}
