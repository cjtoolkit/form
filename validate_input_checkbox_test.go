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

type inputCheckbox struct {
	First bool
}

func (i *inputCheckbox) CJForm(f Fields) {

	// First
	func() {
		f := f.Init("First", InputCheckbox)
		f["mandatory"] = func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		}
	}()
}

func TestInputCheckbox(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputCheckbox

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputCheckbox{}
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
		"First": {"1"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Check Checkbox Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"First": {""},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Check Checkbox Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
