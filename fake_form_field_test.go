package form

/*
Implement:
	FormFieldInterface
*/
type fakeFormField struct{}

func (fFF *fakeFormField) GetErrorPtr() *error {
	return nil
}

func (fFF *fakeFormField) PopulateNorm(values ValuesInterface) {

}

func (fFF *fakeFormField) Transform() {

}

func (fFF *fakeFormField) ReverseTransform() {

}

func (fFF *fakeFormField) ValidateModel() {

}
