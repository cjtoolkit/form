package form

type FormFieldInterface interface {
	GetErrorPtr() *error
	PopulateNorm(value ValuesInterface)
	Transform()
	ReverseTransform()
	ValidateModel()
}
