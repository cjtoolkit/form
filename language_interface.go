package form

type LanguageInterface interface {
	Translate(name string, value interface{}) (msg string)
}
