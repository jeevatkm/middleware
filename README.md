# Middleware [![Build Status](https://travis-ci.org/jeevatkm/middleware.svg?branch=master)](https://travis-ci.org/jeevatkm/middleware) [![GoDoc](https://godoc.org/github.com/jeevatkm/middleware?status.svg)](https://godoc.org/github.com/jeevatkm/middleware) [![GoCover](http://gocover.io/_badge/github.com/jeevatkm/middleware)](http://gocover.io/github.com/jeevatkm/middleware)
A collection of HTTP middleware/Handler function for use with Go's net/http package. Compatible with Goji, Gorilla, Negroni & net/http (amongst many others).

* Minify HTTP middleware using [tdewolff/minify](https://github.com/tdewolff/minify)
* will be adding few...

## Start using it
* Installation
```sh
go get github.com/jeevatkm/middleware
                   OR
go get gopkg.in/jeevatkm/middleware.v0
```

* Import it in your code
```go
import "github.com/jeevatkm/middleware"
                   OR
import "gopkg.in/jeevatkm/middleware.v0"
```

## Examples
Refer [examples](https://github.com/jeevatkm/middleware/tree/master/examples)

#### Vanilla net/http

```go
func main() {
	homeHandler := http.HandlerFunc(home)

	// Adding Minify middleware
	// Note: If you use any Gzip middleware, add Minify middleware after that
	http.Handle("/", middleware.Minify(homeHandler))

	log.Println("Starting server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
```

#### Goji web framework

```go
func main() {
	// Adding Minify middleware
	// Note: If you use any Gzip middleware, add Minify middleware after that
	goji.Use(middleware.Minify)

	goji.Get("/", gojiHome)
	goji.Serve()
}
```

#### Gorilla Mux

```go
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", gorillaHome)

	// Adding Minify middleware
	// Note: If you use any Gzip middleware, add Minify middleware after that
	http.Handle("/", middleware.Minify(r))

	log.Println("Starting server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
```

#### Negroni web middleware

```go
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", negroniHome)

	n := negroni.Classic()

	// Adding Minify middleware
	n.UseHandler(middleware.Minify(mux))

	n.Run(":8000")
}
```

## License
Middleware released under [MIT License](https://github.com/jeevatkm/middleware/blob/master/LICENSE)
