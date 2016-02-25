package form

/*
Implement:
	TranslatableErrorInterface
	error in "builtin"
*/
type ErrorValidateModel struct {
	msg   string
	Key   string
	Value interface{}
}

func (eVM *ErrorValidateModel) Error() string {
	return eVM.msg
}

func (eVM *ErrorValidateModel) Translate(language Langauge) {
	eVM.msg = language.Translate(eVM.Key, eVM.Value)
}
