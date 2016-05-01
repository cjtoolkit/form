package fields

func BoolSuffix(suffix *string) func(*Bool) {
	return func(b *Bool) {
		b.Suffix = suffix
	}
}

func BoolRequired(required bool, errKey string) func(*Bool) {
	return func(b *Bool) {
		b.Required = required
		b.RequiredErrKey = errKey
	}
}