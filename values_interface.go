package form

type ValuesInterface interface {
	GetOne(name string) string
	GetAll(name string) []string
}
