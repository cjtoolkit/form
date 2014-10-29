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

	func (t *TestForm) TextField() form.FieldFuncs {
		return form.FieldFuncs{
			"form": func(m map[string]interface{}) {
				*(m["type"].(*form.TypeCode)) = form.InputText
			},
			"html": func(m map[string]interface{}) {
				*(m["before"].(*string)) = "<h1>File Test</h1>"
			},
		}
	}

	func (t *TestForm) FileField() form.FieldFuncs {
		return form.FieldFuncs{
			"form": func(m map[string]interface{}) {
				*(m["type"].(*form.TypeCode)) = form.InputFile
			},
			"file": func(m map[string]interface{}) {
				*(m["size"].(*int64)) = 10 * 1024 * 1024
				*(m["accept"].(*[]string)) = []string{"image/jpeg", "image/png"}
			},
		}
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
				if !v.MustValidate(req, f) {
					get()
					return
				}

				if f.File == nil {
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
