package form

/*
Implement:
	FormBuilderInterface
*/
type fakeForm struct {
}

func (fF *fakeForm) Fields() []FormFieldInterface {
	return nil
}
