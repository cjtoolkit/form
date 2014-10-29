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

type select_ struct {
	Str  string
	Strs []string
	W    int64
	Ws   []int64
	F    float64
	Fs   []float64
}

func (i *select_) StrField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]Option)) = []Option{
				{Value: "Hello", Label: "Hello"},
				{Value: "World", Label: "World"},
			}
		},
	}
}

func (i *select_) StrsField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]Option)) = []Option{
				{Value: "Hello", Label: "Hello"},
				{Value: "World", Label: "World"},
			}
		},
	}
}

func (i *select_) WField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]OptionInt)) = []OptionInt{
				{Value: 1, Label: "Hello"},
				{Value: 2, Label: "World"},
			}
		},
	}
}

func (i *select_) WsField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]OptionInt)) = []OptionInt{
				{Value: 1, Label: "Hello"},
				{Value: 2, Label: "World"},
			}
		},
	}
}

func (i *select_) FField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]OptionFloat)) = []OptionFloat{
				{Value: 1.5, Label: "Hello"},
				{Value: 2.5, Label: "World"},
			}
		},
	}
}

func (i *select_) FsField() FieldFuncs {
	return FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*TypeCode)) = Select
		},
		"mandatory": func(m map[string]interface{}) {
			*(m["mandatory"].(*bool)) = true
		},
		"option": func(m map[string]interface{}) {
			*(m["option"].(*[]OptionFloat)) = []OptionFloat{
				{Value: 1.5, Label: "Hello"},
				{Value: 2.5, Label: "World"},
			}
		},
	}
}

func TestSelect(t *testing.T) {
	mux := http.NewServeMux()

	var outform select_

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form := select_{}
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
		"Str":  {"Hello"},
		"Strs": {"Hello", "World"},
		"W":    {"1"},
		"Ws":   {"1", "2"},
		"F":    {"1.5"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ := ioutil.ReadAll(res.Body)

	if string(b) != "true" {
		t.Errorf("Init: Expected 'true', return %s. \r\n %v", b, outform)
	}

	// Str Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {""},
		"Strs": {"Hello", "World"},
		"W":    {"1"},
		"Ws":   {"1", "2"},
		"F":    {"1.5"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Str Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Strs Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {"Hello"},
		"Strs": {""},
		"W":    {"1"},
		"Ws":   {"1", "2"},
		"F":    {"1.5"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Strs Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// W Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {"Hello"},
		"Strs": {"Hello", "World"},
		"W":    {"0"},
		"Ws":   {"1", "2"},
		"F":    {"1.5"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("W Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Ws Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {"Hello"},
		"Strs": {"Hello", "World"},
		"W":    {"1"},
		"Ws":   {"0"},
		"F":    {"1.5"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Ws Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// F Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {"Hello"},
		"Strs": {"Hello", "World"},
		"W":    {"1"},
		"Ws":   {"1", "2"},
		"F":    {"0"},
		"Fs":   {"1.5", "2.5"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("F Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}

	// Fs Mandatory
	res, _ = http.PostForm(ts.URL, url.Values{
		"Str":  {"Hello"},
		"Strs": {"Hello", "World"},
		"W":    {"1"},
		"Ws":   {"1", "2"},
		"F":    {"1"},
		"Fs":   {"0"},
	})

	b, _ = ioutil.ReadAll(res.Body)

	if string(b) != "false" {
		t.Errorf("Fs Mandatory: Expected 'false', return %s. \r\n %v", b, outform)
	}
}
