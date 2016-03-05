package form

type TranslatableErrorInterface interface {
	error
	Translate(language LanguageInterface)
}
