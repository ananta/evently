package main

import (
  "fmt"
  "net/http"
)

// logError(): A helper for logging an error message with an http method & URL
func (app *application) logError(r *http.Request, err error){
  var (
    method = r.Method
    uri = r.URL.RequestURI()
  )
  app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// errorResponse(): A helper for sending JSON-formatted error
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any){
  env := envelope{"error": message}

  err := app.writeJson(w, status, env, nil)
  if err != nil {
    app.logError(r, err)
    w.WriteHeader(http.StatusInternalServerError)
  }
}

// serverErrorResponse(): A helper method used when application encounteres an error in runtime sending generic error message to client
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error){
  app.logError(r, err)
  message := "server encountered a problem and could not process your request"
  app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse(): A helper method to send 404 not found status code
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
  message := "the requested resource could not be found"
  app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse(): A helper method to send 404 not found status code
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
  message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
  app.errorResponse(w, r, http.StatusNotFound, message)
}
