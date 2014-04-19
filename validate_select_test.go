package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestSelectForm struct {
	Form
	SelectA string `form:"selectA"`
	SelectB string `form:"selectB"`
}

func (t *TestSelectForm) SelectAType() string {
	return "select"
}

func (t *TestSelectForm) SelectAMandatory() bool {
	return true
}

func (t *TestSelectForm) SelectAOptions() []Option {
	return []Option{
		{"Car", "car", "", false, nil},
		{"Motorbike", "motorbike", "", true, nil},
	}
}

func (t *TestSelectForm) SelectBType() string {
	return "select"
}

func (t *TestSelectForm) SelectBOptions() []Option {
	return []Option{
		{"Car", "car", "", false, nil},
		{"Motorbike", "motorbike", "", true, nil},
	}
}

type TestNotWellFormedSelectForm struct {
	Form
	SelectA string `name:"selectA"`
}

func (t *TestNotWellFormedSelectForm) SelectAType() string {
	return "input:radio"
}

func TestSelect(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestSelectForm{
			SelectA: "motorbike",
			SelectB: "motorbike",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestSelectForm{
			SelectA: "car",
			SelectB: "car",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestSelectForm{
			SelectA: "car",
			SelectB: "",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestSelectForm{
			SelectA: "",
			SelectB: "car",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestSelectForm{
			SelectA: "car",
			SelectB: "van",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestSelectForm{
			SelectA: "van",
			SelectB: "car",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}

		// Testing not well formed radio form
		a := &TestNotWellFormedSelectForm{}
		if ValidateItself(a, res, req) == true {
			fmt.Print(RenderString(a))
			t.Fail()
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.Get(ts.URL)
}
