package fields

func BoolSuffix(suffix *string) func(*Bool) {
	return func(b *Bool) {
		b.Suffix = suffix
	}
}

func BoolRequired(errKey string) func(*Bool) {
	return func(b *Bool) {
		b.Required = true
		b.RequiredErrKey = errKey
	}
}
