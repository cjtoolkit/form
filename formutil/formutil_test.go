package formutil

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUtil(t *testing.T) {

	var call func(w http.ResponseWriter, r *http.Request)

	handler := func(w http.ResponseWriter, r *http.Request) {
		call(w, r)
	}

	testserver := httptest.NewServer(http.HandlerFunc(handler))

	// Parse Url Query String

	val := url.Values{}
	val.Set("hello", "world")

	call = func(w http.ResponseWriter, r *http.Request) {
		ParseUrlQuery(r)

		if r.Form.Get("hello") != "world" {
			t.Error("Parse Url Query String: 'hello' is not equal to 'world'")
		}
	}

	http.DefaultClient.Get(testserver.URL + "?" + val.Encode())

	// Parse Body of Request

	call = func(w http.ResponseWriter, r *http.Request) {
		ParseBody(r)

		if r.PostForm.Get("hello") != "world" {
			t.Error("Parse Body of Request: 'hello' is not equal to 'world'")
		}
	}

	http.DefaultClient.PostForm(testserver.URL, val)

	// Parse Body of Request (Multipart)

	call = func(w http.ResponseWriter, r *http.Request) {
		ParseMultipartBody(r, 10*(1024*2))

		if r.MultipartForm.Value["hello"][0] != "world" {
			t.Error("Parse Body of Request (Multipart): 'hello' is not equal to 'world'")
		}
	}

	buffer := &bytes.Buffer{}

	mp := multipart.NewWriter(buffer)

	bodyType := fmt.Sprintf("multipart/mixed; boundary=%s", mp.Boundary())

	mp.WriteField("hello", "world")
	mp.Close()

	http.DefaultClient.Post(testserver.URL, bodyType, buffer)

	call = func(w http.ResponseWriter, r *http.Request) {
		ParseJQuerySerializeArrayBody(r)

		if r.PostForm.Get("hello") != "world" {
			t.Error("Parse Body of Request (jQuery): 'hello' is not equal to 'world'")
		}

		if r.PostForm.Get("world") != "hello" {
			t.Error("Parse Body of Request (jQuery): 'world' is not equal to 'hello'")
		}
	}

	strr := strings.NewReader(`[{"name":"hello","value":"world"},{"name":"world","value":"hello"}]`)

	http.DefaultClient.Post(testserver.URL, "application/json", strr)

}
