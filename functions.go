package form

import (
	"fmt"
	"mime/multipart"
	"sort"
	"strconv"
	text "text/template"
)

// Get One File from Value
func GetOneFile(values ValuesInterface, name string) (fh *multipart.FileHeader) {
	if values, ok := values.(ValuesFileInterface); ok {
		fh = values.GetOneFile(name)
	}
	return
}

// Get All File from Value
func GetAllFile(values ValuesInterface, name string) (fhs []*multipart.FileHeader) {
	if values, ok := values.(ValuesFileInterface); ok {
		fhs = values.GetAllFile(name)
	}
	return
}

func BuildLanguageTemplate(tpl string) *text.Template {
	return text.Must(text.New("").Funcs(text.FuncMap{
		"list":      templateListFilter,
		"pluralise": templatePluraliseFilter,
		"pluralize": templatePluraliseFilter, // Knowing the Americans
	}).Parse(tpl))
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

func templatePluraliseFilter(plural string, count interface{}) (output string) {
	switch count := count.(type) {
	case int64:
		if count > 1 {
			output = plural
		}
	case int:
		if count > 1 {
			output = plural
		}
	case uint64:
		if count > 1 {
			output = plural
		}
	case float64:
		if count > 1 {
			output = plural
		}
	}
	return
}

// Returns nil if all is good.
func CheckAllFields(fields []FormFieldInterface) interface{} {
	panicChannel := make(chan interface{})

	go func() {
		defer func() {
			panicChannel <- recover()
		}()

		for _, field := range fields {
			field.PreCheck()
		}
	}()

	return <-panicChannel
}

// Binary Search String
func SearchString(a []string, x string) bool {
	count := len(a)
	if 0 == count {
		return false
	}
	i := sort.Search(count, func(i int) bool { return a[i] >= x })
	return x == a[i]
}

// Binary Search Int64
func SearchInt64(a []int64, x int64) bool {
	count := len(a)
	if 0 == count {
		return false
	}
	i := sort.Search(count, func(i int) bool { return a[i] >= x })
	return x == a[i]
}
