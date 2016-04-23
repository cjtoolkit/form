package fields

import (
	"github.com/cjtoolkit/form"
	"strings"
	"time"
)

type Time struct {
	Name     string         // Mandatory
	Label    string         // Mandatory
	Norm     *string        // Mandatory
	Model    *time.Time     // Mandatory
	Err      *error         // Mandatory
	Location *time.Location // Mandatory
	Formats  []string       // Mandatory, most ideal should be at top
	Required bool
	Min      time.Time
	MinZero  bool
	Max      time.Time
	MaxZero  bool
	Extra    func()
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
	*t.Norm = strings.TrimSpace(values.GetOne(t.Name))
}

func (t Time) Transform() {
	*t.Norm = (*t.Model).Format(t.Formats[0])
}

func (t Time) ReverseTransform() {
	for _, format := range t.Formats {
		out, err := time.ParseInLocation(format, *t.Norm, t.Location)
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
			Key: form.LANG_FIELD_REQUIRED,
			Value: map[string]interface{}{
				"Label": t.Label,
			},
		})
	}
}

func (t Time) validateMin() {
	switch {
	case t.Min.IsZero() && t.MinZero:
		return
	case t.Min.Unix() > (*t.Model).Unix() && t.Min.UnixNano() > (*t.Model).UnixNano():
		panic(&form.ErrorValidateModel{
			Key: form.LANG_TIME_MIN,
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
	case t.Max.IsZero() && t.MaxZero:
		return
	case t.Max.Unix() < (*t.Model).Unix() && t.Max.UnixNano() < (*t.Model).UnixNano():
		panic(&form.ErrorValidateModel{
			Key: form.LANG_TIME_MAX,
			Value: map[string]interface{}{
				"Label": t.Label,
				"Max":   t.Max.Format(t.Formats[0]),
			},
		})
	}
}

func (t Time) MaxStr() {
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

func DateFormats() []string{
	// All meets html5 specification.
	return []string{
		"2006-01-02",
	}
}

func TimeFormats() []string{
	// All meets html5 specification.
	return []string{
		"15:04.05",
		"15:04:05",
		"15:04",
	}
}