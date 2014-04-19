package form

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

type TestPopulatorForm struct {
	Form
	Text    string    `form:"text"`
	Number  int64     `form:"number"`
	FNumber float64   `form:"fnumber"`
	Logic   bool      `form:"logic"`
	TimeA   time.Time `form:"timeA"`
	TimeB   time.Time `form:"timeB"`
}

func (t *TestPopulatorForm) TextType() string {
	return "input:text"
}

func (t *TestPopulatorForm) NumberType() string {
	return "input:number"
}

func (t *TestPopulatorForm) FNumberType() string {
	return "input:number"
}

func (t *TestPopulatorForm) LogicType() string {
	return "input:checkbox"
}

func (t *TestPopulatorForm) TimeAType() string {
	return "input:datetime"
}

func (t *TestPopulatorForm) TimeBType() string {
	return "input:week"
}

func TestPopulator(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		form := &TestPopulatorForm{}
		Validate(form, res, req)

		if form.Text != "Hello" {
			t.Fail()
		}

		if form.Number != 3 {
			t.Fail()
		}

		if form.FNumber != 3.5 {
			t.Fail()
		}

		if form.Logic != true {
			t.Fail()
		}

		if form.TimeA.Unix() != time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC).Unix() {
			t.Fail()
		}

		if _, week := form.TimeB.ISOWeek(); week != 12 {
			t.Fail()
		}

		form = &TestPopulatorForm{}
		Validate(form, res, req)

		if form.Text != "World" {
			t.Fail()
		}

		if form.Number != 6 {
			t.Fail()
		}

		if form.FNumber != 6.5 {
			t.Fail()
		}

		if form.Logic != false {
			t.Fail()
		}

		if form.TimeA.Unix() != time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).Unix() {
			t.Fail()
		}

		if _, week := form.TimeB.ISOWeek(); week != 52 {
			t.Fail()
		}

		form = &TestPopulatorForm{}
		Validate(form, res, req)

		if form.Text != "!" {
			t.Fail()
		}

		if form.Number != 9 {
			t.Fail()
		}

		if form.FNumber != 9.5 {
			t.Fail()
		}

		if form.Logic != true {
			t.Fail()
		}

		if form.TimeA.Unix() != time.Date(2030, 0, 0, 0, 0, 0, 0, time.UTC).Unix() {
			t.Fail()
		}

		if _, week := form.TimeB.ISOWeek(); week != 53 {
			t.Fail()
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.PostForm(ts.URL, url.Values{
		"text":    {"Hello", "World", "!"},
		"number":  {"3", "6", "9"},
		"fnumber": {"3.5", "6.5", "9.5"},
		"logic":   {"1", "", "1"},
		"timeA":   {time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC).Format(dateTimeFormat), time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC).Format(dateTimeFormat), time.Date(2030, 0, 0, 0, 0, 0, 0, time.UTC).Format(dateTimeFormat)},
		"timeB":   {"2014-W12", "2014-W53", "2015-W53"},
	})
}
