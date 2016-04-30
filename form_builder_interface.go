package form

type FormBuilderInterface interface {
	Fields() []FormFieldInterface
}

func FormBuilderInterfaceCheck(v FormBuilderInterface) {

}
