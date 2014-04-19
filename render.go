package form

import (
	"bytes"
	"fmt"
	"github.com/gorail/core"
	"io"
	"mime/multipart"
	"reflect"
	"time"
)

type render struct {
	form
	w io.Writer
}

func (r render) str() {
	switch r.ftype {
	case "input:text":
		r.strInputText("text")
	case "input:search":
		r.strInputText("search")
	case "input:password":
		r.strInputText("password")
	case "input:hidden":
		r.strInputText("hidden")
	case "input:url":
		r.strInputText("url")
	case "input:tel":
		r.strInputText("tel")
	case "input:email":
		r.strInputEmail()
	case "input:radio":
		r.strInputRadio()
	case "input:color":
		r.strInputColor()
	case "textarea":
		r.strTextarea()
	case "select":
		r.strSelect()
	}
}

func (r render) strs() {
	switch r.ftype {
	case "select":
		r.strsSelect()
	}
}

func (r render) wnum() {
	switch r.ftype {
	case "input:number":
		r.numInputNumber(false, "number")
	case "input:range":
		r.numInputNumber(false, "range")
	case "input:hidden":
		r.numInputNumber(false, "hidden")
	case "input:radio":
		r.wnumInputRadio()
	case "select":
		r.wnumSelect()
	}
}

func (r render) wnums() {
	switch r.ftype {
	case "select":
		r.wnumsSelect()
	}
}

func (r render) fnum() {
	switch r.ftype {
	case "input:number":
		r.numInputNumber(true, "number")
	case "input:range":
		r.numInputNumber(true, "range")
	case "input:hidden":
		r.numInputNumber(true, "hidden")
	case "input:radio":
		r.fnumInputRadio()
	case "select":
		r.fnumSelect()
	}
}

func (r render) fnums() {
	switch r.ftype {
	case "select":
		r.fnumsSelect()
	}
}

func (r render) b() {
	switch r.ftype {
	case "input:checkbox":
		r.bInputCheckbox(false)
	case "input:hidden":
		r.bInputCheckbox(true)
	}
}

func (r render) time() {
	switch r.ftype {
	case "input:datetime":
		r.timeInputTime("datetime")
	case "input:datetime-local":
		r.timeInputTime("datetime-local")
	case "input:time":
		r.timeInputTime("time")
	case "input:date":
		r.timeInputTime("date")
	case "input:month":
		r.timeInputTime("month")
	case "input:week":
		r.timeInputTime("week")
	}
}

func (r render) file() {
	switch r.ftype {
	case "input:file":
		r.fileInputFile()
	}
}

func (r render) attr(exclude ...string) {
	mstring, ok := r.get("Attr").(map[string]string)
	if !ok {
		return
	}

	for _, value := range exclude {
		delete(mstring, value)
	}

	for key, value := range mstring {
		fmt.Fprintf(r.w, `%s="%s" `, es(key), es(value))
	}
}

// Render Form to Writer (Must be struct with pointer)
func Render(w io.Writer, v FormInterface) (err error) {
	t := reflect.TypeOf(v)
	vc := reflect.ValueOf(v)
	vcc := vc

	switch {
	case isStructPtr(t):
		t = t.Elem()
		vc = vc.Elem()
	default:
		err = fmt.Errorf("%v must be a struct pointer", v)
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := vc.Field(i)
		if !field.CanSet() {
			continue
		}
		name := t.Field(i).Name
		preferedName := name
		tag := t.Field(i).Tag.Get("form")
		if tag == "-" {
			continue
		}
		if tag != "" {
			preferedName = tag
		}

		ww := &bytes.Buffer{}
		wrap := &bytes.Buffer{}

		r := render{form{vcc, t, vc, t.Field(i), vc.Field(i), name, preferedName, ""}, wrap}

		ftype, ok := r.getStr("Type")
		if !ok {
			continue
		}

		r.form.ftype = ftype

		groupFormat := v.Group()

		WrapLabel := false

		switch label := r.get("Label").(type) {
		case string:
			fmt.Fprint(ww, `<label>`, es(label), `: `)
			WrapLabel = true
		case Label:
			labelAttr := func() string {
				if label.Attr == nil {
					return ""
				}

				labelBuf := &bytes.Buffer{}
				defer labelBuf.Reset()

				for key, value := range label.Attr {
					fmt.Fprintf(labelBuf, ` %s="%s"`, es(key), es(value))
				}

				return labelBuf.String()
			}
			fmt.Fprint(ww, `<label for="`, es(label.For), `"`, labelAttr(), `>`, es(label.Content), `</label>`)
			fmt.Fprintln(ww)
		}

		switch vc.Field(i).Interface().(type) {
		case string:
			r.str()
		case []string:
			r.strs()
		case int64:
			r.wnum()
		case []int64:
			r.wnums()
		case float64:
			r.fnum()
		case []float64:
			r.fnums()
		case bool:
			r.b()
		case time.Time:
			r.time()
		case *multipart.FileHeader:
			r.file()
		}

		fmt.Fprintf(ww, v.Wrap(), wrap.String())
		wrap.Reset()

		err := v.Err(name)
		if err != nil {
			fmt.Fprintf(ww, v.ErrFormat(), es(err.Error()))
			fmt.Fprint(ww, ` `)
			groupFormat = v.GroupError()
		} else if warning := v.GetWarning(name); warning != "" {
			fmt.Fprintf(ww, v.WarningFormat(), es(warning))
			fmt.Fprint(ww, ` `)
			groupFormat = v.GroupWarning()
		} else if v.BeenChecked() {
			groupFormat = v.GroupSuccess()
		}

		if WrapLabel {
			fmt.Fprint(ww, `</label>`)
		}

		fmt.Fprintf(w, groupFormat, ww.String())
		ww.Reset()

		fmt.Fprintln(w)
	}

	return
}

// Render Form and Return Byte (Must be struct with pointer)
func RenderBytes(v FormInterface) []byte {
	buf := &bytes.Buffer{}
	defer buf.Reset()
	core.Check(Render(buf, v))
	return buf.Bytes()
}

// Render Form and Return String (Must be struct with pointer)
func RenderString(v FormInterface) string {
	return string(RenderBytes(v))
}
