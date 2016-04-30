package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorReverseTransform struct {
	Msg   string
	Key   string
	Value interface{}
}

func (eRT *ErrorReverseTransform) Error() string {
	return eRT.Msg
}

func (eRT *ErrorReverseTransform) Translate(language LanguageInterface) {
	eRT.Msg = language.Translate(eRT.Key, eRT.Value)
}
