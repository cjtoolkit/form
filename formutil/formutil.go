package formutil

import (
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// Parse Url Query string
func ParseUrlQuery(r *http.Request) {
	r.Form, _ = url.ParseQuery(r.URL.RawQuery)
}

// Parse Body of Request
func ParseBody(r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	r.PostForm, _ = url.ParseQuery(string(body))
}

// Parse Body of Request (Multipart)
func ParseMultipartBody(r *http.Request, maxMemory int64) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return
	}
	if !strings.HasPrefix(mediaType, "multipart/") {
		return
	}

	mr := multipart.NewReader(r.Body, params["boundary"])

	r.MultipartForm, _ = mr.ReadForm(maxMemory)
}
