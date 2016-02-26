package form

import (
	"mime/multipart"
	text "text/template"
)

func GetOneFile(values ValuesInterface, name string) (fh *multipart.FileHeader) {
	if values, ok := values.(ValuesFileInterface); ok {
		fh = values.GetOneFile(name)
	}
	return
}

func GetAllFile(values ValuesInterface, name string) (fhs []*multipart.FileHeader) {
	if values, ok := values.(ValuesFileInterface); ok {
		fhs = values.GetAllFile(name)
	}
	return
}

func BuildLanguageTemplate(tpl string) *text.Template {
	tmpl, err := text.New("").Parse(tpl)
	if nil != err {
		panic(err)
	}
	return tmpl
}
