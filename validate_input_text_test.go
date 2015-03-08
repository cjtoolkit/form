package form

import (
	"fmt"
	_ "github.com/cjtoolkit/form/lang/enGB"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
)

type inputText struct {
	First  string
	Second string
	Re     string
}

var rePattern = regexp.MustCompile("^[a-z]{1,5}[0-9]{1,5}$")

func (i *inputText) CJForm(f *Fields) {

	// First
	func() {
		f := f.Init("First", InputText)
		f.Mandatory()

		size := f.Size()
		size.Max = 8
	}()

	// Second
	func() {
		f := f.Init("Second", InputText)

		match := f.MustMatch()
		match.Name = "First"
		match.Value = &i.First
	}()

	// Re
	func() {
		f := f.Init("Re", InputText)
		f.Mandatory()
		f.Pattern(rePattern)
	}()
}

func TestInputText(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputText

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputText{}
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
		"First":  {"Hello"},
		"Second": {"Hello"},
		"Re":     {"abcde12345"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Check Matching Rule
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"Hello"},
		"Second": {"World"},
		"Re":     {"abcde12345"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Matching Rule: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Check Size
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"HelloHello"},
		"Second": {"HelloHello"},
		"Re":     {"abcde12345"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Size: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {""},
		"Second": {""},
		"Re":     {"abcde12345"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// RegExp Rule
	res, _ = http.PostForm(ts.URL, url.Values{
		"First":  {"Hello"},
		"Second": {"Hello"},
		"Re":     {"abcdef123456"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("RegExp Rule: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
