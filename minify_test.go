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

func TestMinify(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/", nil)
	bodyBytes := []byte(testHtmlString)

	// Auto detect content type
	t1 := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(bodyBytes)
		}),
	)
	if t1.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", t1.Code)
	}
	if len(bodyBytes) == len(t1.Body.Bytes()) {
		t.Errorf("Request Body is not minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(t1.Body.Bytes()))
	}

	// Using content type value as text/html
	t2 := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(bodyBytes)
		}),
	)
	if t2.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", t2.Code)
	}
	if len(bodyBytes) == len(t2.Body.Bytes()) {
		t.Errorf("Request Body is not minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(t2.Body.Bytes()))
	}

	// Using content type value as application/html
	t3 := performTest(r,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/html; charset=utf-8")
			w.Write(bodyBytes)
		}),
	)
	if t3.Code != http.StatusOK {
		t.Errorf("Request failed %d, not 200", t3.Code)
	}
	if len(bodyBytes) == len(t3.Body.Bytes()) {
		t.Errorf("Request Body is not minified [original bytes: %d, response bytes: %d]", len(bodyBytes), len(t3.Body.Bytes()))
	}
}
