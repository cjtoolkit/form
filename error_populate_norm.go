package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorPopulateNorm struct {
	msg   string
	Key   string
	Value interface{}
}

func (ePN *ErrorPopulateNorm) Error() string {
	return ePN.msg
}

func (ePN *ErrorPopulateNorm) Translate(language Langauge) {
	ePN.msg = language.Translate(ePN.Key, ePN.Value)
}
