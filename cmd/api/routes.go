package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler{
  router := httprouter.New()

  // use our custom error handlers for error consistency
  router.NotFound = http.HandlerFunc(app.notFoundResponse)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

  // retrieve the account information
  router.HandlerFunc(http.MethodGet, "/accounts/:id", app.getAccount)
  // create an account
  router.HandlerFunc(http.MethodPost, "/accounts", app.createAccount)
  // create a transaction

  return app.recoverPanic(app.logRequest(commonHeaders(router))) 
}
