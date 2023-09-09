package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// create a new Middleware
	return func(next http.HandlerFunc) http.HandlerFunc {

		// define the handlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// call the next middleware/handler in chain
			next(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 bad request
func Method(m string) Middleware {

	// Create a middleware
	return func(next http.HandlerFunc) http.HandlerFunc {

		// define a handler func
		return func(w http.ResponseWriter, r *http.Request) {

			// do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// call the next middleware/handler in chain
			next(w, r)
		}
	}
}

// Chain Applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
