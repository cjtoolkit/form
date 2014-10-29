package form

import (
	"fmt"
	_ "github.com/cjtoolkit/form/lang/enGB"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

type inputTime struct {
	Datetime      time.Time
	Datetimelocal time.Time
	Date          time.Time
	Time          time.Time
	Month         time.Time
	Week          time.Time
}

func (i *inputTime) DatetimeField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputDatetime
		},
		"range": func(m map[string]interface{}) {
			// time.Date(year, month, day, hour, min, sec, nsec, loc)
			*(m["min"].(*time.Time)) = time.Date(2014, 2, 15, 5, 0, 1, 0, time.UTC)
			*(m["max"].(*time.Time)) = time.Date(2014, 10, 15, 4, 59, 59, 0, time.UTC)
		},
	}
}

func (i *inputTime) DatetimelocalField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputDatetimeLocal
		},
	}
}

func (i *inputTime) DateField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputDate
		},
	}
}

func (i *inputTime) TimeField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputTime
		},
	}
}

func (i *inputTime) MonthField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputMonth
		},
	}
}

func (i *inputTime) WeekField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputWeek
		},
	}
}

func TestInputTime(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputTime

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputTime{}
		check := New(nil, "en-GB")

		r.ParseForm()
		b := check.MustValidate(r, &form)
		outform = form
		if b {
			fmt.Fprint(w, "true")
		} else {
			fmt.Fprint(w, "false")
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Init
	res, _ := http.PostForm(ts.URL, url.Values{
		"Datetime":      {"2014-06-15T05:00:00Z"},
		"Datetimelocal": {"2014-06-15T05:00:00"},
		"Date":          {"2014-06-15"},
		"Time":          {"05:00:00"},
		"Month":         {"2014-06"},
		"Week":          {"2014-W02"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Populate

	if outform.Datetime.Unix() == -62135596800 {
		t.Errorf("Populate: 'Datetime' \r\n %v", outform)
	}

	if outform.Datetimelocal.Unix() == -62135596800 {
		t.Errorf("Populate: 'Datetimelocal' \r\n %v", outform)
	}

	if outform.Date.Unix() == -62135596800 {
		t.Errorf("Populate: 'Date' \r\n %v", outform)
	}

	if outform.Time.Unix() == -62135596800 {
		t.Errorf("Populate: 'Time' \r\n %v", outform)
	}

	if outform.Month.Unix() == -62135596800 {
		t.Errorf("Populate: 'Month' \r\n %v", outform)
	}

	if outform.Week.Unix() == -62135596800 {
		t.Errorf("Populate: 'Week' \r\n %v", outform)
	}

	// Below Min
	res, _ = http.PostForm(ts.URL, url.Values{
		"Datetime":      {"2014-02-15T05:00:00Z"},
		"Datetimelocal": {"2014-06-15T05:00:00"},
		"Date":          {"2014-06-15"},
		"Time":          {"05:00:00"},
		"Month":         {"2014-06"},
		"Week":          {"2014-W02"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Below Min: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Above Max
	res, _ = http.PostForm(ts.URL, url.Values{
		"Datetime":      {"2014-10-15T05:00:00Z"},
		"Datetimelocal": {"2014-06-15T05:00:00"},
		"Date":          {"2014-06-15"},
		"Time":          {"05:00:00"},
		"Month":         {"2014-06"},
		"Week":          {"2014-W02"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Above Max: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
