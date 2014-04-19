package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestInputTimeForm struct {
	Form
	TimeA time.Time `form:"timeA"`
	TimeB time.Time `form:"timeB"`
}

func (t *TestInputTimeForm) TimeAType() string {
	return "input:datetime"
}

func (t *TestInputTimeForm) TimeAMin() time.Time {
	return time.Date(2010, 0, 0, 0, 0, 0, 0, time.UTC)
}

func (t *TestInputTimeForm) TimeAMax() time.Time {
	return time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
}

func (t *TestInputTimeForm) TimeBType() string {
	return "input:datetime"
}

func (t *TestInputTimeForm) TimeBMandatory() bool {
	return true
}

func TestInputTime(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestInputTimeForm{
			TimeA: time.Date(2014, 0, 0, 0, 0, 0, 0, time.UTC),
			TimeB: time.Date(1987, 0, 0, 0, 0, 0, 0, time.UTC),
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTimeForm{
			TimeA: time.Date(2014, 0, 0, 0, 0, 0, 0, time.UTC),
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTimeForm{
			TimeA: time.Date(1987, 0, 0, 0, 0, 0, 0, time.UTC),
			TimeB: time.Date(1987, 0, 0, 0, 0, 0, 0, time.UTC),
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputTimeForm{
			TimeA: time.Date(2080, 0, 0, 0, 0, 0, 0, time.UTC),
			TimeB: time.Date(1987, 0, 0, 0, 0, 0, 0, time.UTC),
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
