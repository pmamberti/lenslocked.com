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
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
