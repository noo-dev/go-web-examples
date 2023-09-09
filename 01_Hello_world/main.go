package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You requested to %v, and value of token paramater is %s \n", r.URL.Path, r.URL.Query().Get("token"))
	})

	http.ListenAndServe(":8080", nil)
}
