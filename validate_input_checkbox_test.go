package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestInputCheckboxForm struct {
	Form
	CheckboxA bool `form:"checkboxA"`
	CheckboxB bool `form:"checkboxB"`
}

func (t *TestInputCheckboxForm) CheckboxAType() string {
	return "input:checkbox"
}

func (t *TestInputCheckboxForm) CheckboxBType() string {
	return "input:checkbox"
}

func (t *TestInputCheckboxForm) CheckboxBMandatory() bool {
	return true
}

func TestInputCheckbox(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestInputCheckboxForm{
			CheckboxA: true,
			CheckboxB: true,
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputCheckboxForm{
			CheckboxA: false,
			CheckboxB: true,
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputCheckboxForm{
			CheckboxA: true,
			CheckboxB: false,
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
