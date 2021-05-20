package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is my Homepage</h1>")
}

func contact(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(
		w,
		"To get in touch, please email to <a href=\"mailto:contact@example.com\">contact@example.com</a>.",
	)
}

func notFound(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(
		w,
		"<h2>We could not find the page!</h2>",
	)
}

func main() {
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact", contact)
	r.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}
