package form

type FormFieldInterface interface {
	GetErrorPtr() *error
	PopulateNorm(values ValuesInterface)
	Transform()
	ReverseTransform()
	ValidateModel()
}
