package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestInputNumberForm struct {
	Form
	Number1 int64 `form:"number1"`
	Number2 int64 `form:"number2"`
}

func (t *TestInputNumberForm) Number1Type() string {
	return "input:number"
}

func (t *TestInputNumberForm) Number1Min() int64 {
	return 5
}

func (t *TestInputNumberForm) Number1Max() int64 {
	return 15
}

func (t *TestInputNumberForm) Number2Type() string {
	return "input:range"
}

func (t *TestInputNumberForm) Number2Step() int64 {
	return 5
}

type TestInputNumberFloatForm struct {
	Form
	Number1 float64 `form:"number1"`
	Number2 float64 `form:"number2"`
}

func (t *TestInputNumberFloatForm) Number1Type() string {
	return "input:number"
}

func (t *TestInputNumberFloatForm) Number1Min() float64 {
	return 2.3
}

func (t *TestInputNumberFloatForm) Number1Max() float64 {
	return 17.5
}

func (t *TestInputNumberFloatForm) Number2Type() string {
	return "input:range"
}

func (t *TestInputNumberFloatForm) Number2Step() float64 {
	return 0.5
}

func TestInputNumber(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// Whole Number
		s := &TestInputNumberForm{
			Number1: 10,
			Number2: 10,
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputNumberForm{
			Number1: 3,
			Number2: 10,
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputNumberForm{
			Number1: 18,
			Number2: 10,
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputNumberForm{
			Number1: 10,
			Number2: 8,
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}

		// Floating Number
		ff := &TestInputNumberFloatForm{
			Number1: 5.3,
			Number2: 2.5,
		}
		if ValidateItself(ff, res, req) == false {
			fmt.Print(RenderString(ff))
			t.Fail()
		}
		ff = &TestInputNumberFloatForm{
			Number1: 2.2,
			Number2: 2.5,
		}
		if ValidateItself(ff, res, req) == true {
			fmt.Print(RenderString(ff))
			t.Fail()
		}
		ff = &TestInputNumberFloatForm{
			Number1: 17.6,
			Number2: 2.5,
		}
		if ValidateItself(ff, res, req) == true {
			fmt.Print(RenderString(ff))
			t.Fail()
		}
		ff = &TestInputNumberFloatForm{
			Number1: 5.3,
			Number2: 2.4,
		}
		if ValidateItself(ff, res, req) == true {
			fmt.Print(RenderString(ff))
			t.Fail()
		}
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	http.Get(ts.URL)
}
