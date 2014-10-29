package form

func ExampleOption() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]Option)) = []Option{
				{Content: "Car", Value: "car", Label: "Car"},
				{Content: "Motorbike", Value: "motorbike", Label: "Motorbike"},
			}
		},
	}
}

func ExampleOptionFloat() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]OptionFloat)) = []OptionFloat{
				{Content: "Car", Value: 1.5, Label: "Car"},
				{Content: "Motorbike", Value: 2.5, Label: "Motorbike"},
			}
		},
	}
}

func ExampleOptionInt() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]OptionInt)) = []OptionInt{
				{Content: "Car", Value: 1, Label: "Car"},
				{Content: "Motorbike", Value: 2, Label: "Motorbike"},
			}
		},
	}
}
