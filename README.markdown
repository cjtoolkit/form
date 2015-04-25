# CJToolkit Form

Automated Form Rendering and Validation Library for Google Go.

- Integration with 'github.com/cjtoolkit/i18n'.
  - So it can speak your lingo.
- Dual layer rendering system.
  - So it can easily be adapted to any CSS framework, such as Bootstrap or Foundation.
  - Currently support Bootstrap and Foundation out of the box.
  - First layer is fixed.
  - Second layer is user definable.
    - So you can use your own CSS framework.
- Relies on Struct Pointers with only one methods
  - So you can pretty much do anything you desire.
    - Defining your own rules (see below).
    - i18n integration.
    - Database integration either with or without ORM, Your choice.
  - No Struct tags are needed, not that there anything wrong with them.
  - See example below, than have a look at document, it will help you understand the system.


Documentation can be found at.

 - https://godoc.org/github.com/cjtoolkit/form

## Installation

~~~
go get github.com/cjtoolkit/form
~~~

## Example

### Defining Custom Rules

~~~ go 
func (e *Example) CJForm(f *form.Fields) {
	// Text
	func() {
		f := f.Init("Text", form.InputText)

		f.Custom(func(e *error, w *string) {
			*e = "Error message here!"
			*w = "Warning message here!"
		})
	}()
}
~~~

### Actual Application

~~~ go
package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cjtoolkit/form"
	_ "github.com/cjtoolkit/form/lang/enGB"
)

type TestForm struct {
	Text string
	File *multipart.FileHeader
}

func (t *TestForm) CJForm(f *form.Fields) {
	// Text
	func() {
		f := f.Init("Text", form.InputText)
		html := f.HTML()
		html.Before = "<h1>File Test</h1>"
	}()

	// File
	func() {
		f := f.Init("File", form.InputFile)
		file := f.File()
		file.Size = 10 * 1024 * 1024
		file.Accept = []string{"image/jpeg", "image/png"}
	}()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		f := &TestForm{}
		v := form.New(nil, "en-GB")
		get := func() {
			res.Header().Set("Content-Type", "text/html")

			fmt.Fprint(res, `<form action="/" method="post" enctype="multipart/form-data">
`)
			v.Render(f, res)

			fmt.Fprint(res, `<input type="submit" name="submit" value="Submit">
</form>`)
		}

		switch req.Method {
		case "GET":
			get()
		case "POST":
			req.ParseMultipartForm(10 * 1024 * 1024)
			if !v.MustValidate(req, f) && f.File == nil {
				get()
				return
			}

			file, _ := f.File.Open()
			io.Copy(res, file)
		}
	})

	http.ListenAndServe(":8080", mux)
}
~~~

## Demo

https://formdemo.cj-jackson.com/

## Buy me a beer!

Bitcoin - 1MieXR5ANYY6VstNanhuLRtGQGn6zpjxK3