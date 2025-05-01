package main

import (
	"log"
	"net/http"
	"social/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string //time.Duration
}

func (app *application) mount() http.Handler {
	// reference: https://github.com/go-chi/chi
	// Create a new chi router instance
	r := chi.NewRouter()

	r.Use(middleware.Logger)    // Log the start and end of each request
	r.Use(middleware.Recoverer) // Recover from panics and log the error
	r.Use(middleware.RealIP)    // Get the real IP address of the client
	r.Use(middleware.RequestID) // Generate a unique request ID for each request

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		// /v1/health
		r.Get("/health", app.healthCheckHandler)

		// /v1/posts
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostsHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", app.getPostHandler)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Set a write timeout of 30 seconds
		ReadTimeout:  time.Second * 10, // Set a read timeout of 10 seconds
		IdleTimeout:  time.Minute,      // Set an idle timeout of 1 minute
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()
}
