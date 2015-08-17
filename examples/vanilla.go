package main

import (
	"log"
	"net/http"

	"github.com/jeevatkm/middleware"
)

func entryHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing entryHandler request in")
		h.ServeHTTP(w, r)
		log.Println("Executing entryHandler response out")
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to Home handler!")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>

		<style>body {
		    background-color: #d0e4fe;
		}

		h1 {
		    color: orange;
		    text-align: center;
		}

		p {
		    font-family: "Times New Roman";
		    font-size: 20px;
		}
		</style>
		</head>

		<body>

		<h1>My Minify Vanilla Testing</h1>

		<p>My vanilla paragraph.</p>
		<div id="demo"></div>

		<script>
			var day;
			switch (new Date().getDay()) {
			    case 0:
			        day = "Sunday";
			        break;
			    case 1:
			        day = "Monday";
			        break;
			    case 2:
			        day = "Tuesday";
			        break;
			    case 3:
			        day = "Wednesday";
			        break;
			    case 4:
			        day = "Thursday";
			        break;
			    case 5:
			        day = "Friday";
			        break;
			    case  6:
			        day = "Saturday";
			        break;
			}
			document.getElementById("demo").innerHTML = "Today is " + day;
			</script>

		</body>
		</html>`))
}

func main() {
	homeHandler := http.HandlerFunc(home)

	http.Handle("/", middleware.Minify(entryHandler(homeHandler)))
	http.ListenAndServe(":3000", nil)
}
