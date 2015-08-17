# Middleware [![Build Status](https://travis-ci.org/jeevatkm/middleware.svg?branch=master)](https://travis-ci.org/jeevatkm/middleware) [![GoDoc](https://godoc.org/github.com/jeevatkm/middleware?status.svg)](https://godoc.org/github.com/jeevatkm/middleware)
A collection of HTTP middleware/Handler function for use with Go's net/http package. Compatible with Goji, Gorilla, Gin & net/http (amongst many others).

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

	goji.Get("/", home)
	goji.Serve()
}
```

## License
Middleware released under [MIT License](https://github.com/jeevatkm/middleware/blob/master/LICENSE)
