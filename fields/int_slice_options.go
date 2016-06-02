package fields

type IntSliceOption func(*IntSlice)

func IntSliceSuffix(suffix *string) IntSliceOption {
	return func(i *IntSlice) {
		i.Suffix = suffix
	}
}

func IntSliceRequired(errKey string) IntSliceOption {
	return func(i *IntSlice) {
		i.Required = true
		i.RequiredErrKey = errKey
	}
}

func IntSliceExtra(extra func()) IntSliceOption {
	return func(i *IntSlice) {
		i.Extra = extra
	}
}
