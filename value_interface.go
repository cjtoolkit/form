package form

type ValueInterface interface {
	GetOne(name string) string
	GetAll(name string) []string
}
