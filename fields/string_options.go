package fields

import "regexp"

func StringSuffix(suffix *string) func(*String) {
	return func(s *String) {
		s.Suffix = suffix
	}
}

func StringRequired(errKey string) func(*String) {
	return func(s *String) {
		s.Required = true
		s.RequiredErrKey = errKey
	}
}

func StringMinRune(minRune int, errKey string) func(*String) {
	return func(s *String) {
		s.MinRune = minRune
		s.MinRuneErrKey = errKey
	}
}

func StringMaxRune(maxRune int, errKey string) func(*String) {
	return func(s *String) {
		s.MaxRune = maxRune
		s.MaxRuneErrKey = errKey
	}
}

func StringMustMatch(name, label string, model *string, errKey string) func(*String) {
	return func(s *String) {
		s.MustMatchName = name
		s.MustMatchLabel = label
		s.MustMatchModel = model
		s.MustMatchErrKey = errKey
	}
}

func StringPattern(pattern *regexp.Regexp, errKey string) func(*String) {
	return func(s *String) {
		s.Pattern = pattern
		s.PatternErrKey = errKey
	}
}

func StringInList(errKey string, inList ...string) func(*String) {
	return func(s *String) {
		s.InListErrKey = errKey
		s.InList = inList
	}
}

func StringExtra(extra func()) func(*String) {
	return func(s *String) {
		s.Extra = extra
	}
}
