package fields

import "regexp"

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

func getMessageFromError(err error) string {
	if nil == err {
		return ""
	}
	return err.Error()
}

func getPatternFromRegExp(re *regexp.Regexp) string {
	if nil == re {
		return ""
	}
	return re.String()
}

func execFnIfNotNil(fn func()) {
	if nil != fn {
		fn()
	}
}

func addSuffix(name string, suffix *string) string {
	if nil != suffix {
		return name + "-" + *suffix
	}
	return name
}
