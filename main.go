package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>This is my Homepage</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(
			w,
			"To get in touch, please email to <a href=\"mailto:contact@example.com\">contact@example.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find your page.</h1> <p>If you keep seeeing this, send us an email.</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
