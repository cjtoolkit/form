package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorTransform struct {
	Msg   string
	Key   string
	Value interface{}
}

func (eT *ErrorTransform) Error() string {
	return eT.Msg
}

func (eT *ErrorTransform) Translate(language Langauge) {
	eT.Msg = language.Translate(eT.Key, eT.Value)
}
