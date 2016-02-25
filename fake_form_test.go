package form

/*
Implement:
	FormBuilderInterface
*/
type fakeForm struct {
}

func (fF *fakeForm) BuildForm() []FormFieldInterface {
	return nil
}
