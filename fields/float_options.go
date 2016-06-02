package fields

type FloatOption func(*Float)

func FloatSuffix(suffix *string) FloatOption {
	return func(f *Float) {
		f.Suffix = suffix
	}
}

func FloatRequired(errKey string) FloatOption {
	return func(f *Float) {
		f.Required = true
		f.RequiredErrKey = errKey
	}
}

func FloatMin(min float64, errKey string) FloatOption {
	return func(f *Float) {
		f.Min = min
		f.MinZero = true
		f.MinErrKey = errKey
	}
}

func FloatMax(max float64, errKey string) FloatOption {
	return func(f *Float) {
		f.Max = max
		f.MaxZero = true
		f.MaxErrKey = errKey
	}
}

func FloatStep(step float64, errKey string) FloatOption {
	return func(f *Float) {
		f.Step = step
		f.StepErrKey = errKey
	}
}

func FloatInList(errKey string, inList ...float64) FloatOption {
	return func(f *Float) {
		f.InListErrKey = errKey
		f.InList = inList
	}
}

func FloatExtra(extra func()) FloatOption {
	return func(f *Float) {
		f.Extra = extra
	}
}
