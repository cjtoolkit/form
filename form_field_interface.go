package form

type FormFieldInterface interface {
	PreCheck()
	GetErrorPtr() *error
	PopulateNorm(values ValuesInterface)
	Transform()
	ReverseTransform()
	ValidateModel()
}
