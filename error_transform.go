package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorTransform struct {
	msg   string
	Key   string
	Value interface{}
}

func (eT *ErrorTransform) Error() string {
	return eT.msg
}

func (eT *ErrorTransform) Translate(language Langauge) {
	eT.msg = language.Translate(eT.Key, eT.Value)
}
