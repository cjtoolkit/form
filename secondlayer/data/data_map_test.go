package data

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cjtoolkit/form"
	_ "github.com/cjtoolkit/form/lang/enGB"
)

type dataMapForm struct {
	First  string
	Second string
}

func (i *dataMapForm) FirstField() form.FieldFuncs {
	return form.FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*form.TypeCode)) = form.InputText
		},
	}
}

func (i *dataMapForm) SecondField() form.FieldFuncs {
	return form.FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*form.TypeCode)) = form.InputText
		},
	}
}

func TestDataMap(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f := dataMapForm{}
		dataCollector := NewDataMap()
		check := form.New(dataCollector, "en-GB")
		check.RenderStr(&f)
		data := dataCollector.PullAndReset()

		if data["First"][0].Name != "First" {
			t.Errorf(`Expected "First", got "%s"`, data["First"][0].Name)
		}

		if data["Second"][0].Name != "Second" {
			t.Errorf(`Expected "Second", got "%s"`, data["Second"][0].Name)
		}

		w.Write([]byte("Hello"))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.PostForm(ts.URL, url.Values{
		"First":  {"Hello"},
		"Second": {"World"},
	})
}
