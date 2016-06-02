package fields

import "regexp"

type StringOption func(*String)

func StringSuffix(suffix *string) StringOption {
	return func(s *String) {
		s.Suffix = suffix
	}
}

func StringRequired(errKey string) StringOption {
	return func(s *String) {
		s.Required = true
		s.RequiredErrKey = errKey
	}
}

func StringMinRune(minRune int, errKey string) StringOption {
	return func(s *String) {
		s.MinRune = minRune
		s.MinRuneErrKey = errKey
	}
}

func StringMaxRune(maxRune int, errKey string) StringOption {
	return func(s *String) {
		s.MaxRune = maxRune
		s.MaxRuneErrKey = errKey
	}
}

func StringMustMatch(name, label string, model *string, errKey string) StringOption {
	return func(s *String) {
		s.MustMatchName = name
		s.MustMatchLabel = label
		s.MustMatchModel = model
		s.MustMatchErrKey = errKey
	}
}

func StringPattern(pattern *regexp.Regexp, errKey string) StringOption {
	return func(s *String) {
		s.Pattern = pattern
		s.PatternErrKey = errKey
	}
}

func StringInList(errKey string, inList ...string) StringOption {
	return func(s *String) {
		s.InListErrKey = errKey
		s.InList = inList
	}
}

func StringExtra(extra func()) StringOption {
	return func(s *String) {
		s.Extra = extra
	}
}
