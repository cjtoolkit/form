package fields

type BoolOption func(*Bool)

func BoolSuffix(suffix *string) BoolOption {
	return func(b *Bool) {
		b.Suffix = suffix
	}
}

func BoolRequired(errKey string) BoolOption {
	return func(b *Bool) {
		b.Required = true
		b.RequiredErrKey = errKey
	}
}
