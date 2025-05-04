package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	log.Printf("Internal server error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusInternalServerError, "The server encountered an internal error and was unable to complete your request.")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	log.Printf("Bad Request error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	log.Printf("Not Found error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusNotFound, "Not Found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	log.Printf("Conflict error: %s Path: %s Error: %s", r.Method, r.URL.Path, err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}
