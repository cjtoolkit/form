package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorPopulateNorm struct {
	Msg   string
	Key   string
	Value interface{}
}

func (ePN *ErrorPopulateNorm) Error() string {
	return ePN.Msg
}

func (ePN *ErrorPopulateNorm) Translate(language LanguageInterface) {
	ePN.Msg = language.Translate(ePN.Key, ePN.Value)
	ePN.clean()
}

func (ePN *ErrorPopulateNorm) clean() {
	ePN.Key = ""
	ePN.Value = nil
}
