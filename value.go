package form

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
)

type value struct {
	Form          url.Values
	PostForm      url.Values
	MultipartForm *multipart.Form
}

func newValue(r *http.Request) *value {
	v := &value{}
	count := 0
	if r.Form != nil {
		if len(r.Form) > 0 {
			v.Form = url.Values{}
		}
		for key, value := range r.Form {
			count++
			v.Form[key] = append(v.Form[key], value...)
		}
	}
	if r.PostForm != nil {
		if len(r.PostForm) > 0 {
			v.PostForm = url.Values{}
		}
		for key, value := range r.PostForm {
			count++
			v.PostForm[key] = append(v.PostForm[key], value...)
		}
	}
	if r.MultipartForm != nil {
		if len(r.MultipartForm.Value) > 0 {
			if v.MultipartForm == nil {
				v.MultipartForm = &multipart.Form{}
			}
			v.MultipartForm.Value = map[string][]string{}
		}
		for key, value := range r.MultipartForm.Value {
			count++
			v.MultipartForm.Value[key] = append(v.MultipartForm.Value[key], value...)
		}
		if len(r.MultipartForm.File) > 0 {
			if v.MultipartForm == nil {
				v.MultipartForm = &multipart.Form{}
			}
			v.MultipartForm.File = map[string][]*multipart.FileHeader{}
		}
		for key, value := range r.MultipartForm.File {
			count++
			v.MultipartForm.File[key] = append(v.MultipartForm.File[key], value...)
		}
	}
	if count == 0 {
		return nil
	}
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

	return v
}

func (r *value) Shift(name string) string {
	str := ""
	if r.MultipartForm != nil && len(r.MultipartForm.Value[name]) > 0 {
		if len(r.MultipartForm.Value[name]) <= 1 {
			if len(r.MultipartForm.Value[name]) == 1 {
				str = r.MultipartForm.Value[name][0]
			}
			delete(r.MultipartForm.Value, name)
		} else {
			str, r.MultipartForm.Value[name] =
				r.MultipartForm.Value[name][0], r.MultipartForm.Value[name][1:]
		}
	} else if r.Form != nil && len(r.Form[name]) > 0 {
		if len(r.Form[name]) <= 1 {
			if len(r.Form[name]) == 1 {
				str = r.Form[name][0]
			}
			delete(r.Form, name)
		} else {
			str, r.Form[name] =
				r.Form[name][0], r.Form[name][1:]
		}
	} else if r.PostForm != nil && len(r.PostForm[name]) > 0 {
		if len(r.PostForm[name]) <= 1 {
			if len(r.PostForm[name]) == 1 {
				str = r.PostForm[name][0]
			}
			delete(r.PostForm, name)
		} else {
			str, r.PostForm[name] =
				r.PostForm[name][0], r.PostForm[name][1:]
		}
	}
	return str
}

func (r *value) All(name string) []string {
	strs := []string{}
	if r.MultipartForm != nil && len(r.MultipartForm.Value[name]) > 0 {
		strs = r.MultipartForm.Value[name]
		delete(r.MultipartForm.Value, name)
	} else if r.Form != nil && len(r.Form[name]) > 0 {
		strs = r.Form[name]
		delete(r.Form, name)
	} else if r.PostForm != nil && len(r.PostForm[name]) > 0 {
		strs = r.PostForm[name]
		delete(r.PostForm, name)
	}
	return strs
}

func (r *value) FileShift(name string) *multipart.FileHeader {
	var fileHeader *multipart.FileHeader
	if r.MultipartForm != nil && len(r.MultipartForm.File[name]) > 0 {
		if len(r.MultipartForm.File[name]) <= 1 {
			if len(r.MultipartForm.File[name]) == 1 {
				fileHeader = r.MultipartForm.File[name][0]
			}
			delete(r.MultipartForm.File, name)
		} else {
			fileHeader, r.MultipartForm.File[name] =
				r.MultipartForm.File[name][0], r.MultipartForm.File[name][1:]
		}
	}
	return fileHeader
}
