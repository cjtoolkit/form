package fields

type StringSliceOption func(*StringSlice)

func StringSliceSuffix(suffix *string) StringSliceOption {
	return func(s *StringSlice) {
		s.Suffix = suffix
	}
}

func StringSliceRequired(errKey string) StringSliceOption {
	return func(s *StringSlice) {
		s.Required = true
		s.RequiredErrKey = errKey
	}
}

func StringSliceExtra(extra func()) StringSliceOption {
	return func(s *StringSlice) {
		s.Extra = extra
	}
}
