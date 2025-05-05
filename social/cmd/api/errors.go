package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	app.logger.Errorw("Internal server error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusInternalServerError, "The server encountered an internal error and was unable to complete your request.")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	app.logger.Warnf("Bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	app.logger.Errorf("Not found", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusNotFound, "Not Found")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	// Log the error and send a generic message to the client.
	app.logger.Warnf("Conflict", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusConflict, err.Error())
}
