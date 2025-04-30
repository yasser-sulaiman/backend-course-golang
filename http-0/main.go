package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Welcome to the home page!"))
			return
		case "/users":
			w.Write([]byte("users us page"))
			return
		default:
			w.Write([]byte("404 Not Found"))
			return
		}
	case http.MethodPost:
		switch r.URL.Path {
		case "/submit":
			w.Write([]byte("Form submitted successfully!"))
		}
	}
}

func main() {
	api := &api{addr: ":8080"}

	srv := &http.Server{
		Addr:    api.addr,
		Handler: api,
	}

	srv.ListenAndServe()

}
