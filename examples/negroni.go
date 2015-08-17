/*
Example for How to use with Negroni web middleware
https://github.com/codegangsta/negroni
*/
package main

import (
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/jeevatkm/middleware"
)

var htmlString = `
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

		<h1>Minify #NAME# Example</h1>

		<p>Example #NAME# paragraph.</p>
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
		</html>`

func negroniHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Replace(htmlString, "#NAME#", "Negroni", -1)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", negroniHome)

	n := negroni.Classic()

	// Adding Minify middleware
	n.UseHandler(middleware.Minify(mux))

	n.Run(":8000")
}
