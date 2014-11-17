package data

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cjtoolkit/form"
	_ "github.com/cjtoolkit/form/lang/enGB"
)

type dataSliceForm struct {
	First  string
	Second string
}

func (i *dataSliceForm) FirstField() form.FieldFuncs {
	return form.FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*form.TypeCode)) = form.InputText
		},
	}
}

func (i *dataSliceForm) SecondField() form.FieldFuncs {
	return form.FieldFuncs{
		"form": func(m map[string]interface{}) {
			*(m["type"].(*form.TypeCode)) = form.InputText
		},
	}
}

func TestDataSlice(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f := dataSliceForm{}
		dataCollector := NewDataSlice()
		check := form.New(dataCollector, "en-GB")
		check.RenderStr(&f)
		data := dataCollector.PullAndReset()

		if data[0].Name != "First" {
			t.Errorf(`Expected "First", got "%s"`, data[0].Name)
		}

		if data[1].Name != "Second" {
			t.Errorf(`Expected "Second", got "%s"`, data[1].Name)
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
