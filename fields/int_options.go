package fields

type IntOption func(*Int)

func IntSuffix(suffix *string) IntOption {
	return func(i *Int) {
		i.Suffix = suffix
	}
}

func IntRequired(errKey string) IntOption {
	return func(i *Int) {
		i.Required = true
		i.RequiredErrKey = errKey
	}
}

func IntMin(min int64, errKey string) IntOption {
	return func(i *Int) {
		i.Min = min
		i.MinZero = true
		i.MinErrKey = errKey
	}
}

func IntMax(max int64, errKey string) IntOption {
	return func(i *Int) {
		i.Max = max
		i.MaxZero = true
		i.MaxErrKey = errKey
	}
}

func IntStep(step int64, errKey string) IntOption {
	return func(i *Int) {
		i.Step = step
		i.StepErrKey = errKey
	}
}

func IntInList(errKey string, inList ...int64) IntOption {
	return func(i *Int) {
		i.IsListErrKey = errKey
		i.InList = inList
	}
}

func IntExtra(extra func()) IntOption {
	return func(i *Int) {
		i.Extra = extra
	}
}
