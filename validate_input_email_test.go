package form

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var exampleMail = regexp.MustCompile(`example.com$`)

type TestInputEmailForm struct {
	Form
	EmailA string `form:"textA"`
	EmailB string `form:"textB"`
}

func (t *TestInputEmailForm) EmailAType() string {
	return "input:email"
}

func (t *TestInputEmailForm) EmailAExt() {
	if exampleMail.MatchString(t.EmailA) {
		t.SetErr("EmailA", FormError("We do not accept example.com email address, sorry!"))
	}
}

func (t *TestInputEmailForm) EmailBType() string {
	return "input:email"
}

func (t *TestInputEmailForm) EmailBMustMatch() string {
	return "EmailA"
}

func (t *TestInputEmailForm) EmailBMustMatchErr() string {
	return "Does not match EmailA"
}

func TestInputEmail(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		s := &TestInputEmailForm{
			EmailA: "hello@cj-jackson.com",
			EmailB: "hello@cj-jackson.com",
		}
		if ValidateItself(s, res, req) == false {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputEmailForm{
			EmailA: "hello@cj-jackson.com",
			EmailB: "support@cj-jackson.com",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputEmailForm{
			EmailA: "hello_cj-jackson.com",
			EmailB: "hello_cj-jackson.com",
		}
		if ValidateItself(s, res, req) == true {
			fmt.Print(RenderString(s))
			t.Fail()
		}
		s = &TestInputEmailForm{
			EmailA: "hello@example.com",
			EmailB: "hello@example.com",
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
