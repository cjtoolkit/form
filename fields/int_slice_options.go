package fields

func IntSliceSuffix(suffix *string) func(*IntSlice) {
	return func(i *IntSlice) {
		i.Suffix = suffix
	}
}

func IntSliceRequired(errKey string) func(*IntSlice) {
	return func(i *IntSlice) {
		i.Required = true
		i.RequiredErrKey = errKey
	}
}

func IntSliceExtra(extra func()) func(*IntSlice) {
	return func(i *IntSlice) {
		i.Extra = extra
	}
}
