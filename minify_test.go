package middleware

import (
	"net/http"
	"net/http/httptest"

	"testing"
)

var testHtmlString = `
		<!DOCTYPE html>
		<html>
		<head>

		    <meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">


<meta name="viewport" content="width=device-width, initial-scale=1">

		</head>

		<body>

		<h1>Minify Test</h1>

		<p>Test paragraph.</p>
		<div id="example"></div>



		</body>
		</html>`

func performTest(r *http.Request, h http.Handler) *httptest.ResponseRecorder {
	m := Minify(h)
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)

	return w
}

func TestMinifyAutoDetectContentType(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	bodyBytes := []byte(testHtmlString)

	res := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(bodyBytes)
		}),
	)
	if res.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", res.Code)
	}
	if len(bodyBytes) == len(res.Body.Bytes()) {
		t.Errorf("Request Body is not minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(res.Body.Bytes()))
	}
}

func TestMinifyContentTypeTextHTML(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	bodyBytes := []byte(testHtmlString)

	res := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(bodyBytes)
		}),
	)
	if res.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", res.Code)
	}
	if len(bodyBytes) == len(res.Body.Bytes()) {
		t.Errorf("Request Body is not minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(res.Body.Bytes()))
	}
}

func TestNonMinifyContentTypeApplicationHTML(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	bodyBytes := []byte(testHtmlString)

	res := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/html; charset=utf-8")
			w.Write(bodyBytes)
		}),
	)
	if res.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", res.Code)
	}
	if len(bodyBytes) != len(res.Body.Bytes()) {
		t.Errorf("Request Body is minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(res.Body.Bytes()))
	}
}

func TestNonMinifyContentType(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	bodyBytes := []byte(`{"type":"Non-Minify contentType"}`)

	res := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(bodyBytes)
		}),
	)
	if res.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", res.Code)
	}
	if len(bodyBytes) != len(res.Body.Bytes()) {
		t.Errorf("Request Body is minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(res.Body.Bytes()))
	}
}
