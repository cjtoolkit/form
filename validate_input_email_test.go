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

type inputEmail struct {
	First  string
	Second string
}

func (i *inputEmail) CJForm(f Fields) {

	// First
	func() {
		f := f.Init("First", InputEmail)
		f.Mandatory()
	}()

	// Second
	func() {
		f := f.Init("Second", InputEmail)

		match := f.MustMatch()
		match.Name = "First"
		match.Value = &i.First
	}()
}

func TestInputEmail(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputEmail

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputEmail{}
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
		"First":  {"hello@example.com"},
		"Second": {"hello@example.com"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Valid Email Address
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"hello_example.com"},
		"Second": {"hello_example.com"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Matching Rule: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Check Matching Rule
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"hello@example.com"},
		"Second": {"world@example.com"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Matching Rule: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {""},
		"Second": {""},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
