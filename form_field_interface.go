package form

type FormFieldInterface interface {
	GetErrorPtr() *error
	PopulateNorm(value ValueInterface)
	Transform()
	ReverseTransform()
	ValidateModel()
}
