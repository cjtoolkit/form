package form

import (
	"fmt"
	_ "github.com/cjtoolkit/form/lang/enGB"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type inputNumber struct {
	First  int64
	Second float64
}

func (i *inputNumber) CJForm(f *Fields) {

	// First
	func() {
		f := f.Init(&i.First, "First", InputNumber)

		r := f.RangeInt()
		r.Min = 4
		r.Max = 8

		f.StepInt(2)
	}()

	// Second
	func() {
		f := f.Init(&i.Second, "Second", InputNumber)

		r := f.RangeFloat()
		r.Min = 2.65
		r.Max = 7.45

		f.StepFloat(0.5)
	}()
}

func TestInputNumber(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputNumber

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputNumber{}
		check := New(r, nil, "en-GB")

		r.ParseForm()
		b := check.MustValidate(&form)
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
		"First":  {"6"},
		"Second": {"5.5"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// First Above Max
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"10"},
		"Second": {"5.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("First Above Max: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// First Below Min
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"2"},
		"Second": {"5.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("First Below Min: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// First Out of Step
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"7"},
		"Second": {"5.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("First Out of Step: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Second Below Min
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"6"},
		"Second": {"2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Second Below Min: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Second Above Max
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"6"},
		"Second": {"7.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Second Above Max: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Second Out of Step
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"6"},
		"Second": {"5.6"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Second Out of Step: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
