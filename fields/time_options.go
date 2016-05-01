package fields

import "time"

func TimeSuffix(suffix *string) func(*Time) {
	return func(t *Time) {
		t.Suffix = suffix
	}
}

func TimeRequired(errKey string) func(*Time) {
	return func(t *Time) {
		t.Required = true
		t.RequiredErrKey = errKey
	}
}

func TimeMin(min time.Time, errKey string) func(*Time) {
	return func(t *Time) {
		t.Min = min
		t.MinZero = true
		t.MinErrKey = errKey
	}
}

func TimeMax(max time.Time, errKey string) func(*Time) {
	return func(t *Time) {
		t.Max = max
		t.MaxZero = true
		t.MaxErrKey = errKey
	}
}

func TimeExtra(extra func()) func(*Time) {
	return func(t *Time) {
		t.Extra = extra
	}
}
