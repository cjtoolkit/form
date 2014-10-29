package form

// For expressing form input type, eg. <input type="text"... , <textarea....
type TypeCode uint8

// List of avaliable form input type, 'Invalid' does not count!
const (
	Invalid TypeCode = iota
	InputCheckbox
	InputColor
	InputDate
	InputDatetime
	InputDatetimeLocal
	InputEmail
	InputFile
	InputHidden
	InputMonth
	InputNumber
	InputPassword
	InputRadio
	InputRange
	InputSearch
	InputTel
	InputText
	InputTime
	InputUrl
	InputWeek
	Select
	Textarea
	terminate
)
