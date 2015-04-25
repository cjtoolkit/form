package form

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
)

type value struct {
	Form                url.Values
	PostForm            url.Values
	MultipartForm       *multipart.Form
	formCount           map[string]int
	postFormCount       map[string]int
	multipartValueCount map[string]int
	multipartFileCount  map[string]int
}

func newValue(r *http.Request) *value {
	v := &value{}
	count := 0
	if len(r.Form) > 0 {
		v.Form = r.Form
		count++
	}
	if len(r.PostForm) > 0 {
		v.PostForm = r.PostForm
		count++
	}
	if r.MultipartForm != nil {
		v.MultipartForm = r.MultipartForm
		count++
	}
	if count == 0 {
		return nil
	}
	v.formCount = map[string]int{}
	v.postFormCount = map[string]int{}
	v.multipartValueCount = map[string]int{}
	v.multipartFileCount = map[string]int{}
	return v
}

func newValueSerializeArray(r *http.Request) *value {
	v := &value{PostForm: url.Values{}}

	data := []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{}

	jDec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	jDec.Decode(&data)

	for _, item := range data {
		v.PostForm[item.Name] = append(v.PostForm[item.Name], item.Value)
	}
	v.formCount = map[string]int{}
	v.postFormCount = map[string]int{}
	v.multipartValueCount = map[string]int{}
	v.multipartFileCount = map[string]int{}
	return v
}

func (r *value) Shift(name string) (str string) {
	switch {
	case r.MultipartForm != nil &&
		len(r.MultipartForm.Value[name]) > r.multipartValueCount[name]:

		str = r.MultipartForm.Value[name][r.multipartValueCount[name]]
		r.multipartValueCount[name]++

	case len(r.Form[name]) > r.formCount[name]:

		str = r.Form[name][r.formCount[name]]
		r.formCount[name]++

	case len(r.PostForm[name]) > r.postFormCount[name]:

		str = r.PostForm[name][r.postFormCount[name]]
		r.postFormCount[name]++
	}

	return
}

func (r *value) All(name string) []string {
	strs := []string{}

	switch {
	case r.MultipartForm != nil &&
		len(r.MultipartForm.Value[name]) > r.multipartValueCount[name]:

		strs = r.MultipartForm.Value[name][r.multipartValueCount[name]:]
		r.multipartValueCount[name] = len(r.MultipartForm.Value[name])

	case len(r.Form[name]) > r.formCount[name]:

		strs = r.Form[name][r.formCount[name]:]
		r.formCount[name] = len(r.Form[name])

	case len(r.PostForm[name]) > r.postFormCount[name]:

		strs = r.PostForm[name][r.postFormCount[name]:]
		r.postFormCount[name] = len(r.PostForm[name])
	}

	return strs
}

func (r *value) FileShift(name string) (fileHeader *multipart.FileHeader) {
	if r.MultipartForm != nil &&
		len(r.MultipartForm.File[name]) > r.multipartFileCount[name] {

		fileHeader = r.MultipartForm.File[name][r.multipartFileCount[name]]
		r.multipartFileCount[name]++

	}

	return
}
