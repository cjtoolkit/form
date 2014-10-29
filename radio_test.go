package form

func ExampleRadio() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]Radio)) = []Radio{
				{Value: "car", Label: "Car"},
				{Value: "motorbike", Label: "Motorbike"},
			}
		},
	}
}

func ExampleRadioFloat() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]RadioFloat)) = []RadioFloat{
				{Value: 1.5, Label: "Car"},
				{Value: 2.5, Label: "Motorbike"},
			}
		},
	}
}

func ExampleRadioInt() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]RadioInt)) = []RadioInt{
				{Value: 1, Label: "Car"},
				{Value: 2, Label: "Motorbike"},
			}
		},
	}
}
