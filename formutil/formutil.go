package formutil

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// Parse Url Query string
func ParseUrlQuery(r *http.Request) {
	Form, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return
	}
	r.Form = Form
}

// Parse Body of Request
func ParseBody(r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	r.Body.Close()

	PostForm, err := url.ParseQuery(string(body))
	if err != nil {
		return
	}
	r.PostForm = PostForm
}

// Parse Body of Request (Multipart)
func ParseMultipartBody(r *http.Request, maxMemory int64) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || !strings.HasPrefix(mediaType, "multipart/") {
		return
	}

	mr := multipart.NewReader(r.Body, params["boundary"])
	defer r.Body.Close()

	MultipartForm, err := mr.ReadForm(maxMemory)
	if err != nil {
		return
	}

	r.MultipartForm = MultipartForm
}

// Parse http://api.jquery.com/serializearray/ in JSON format
// Request Body must be in JSON format. 'JSON.stringify(object);' in Javascript.
// Eg [{"name":"","value":""},{"name":"","value":""},{"name":"","value":""}...]
func ParseJQuerySerializeArrayBody(r *http.Request) {
	data := []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{}

	jDec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := jDec.Decode(&data)
	if err != nil {
		return
	}

	r.PostForm = url.Values{}

	for _, item := range data {
		r.PostForm.Add(item.Name, item.Value)
	}
}
