package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is foo page")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is bar page")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":80", nil)
}
