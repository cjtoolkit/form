package form

import (
	"fmt"
	"mime/multipart"
	"strconv"
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
	tmpl, err := text.New("").Funcs(text.FuncMap{
		"list": templateListFilter,
	}).Parse(tpl)
	if nil != err {
		panic(err)
	}
	return tmpl
}

func templateListFilter(and string, value interface{}) (str string) {
	switch value := value.(type) {
	case []string:
		switch valueLen := len(value); valueLen {
		case 0:
			return
		case 1:
			str = value[0]
		default:
			last := value[valueLen-1]
			value = value[:valueLen-1]
			newValue := []interface{}{}
			for _, v := range value {
				newValue = append(newValue, v, ", ")
			}
			str = fmt.Sprint(append(newValue[:len(newValue)-1], " ", and, " ", last)...)
		}
	case []int64:
		switch valueLen := len(value); valueLen {
		case 0:
			return
		case 1:
			str = strconv.FormatInt(value[0], 10)
		default:
			last := strconv.FormatInt(value[valueLen-1], 10)
			value = value[:valueLen-1]
			newValue := []interface{}{}
			for _, v := range value {
				newValue = append(newValue, strconv.FormatInt(v, 10), ", ")
			}
			str = fmt.Sprint(append(newValue[:len(newValue)-1], " ", and, " ", last)...)
		}
	case []float64:
		switch valueLen := len(value); valueLen {
		case 0:
			return
		case 1:
			str = strconv.FormatFloat(value[0], 'f', -1, 64)
		default:
			last := strconv.FormatFloat(value[valueLen-1], 'f', -1, 64)
			value = value[:valueLen-1]
			newValue := []interface{}{}
			for _, v := range value {
				newValue = append(newValue, strconv.FormatFloat(v, 'f', -1, 64), ", ")
			}
			str = fmt.Sprint(append(newValue[:len(newValue)-1], " ", and, " ", last)...)
		}
	}
	return
}
