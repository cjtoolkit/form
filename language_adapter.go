package form

/*
Implement:
	LangaugeInterface
*/
type LanguageAdapter func(name string, value interface{}) (msg string)

func (lA LanguageAdapter) Translate(name string, value interface{}) (msg string) {
	return lA(name, value)
}
