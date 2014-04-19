package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestTextAreaForm struct {
	Form
	Textarea string `form:"textarea"`
}

func (t *TestTextAreaForm) TextareaType() string {
	return "textarea"
}

func (t *TestTextAreaForm) TextareaMinChar() int64 {
	return 4
}

func (t *TestTextAreaForm) TextareaMaxChar() int64 {
	return 10
}

func TestTextarea(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestTextAreaForm{
			Textarea: "hello",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestTextAreaForm{
			Textarea: "a",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestTextAreaForm{
			Textarea: "hello world",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.Get(ts.URL)
}
