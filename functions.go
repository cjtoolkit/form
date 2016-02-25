package form

import (
	"mime/multipart"
	text "text/template"
)

func GetOneFile(value ValuesInterface, name string) (fh *multipart.FileHeader) {
	if value, ok := value.(ValuesFileInterface); ok {
		fh = value.GetOneFile(name)
	}
	return
}

func GetAllFile(value ValuesInterface, name string) (fhs []*multipart.FileHeader) {
	if value, ok := value.(ValuesFileInterface); ok {
		fhs = value.GetAllFile(name)
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
