package fields

import "time"

type TimeOption func(*Time)

func TimeSuffix(suffix *string) TimeOption {
	return func(t *Time) {
		t.Suffix = suffix
	}
}

func TimeUserLocation(UserLocation *string) TimeOption {
	return func(t *Time) {
		t.UserLocation = UserLocation
	}
}

func TimeRequired(errKey string) TimeOption {
	return func(t *Time) {
		t.Required = true
		t.RequiredErrKey = errKey
	}
}

func TimeMin(min time.Time, errKey string) TimeOption {
	return func(t *Time) {
		t.Min = min
		t.MinZero = true
		t.MinErrKey = errKey
	}
}

func TimeMax(max time.Time, errKey string) TimeOption {
	return func(t *Time) {
		t.Max = max
		t.MaxZero = true
		t.MaxErrKey = errKey
	}
}

func TimeExtra(extra func()) TimeOption {
	return func(t *Time) {
		t.Extra = extra
	}
}
