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

func (i *select_) CJForm(f Fields) {

	// Str
	func() {
		f := f.Init("Str", Select)
		f.Mandatory()
		f.Options([]Option{
			{Value: "Hello", Label: "Hello"},
			{Value: "World", Label: "World"},
		})
	}()

	// Strs
	func() {
		f := f.Init("Strs", Select)
		f.Mandatory()
		f.Options([]Option{
			{Value: "Hello", Label: "Hello"},
			{Value: "World", Label: "World"},
		})
	}()

	// W
	func() {
		f := f.Init("W", Select)
		f.Mandatory()
		f.Options([]OptionInt{
			{Value: 1, Label: "Hello"},
			{Value: 2, Label: "World"},
		})
	}()

	// Ws
	func() {
		f := f.Init("Ws", Select)
		f.Mandatory()
		f.Options([]OptionInt{
			{Value: 1, Label: "Hello"},
			{Value: 2, Label: "World"},
		})
	}()

	// F
	func() {
		f := f.Init("F", Select)
		f.Mandatory()
		f.Options([]OptionFloat{
			{Value: 1.5, Label: "Hello"},
			{Value: 2.5, Label: "World"},
		})
	}()

	// Fs
	func() {
		f := f.Init("Fs", Select)
		f.Mandatory()
		f.Options([]OptionFloat{
			{Value: 1.5, Label: "Hello"},
			{Value: 2.5, Label: "World"},
		})
	}()

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
