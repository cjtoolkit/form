package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestInputTextForm struct {
	Form
	TextA string `form:"textA"`
	TextB string `form:"textB"`
}

func (t *TestInputTextForm) TextAType() string {
	return "input:text"
}

func (t *TestInputTextForm) TextAMinChar() int64 {
	return 2
}

func (t *TestInputTextForm) TextAMaxChar() int64 {
	return 10
}

func (t *TestInputTextForm) TextAPattern() string {
	return "^([a-zA-Z]*)$"
}

func (t *TestInputTextForm) TextAPatternErr() string {
	return "Letters Only"
}

func (t *TestInputTextForm) TextBType() string {
	return "input:text"
}

func (t *TestInputTextForm) TextBMustMatch() string {
	return "TextA"
}

func (t *TestInputTextForm) TextBMustMatchErr() string {
	return "Does not match Test1"
}

func TestInputText(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestInputTextForm{
			TextA: "hello",
			TextB: "hello",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTextForm{
			TextA: "hello5",
			TextB: "hello3",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTextForm{
			TextA: "hellohellohello",
			TextB: "hellohellohello",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTextForm{
			TextA: "a",
			TextB: "a",
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
