package main

import (
	"net/http"
)


func main() {
	api := &api{addr: ":8080"}

	// init server mux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
