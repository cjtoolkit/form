package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorValidateModel struct {
	Msg   string
	Key   string
	Value interface{}
}

func (eVM *ErrorValidateModel) Error() string {
	return eVM.Msg
}

func (eVM *ErrorValidateModel) Translate(language LanguageInterface) {
	eVM.Msg = language.Translate(eVM.Key, eVM.Value)
}
