package fields

func UseDefaultKeyIfCustomKeyIsEmpty(defaultKey, customKey string) string {
	if "" != customKey {
		return customKey
	}
	return defaultKey
}
