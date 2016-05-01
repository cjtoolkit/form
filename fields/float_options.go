package fields

func FloatSuffix(suffix *string) func(*Float) {
	return func(f *Float) {
		f.Suffix = suffix
	}
}

func FloatRequired(errKey string) func(*Float) {
	return func(f *Float) {
		f.Required = true
		f.RequiredErrKey = errKey
	}
}

func FloatMin(min float64, errKey string) func(*Float) {
	return func(f *Float) {
		f.Min = min
		f.MinZero = true
		f.MinErrKey = errKey
	}
}

func FloatMax(max float64, errKey string) func(*Float) {
	return func(f *Float) {
		f.Max = max
		f.MaxZero = true
		f.MaxErrKey = errKey
	}
}

func FloatStep(step float64, errKey string) func(*Float) {
	return func(f *Float) {
		f.Step = step
		f.StepErrKey = errKey
	}
}

func FloatInList(errKey string, inList ...float64) func(*Float) {
	return func(f *Float) {
		f.InListErrKey = errKey
		f.InList = inList
	}
}

func FloatExtra(extra func()) func(*Float) {
	return func(f *Float) {
		f.Extra = extra
	}
}
