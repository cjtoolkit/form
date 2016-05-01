package fields

func StringSliceSuffix(suffix *string) func(*StringSlice) {
	return func(s *StringSlice) {
		s.Suffix = suffix
	}
}

func StringSliceRequired(errKey string) func(*StringSlice) {
	return func(s *StringSlice) {
		s.Required = true
		s.RequiredErrKey = errKey
	}
}

func StringSliceExtra(extra func()) func(*StringSlice) {
	return func(s *StringSlice) {
		s.Extra = extra
	}
}
