/*
Form Rendering and Validation System.

Usage

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
			f := f.Init(&t.Text, "Text", form.InputText)
			html := f.HTML()
			html.Before = "<h1>File Test</h1>"
		}()

		// File
		func() {
			f := f.Init(&t.File, "File", form.InputFile)
			file := f.File()
			file.Size = 10 * 1024 * 1024
			file.Accept = []string{"image/jpeg", "image/png"}
		}()
	}

	func main() {
		mux := http.NewServeMux()

		mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			f := &TestForm{}
			v := form.New(req, nil, "en-GB")
			get := func() {
				res.Header().Set("Content-Type", "text/html")

				fmt.Fprint(res, `<form action="/" method="post" enctype="multipart/form-data">
	`)
				v.Render(res, f)

				fmt.Fprint(res, `<input type="submit" name="submit" value="Submit">
	</form>`)
			}

			switch req.Method {
			case "GET":
				get()
			case "POST":
				req.ParseMultipartForm(10 * 1024 * 1024)
				if !v.MustValidate(f) && f.File == nil {
					get()
					return
				}

				file, _ := f.File.Open()
				io.Copy(res, file)
			}
		})

		http.ListenAndServe(":8080", mux)
	}
*/
package form
