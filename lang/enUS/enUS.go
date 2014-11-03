// English (American)
package enUS

import (
	"github.com/cjtoolkit/i18n"
	_ "github.com/cjtoolkit/i18n/lang/enUS"
)

func init() {
	i18n.AppendMap("en-US", "cjtoolkit-form", map[string]interface{}{
		"ErrMandatory":         "This field is mandatory!",
		"ErrMinChar":           "Characters must be more than and equal to '{{.Count}}'!",
		"ErrMaxChar":           "Characters must be less than and equal to '{{.Count}}'!",
		"ErrMustMatchMissing":  "Must Match is missing!",
		"ErrMandatoryCheckbox": "This checkbox requires marking!",
		"ErrMimeCheck": i18n.Plural{
			One:   "Invalid file mime, valid mime is : {{.Mimes}}",
			Other: "Invalid file mime, valid mimes are : {{.Mimes}}",
		},
		"ErrSelectValueMissing":     "Select Value is missing!",
		"ErrSelectOptionIsRequired": "Option is required!",
		"ErrFieldDoesNotExist":      "Field name does not exist!",
		"ErrInvalidEmailAddress":    "Invalid Email Address!",
		"ErrFileRequired":           "File is required!",
		"ErrNumberMin":              "Number must be minimum of '{{.Count}}' or above!",
		"ErrNumberMax":              "Number must be maximum of '{{.Count}}' or below!",
		"ErrNumberStep":             "Number must be in step of '{{.Count}}'!",
		"ErrPatternMismatch":        "Value must match pattern of '{{.Pattern}}'!",
		"ErrMustMatchMismatch":      "Value must match '{{.Name}}'!",
		"ErrRadioNotWellFormed":     "Radio validator is not well formed",
		"ErrSelectNotWellFormed":    "Select validator is not well formed",
		"ErrOutOfBound":             "Out of bound!",
		"ErrNotSelect":              "Not selected!",
		"ErrFileSize":               "Filesize cannot exceed '{{.Size}}' in bytes",
		"ErrType":                   "Invalid Form type for Data Type `{{.DataType}}`!",
		"ErrTimeMin":                "Time cannot be less than '{{.Time}}'",
		"ErrTimeMax":                "Time cannot be greater than '{{.Time}}'",
		"ErrInvalidColorCode":       "Invalid color code",
	})
}
