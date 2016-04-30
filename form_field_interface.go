package form

import "encoding/json"

type FormFieldInterface interface {
	json.Marshaler
	PreCheck()
	GetErrorPtr() *error
	PopulateNorm(values ValuesInterface)
	Transform()
	ReverseTransform()
	ValidateModel()
}

func FormFieldInterfaceCheck(v FormFieldInterface) {}
