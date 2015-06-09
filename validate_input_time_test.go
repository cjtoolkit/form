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

func (i *inputTime) CJForm(f *Fields) {

	// Datetime
	func() {
		f := f.Init(&i.Datetime, "Datetime", InputDatetime)
		r := f.RangeTime()
		r.Min = time.Date(2014, 2, 15, 5, 0, 1, 0, time.UTC)
		r.Max = time.Date(2014, 10, 15, 4, 59, 59, 0, time.UTC)
	}()

	// Datetimelocal
	func() {
		f.Init(&i.Datetimelocal, "Datetimelocal", InputDatetimeLocal)
	}()

	// Date
	func() {
		f.Init(&i.Date, "Date", InputDate)
	}()

	// Time
	func() {
		f.Init(&i.Time, "Time", InputTime)
	}()

	// Month
	func() {
		f.Init(&i.Month, "Month", InputMonth)
	}()

	// Week
	func() {
		f.Init(&i.Week, "Week", InputWeek)
	}()

}

func TestInputTime(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputTime

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputTime{}
		check := New(r, nil, "en-GB")

		r.ParseForm()
		b := check.MustValidate(&form)
		outform = form
		if b {
			fmt.Fprint(w, "true")
		} else {
			fmt.Fprint(w, "false")
		}

		check.RenderStr(&form)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Init
	res, _ := http.PostForm(ts.URL, url.Values{
		"Datetime":      {"2014-06-15T05:00Z"},
		"Datetimelocal": {"2014-06-15T05:00"},
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
		"Datetime":      {"2014-02-15T05:00Z"},
		"Datetimelocal": {"2014-06-15T05:00"},
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
		"Datetime":      {"2014-10-15T05:00Z"},
		"Datetimelocal": {"2014-06-15T05:00"},
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
