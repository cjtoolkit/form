package fields

func UseDefaultKeyIfCustomKeyIsEmpty(defaultKey, customKey string) string {
	if "" != customKey {
		return customKey
	}
	return defaultKey
}

func ExecFuncIfErrIsNotNil(err error, fn func()) (b bool) {
	if nil != err {
		fn()
		b = true
	}
	return
}
