package main

import (
	"fmt"
	"github.com/cjtoolkit/form"
	"io"
	"mime/multipart"
	"net/http"
)

type TestForm struct {
	form.Form
	File *multipart.FileHeader
}

func (t *TestForm) FileType() string {
	return "input:file"
}

func (t *TestForm) FileSize() int64 {
	return 10 * 1024 * 1024
}

func (t *TestForm) FileAccept() []string {
	return []string{"image/jpeg", "image/png"}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		f := &TestForm{}
		get := func() {
			fmt.Fprint(res, "<h1>File Test</h1>")

			fmt.Fprint(res, `<form action="/" method="post" enctype="multipart/form-data">
`)
			form.Render(res, f)

			fmt.Fprint(res, `<input type="submit" name="submit" value="Submit">
</form>`)
		}

		switch req.Method {
		case "GET":
			get()
		case "POST":
			if !form.Validate(f, res, req) {
				get()
				return
			}

			file, _ := f.File.Open()
			io.Copy(res, file)
		}
	})

	http.ListenAndServe(":8080", mux)
}
