package form

const (
	ErrMandatory              = "This field is mandatory! (Also first character must not be space or new line)"
	ErrMinChar                = "Characters must be more than and equal to '%v'!"
	ErrMaxChar                = "Characters must be less than and equal to '%v'!"
	ErrMustMatchMissing       = "Must Match is missing!"
	ErrMandatoryCheckbox      = "This checkbox requires marking!"
	ErrMimeCheck              = "Invalid file mime, valid mime are/is : "
	ErrSelectValueMissing     = "Select Value is missing!"
	ErrSelectOptionIsRequired = "Option is required!"
	ErrFieldDoesNotExist      = "Field name does not exist!"
	ErrInvalidEmailAddress    = "Invalid Email Address!"
	ErrFileRequired           = "File is required!"
	ErrNumberMin              = "Number must be minimum of '%v' or above!"
	ErrNumberMax              = "Number must be maximum of '%v' or below!"
	ErrNumberStep             = "Number must be in step of '%v'!"
	ErrPatternMismatch        = "Value must match pattern of '%v'!"
	ErrMustMatchMismatch      = "Value must match '%v'!"
	ErrRadioNotWellFormed     = "Radio validator is not well formed"
	ErrSelectNotWellFormed    = "Select validator is not well formed"
	ErrOutOfBound             = "Out of bound!"
	ErrNotSelect              = "Not selected!"
	ErrFileSize               = "Filesize cannot exceed '%v' in bytes"
	ErrType                   = "Form type '{{.FormType}}' is not a valid for data type `{{.DataType}}`!"
	ErrTimeMin                = "Time cannot be less than '%v'"
	ErrTimeMax                = "Time cannot be greater than '%v'"
	ErrInvalidColorCode       = "Invalid color code"
)

type i18nDummy int

func (_ i18nDummy) Key(name string) string {
	return name
}

type I18n interface {
	Key(string) string
}

var DefaultI18n I18n = i18nDummy(0)
