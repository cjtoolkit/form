package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestInputRadioForm struct {
	Form
	RadioA string `form:"radioA"`
	RadioB string `form:"radioB"`
}

func (t *TestInputRadioForm) RadioAType() string {
	return "input:radio"
}

func (t *TestInputRadioForm) RadioAMandatory() bool {
	return true
}

func (t *TestInputRadioForm) RadioARadio() []Radio {
	return []Radio{
		{"car", "Car", false, nil},
		{"motorbike", "Motorbike", true, nil},
	}
}

func (t *TestInputRadioForm) RadioBType() string {
	return "input:radio"
}

func (t *TestInputRadioForm) RadioBRadio() []Radio {
	return []Radio{
		{"car", "Car", false, nil},
		{"motorbike", "Motorbike", true, nil},
	}
}

type TestNotWellFormedRadioForm struct {
	Form
	RadioA string `name:"radioA"`
}

func (t *TestNotWellFormedRadioForm) RadioAType() string {
	return "input:radio"
}

func TestInputRadio(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestInputRadioForm{
			RadioA: "motorbike",
			RadioB: "motorbike",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputRadioForm{
			RadioA: "car",
			RadioB: "car",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputRadioForm{
			RadioA: "car",
			RadioB: "",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputRadioForm{
			RadioA: "",
			RadioB: "car",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputRadioForm{
			RadioA: "car",
			RadioB: "van",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputRadioForm{
			RadioA: "van",
			RadioB: "car",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}

		// Testing not well formed radio form
		a := &TestNotWellFormedRadioForm{}
		if ValidateItself(a, res, req) == true {
			fmt.Print(RenderString(a))
			t.Fail()
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.Get(ts.URL)
}
