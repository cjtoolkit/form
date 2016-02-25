package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorReverseTransform struct {
	msg   string
	Key   string
	Value interface{}
}

func (eRT *ErrorReverseTransform) Error() string {
	return eRT.msg
}

func (eRT *ErrorReverseTransform) Translate(language Langauge) {
	eRT.msg = language.Translate(eRT.Key, eRT.Value)
}
