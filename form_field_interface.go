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
	GetBridge() *FieldBridge
}

func FormFieldInterfaceCheck(v FormFieldInterface) {}
