package fields

import (
	"encoding/json"
	"github.com/cjtoolkit/form"
	"strings"
	"time"
)

type Time struct {
	Name           string         // Mandatory
	Label          string         // Mandatory
	Norm           *string        // Mandatory
	Model          *time.Time     // Mandatory
	Err            *error         // Mandatory
	Location       *time.Location // Mandatory
	Formats        []string       // Mandatory, most ideal format should be at the top
	Suffix         *string
	Required       bool
	RequiredErrKey string
	Min            time.Time
	MinZero        bool
	MinErrKey      string
	Max            time.Time
	MaxZero        bool
	MaxErrKey      string
	Extra          func()
}

func NewTime(
	name, label string,
	norm *string,
	model *time.Time,
	err *error,
	location *time.Location,
	formats []string,
	options ...func(*Time),
) Time {
	t := Time{
		Name:     name,
		Label:    label,
		Norm:     norm,
		Model:    model,
		Err:      err,
		Location: location,
		Formats:  formats,
	}

	t.PreCheck()

	for _, option := range options {
		option(&t)
	}

	return t
}

type timeJson struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Success  bool   `json:"success"`
	Error    string `json:"error,omitempty"`
	Min      string `json:"min,omitempty"`
	MinUnix  int64  `json:"minUnix,omitempty"`
	Max      string `json:"max,omitempty"`
	MaxUnix  int64  `json:"maxUnix"`
}

func (t Time) NameWithSuffix() string {
	return addSuffix(t.Name, t.Suffix)
}

func (t Time) timeToString(tt time.Time, zero bool) string {
	if tt.IsZero() && !zero {
		return ""
	}
	return tt.Format(t.Formats[0])
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(timeJson{
		Type:     "time",
		Name:     t.Name,
		Required: t.Required,
		Success:  nil == *t.Err,
		Error:    getMessageFromError(*t.Err),
		Min:      t.timeToString(t.Min, t.MinZero),
		Max:      t.timeToString(t.Max, t.MaxZero),
		MinUnix:  t.Min.Unix(),
		MaxUnix:  t.Max.Unix(),
	})
}

func (t Time) PreCheck() {
	switch {
	case "" == strings.TrimSpace(t.Name):
		panic(form.ErrorPreCheck("Time Field: Name cannot be empty string"))
	case "" == strings.TrimSpace(t.Label):
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Label cannot be empty string"))
	case nil == t.Norm:
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Norm cannot be nil value"))
	case nil == t.Model:
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Model cannot be nil value"))
	case nil == t.Err:
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Err cannot be nil value"))
	case nil == t.Location:
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Location cannot be nil value"))
	case 0 == len(t.Formats):
		panic(form.ErrorPreCheck("Time Field: " + t.Name + ": Formats cannot be nil value or empty"))
	}
}

func (t Time) GetErrorPtr() *error {
	return t.Err
}

func (t Time) PopulateNorm(values form.ValuesInterface) {
	*t.Norm = values.GetOne(t.NameWithSuffix())
}

func (t Time) Transform() {
	*t.Norm = (*t.Model).Format(t.Formats[0])
}

func (t Time) ReverseTransform() {
	norm := strings.TrimSpace(*t.Norm)
	*t.Model = time.Time{}
	for _, format := range t.Formats {
		out, err := time.ParseInLocation(format, norm, t.Location)
		if nil != err {
			continue
		}
		*t.Model = out
		return
	}

	panic(form.ErrorReverseTransform{
		Key: form.LANG_TIME_FORMAT,
		Value: map[string]interface{}{
			"Label": t.Label,
		},
	})
}

func (t Time) ValidateModel() {
	t.validateRequired()
	t.validateMin()
	t.validateMax()
	execFnIfNotNil(t.Extra)
}

func (t Time) validateRequired() {
	switch {
	case !t.Required:
		return
	case (*t.Model).IsZero():
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_FIELD_REQUIRED, t.RequiredErrKey),
			Value: map[string]interface{}{
				"Label": t.Label,
			},
		})
	}
}

func (t Time) validateMin() {
	switch {
	case t.Min.IsZero() && !t.MinZero:
		return
	case t.Min.Unix() > (*t.Model).Unix() || (t.Min.Unix() == (*t.Model).Unix() && t.Min.UnixNano() > (*t.Model).UnixNano()):
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_TIME_MIN, t.MinErrKey),
			Value: map[string]interface{}{
				"Label": t.Label,
				"Min":   t.Min.Format(t.Formats[0]),
			},
		})
	}
}

func (t Time) MinStr() string {
	return t.Min.Format(t.Formats[0])
}

func (t Time) validateMax() {
	switch {
	case t.Max.IsZero() && !t.MaxZero:
		return
	case t.Max.Unix() < (*t.Model).Unix() || (t.Max.Unix() == (*t.Model).Unix() && t.Max.UnixNano() < (*t.Model).UnixNano()):
		panic(&form.ErrorValidateModel{
			Key: UseDefaultKeyIfCustomKeyIsEmpty(form.LANG_TIME_MAX, t.MaxErrKey),
			Value: map[string]interface{}{
				"Label": t.Label,
				"Max":   t.Max.Format(t.Formats[0]),
			},
		})
	}
}

func (t Time) MaxStr() string {
	return t.Max.Format(t.Formats[0])
}

func DateTimeLocalFormats() []string {
	// All meets html5 specification.
	return []string{
		"2006-01-02T15:04.05",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02",
		"15:04.05",
		"15:04:05",
		"15:04",
	}
}

func DateFormats() []string {
	// Meets html5 specification.
	return []string{
		"2006-01-02",
	}
}

func TimeFormats() []string {
	// All meets html5 specification.
	return []string{
		"15:04.05",
		"15:04:05",
		"15:04",
	}
}
