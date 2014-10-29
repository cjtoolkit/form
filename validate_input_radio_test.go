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

type inputRadio struct {
	Str string
	W   int64
	F   float64
}

func (i *inputRadio) StrField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]Radio)) = []Radio{
				{Value: "Hello", Label: "Hello"},
				{Value: "World", Label: "World"},
			}
		},
	}
}

func (i *inputRadio) WField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]RadioInt)) = []RadioInt{
				{Value: 1, Label: "Hello"},
				{Value: 2, Label: "World"},
			}
		},
	}
}

func (i *inputRadio) FField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = InputRadio
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"radio": func(m map[string]interface{}) {
			*(m["radio"].(*[]RadioFloat)) = []RadioFloat{
				{Value: 1.5, Label: "Hello"},
				{Value: 2.5, Label: "World"},
			}
		},
	}
}

func TestInputRadio(t *testing.T) {
	mux := http.NewServeMux()

	var outform inputRadio

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := inputRadio{}
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
		"Str": {"Hello"},
		"W":   {"1"},
		"F":   {"1.5"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Str Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str": {""},
		"W":   {"1"},
		"F":   {"1.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Str Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// W Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str": {"Hello"},
		"W":   {"0"},
		"F":   {"1.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("W Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// F Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str": {"Hello"},
		"W":   {"1"},
		"F":   {"0"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("F Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
