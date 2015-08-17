/*
Example for How to use with standard Goji web framework.
*/
package main

import (
	"net/http"

	"github.com/jeevatkm/middleware"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func home(c web.C, w http.ResponseWriter, r *http.Request) {
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

		<h1>Minify Goji Example</h1>

		<p>Minify Goji paragraph.</p>
		<div id="example"></div>

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
			document.getElementById("example").innerHTML = "Today is " + day;
			</script>

		</body>
		</html>`))
}

func main() {

	// Adding Minify middleware
	// Note: If you use any Gzip middleware, add Minify middleware after that
	goji.Use(middleware.Minify)

	goji.Get("/", home)
	goji.Serve()
}
