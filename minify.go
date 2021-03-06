package middleware

import (
	"bytes"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	minifier  *minify.Minify
	mediaType *regexp.Regexp
)

func init() {
	minifier = minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("text/javascript", js.Minify)

	mediaType = regexp.MustCompile("text/[html|css|javascript]")
}

type minifyWriter struct {
	http.ResponseWriter
	Body        *bytes.Buffer
	code        int
	wroteHeader bool
}

func (m *minifyWriter) Header() http.Header {
	return m.ResponseWriter.Header()
}

func (m *minifyWriter) WriteHeader(code int) {
	if !m.wroteHeader {
		m.code = code
		m.wroteHeader = true
		m.ResponseWriter.WriteHeader(code)
	}
}

func (m *minifyWriter) Write(b []byte) (int, error) {
	h := m.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}

	if !m.wroteHeader {
		m.WriteHeader(http.StatusOK)
	}

	if m.Body != nil {
		m.Body.Write(b)
	}
	return len(b), nil
}

/*
Minify middleware is simple and generic using 'tdewolff/minify' package.
Middleware minifies HTML, Inline CSS and Inline JS. Compatible with Goji,
Gorilla, Gin & net/http (amongst many others).
Refer: https://github.com/jeevatkm/middleware/examples
*/
func Minify(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := &minifyWriter{
			ResponseWriter: w,
			Body:           &bytes.Buffer{},
		}

		h.ServeHTTP(mw, r)

		hdr := w.Header()
		ct := hdr.Get("Content-Type")
		if mediaType.MatchString(ct) {
			rb, err := minify.Bytes(minifier, ct, mw.Body.Bytes())
			if err != nil {
				_ = err // unsupported mediatype error or internal
			}

			hdr.Del("Content-Length")
			hdr.Set("Content-Length", strconv.Itoa(len(rb)))

			w.Write(rb)
		} else {
			w.Write(mw.Body.Bytes())
		}
	})
}
