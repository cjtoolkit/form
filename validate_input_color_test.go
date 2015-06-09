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

type inputColor struct {
	First string
}

func (i *inputColor) CJForm(f *Fields) {

	// First
	func() {
		f.Init(&i.First, "First", InputColor)
	}()
}

func TestInputColor(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputColor

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputColor{}
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
		"First": {"#A0A0A0"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Check Color Rule #1
	res, _ = http.PostForm(ts.URL, url.Values{
		"First": {"#AGAGAG"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Color Rule #1: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
