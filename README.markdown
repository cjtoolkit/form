# CJToolkit Form

Highly Flexible Form Rendering and Validation System.

Generic Port of [Gorail Form](https://github.com/gorail/form).  Made to be directly compatible with 'net/http' or any framework compatible with 'net/http' such as [Martini](https://github.com/go-martini/martini) or [Gorail Core](https://github.com/gorail/core).

Documentation can be found at.

http://gowalker.org/github.com/cjtoolkit/form

## Installation

~~~
go get github.com/cjtoolkit/form
~~~

## Example

~~~ go
package main

import (
	"fmt"
	"github.com/cjtoolkit/form"
	"github.com/go-martini/martini"
	"net/http"
)

type ContactNameForm struct {
	form.Form
	Title     string `form:"title"`
	FirstName string `form:"first-name"`
	LastName  string `form:"last-name"`
}

func (co *ContactNameForm) TitleType() string {
	return "select"
}

func (co *ContactNameForm) TitleAttr() map[string]string {
	return map[string]string{"class": "form-control", "id": "contact-title"}
}

func (co *ContactNameForm) TitleLabel() form.Label {
	return form.Label{"Title:", "contact-title", map[string]string{"class": "control-label"}}
}

// Mandatory for 'select'.
func (co *ContactNameForm) TitleOptions() []form.Option {
	return []form.Option{
		{
			Content:  "Mr.",
			Value:    "Mr.",
			Label:    "Mr.",
			Selected: true,
		},
		{
			Content: "Mrs.",
			Value:   "Mrs.",
			Label:   "Mrs.",
		},
		{
			Content: "Miss.",
			Value:   "Miss.",
			Label:   "Miss.",
		},
		{
			Content: "Ms.",
			Value:   "Ms.",
			Label:   "Ms.",
		},
		{
			Content: "Dr.",
			Value:   "Dr.",
			Label:   "Dr.",
		},
	}
}

func (co *ContactNameForm) TitleMandatory() bool {
	return true
}

func (co *ContactNameForm) FirstNameType() string {
	return "input:text"
}

func (co *ContactNameForm) FirstNameAttr() map[string]string {
	return map[string]string{"class": "form-control", "id": "contact-first-name"}
}

func (co *ContactNameForm) FirstNameLabel() form.Label {
	return form.Label{"First Name:", "contact-first-name", map[string]string{"class": "control-label"}}
}

func (co *ContactNameForm) FirstNameMandatory() bool {
	return true
}

func (co *ContactNameForm) FirstNameMaxLength() int64 {
	return 30
}

func (co *ContactNameForm) FirstNameMaxLengthErr() string {
	return "Must be 30 or below"
}

func (co *ContactNameForm) LastNameType() string {
	return "input:text"
}

func (co *ContactNameForm) LastNameAttr() map[string]string {
	return map[string]string{"class": "form-control", "id": "contact-last-name"}
}

func (co *ContactNameForm) LastNameLabel() form.Label {
	return form.Label{"Last Name:", "contact-last-name", map[string]string{"class": "control-label"}}
}

func (co *ContactNameForm) LastNameMaxLength() int64 {
	return 30
}

func (co *ContactNameForm) LastNameMaxLengthErr() string {
	return "Must be 30 or below"
}

func main() {
	m := martini.Classic()

	f := func(res http.ResponseWriter, req *http.Request) {
		flash := ""
		// It's got to be a pointer, always create it inside the user request scope, to avoid race condition.
		contactForm := &ContactNameForm{}
		get := func() {
			res.Header().Set("Content-Type", "text/html")
			fmt.Fprintln(res, `<form action="." method="post">`)
			fmt.Fprintln(res, flash, `<br />`)
			form.Render(res, contactForm)
			fmt.Fprintln(res, `<input type="submit" value="Submit"></form>`)
		}
		switch req.Method {
		case "GET":
			get()
		case "POST":
			if form.Validate(contactForm, res, req) {
				flash = fmt.Sprintf("Contact Name Successfully submitted, %s %s %s", contactForm.Title, contactForm.FirstName, contactForm.LastName)
			} else {
				flash = "Fail"
			}
			// Save you from relying on sessions!
			get()
		}
	}

	m.Get("/", f)
	m.Post("/", f)
	m.Run()
}
~~~

## Note

There is also a cmd app that helps save you from repetitive work. Assuming you got 'bin' in your 'path' environment variable. 

~~~
go install github.com/cjtoolkit/form_method
form_method co *ContactNameForm Title FirstName LastName
~~~

It's generate the method for you.