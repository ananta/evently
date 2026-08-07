package main

import (
	"errors"
	"net/http"

	"github.com/ananta/evently/internal/data"
)

func (app *application) getAccount(w http.ResponseWriter, r *http.Request){
  id, err := app.readIDParam(r)
  if err != nil {
    app.notFoundResponse(w, r)
  }
  account, err := app.models.Accounts.Get(int64(id))
  if err != nil {
    switch{
      case errors.Is(err, data.ErrRecordNotFound):
        app.notFoundResponse(w, r)
      default:
        app.serverErrorResponse(w, r, err)
    }
    return
  }
  err = app.writeJson(w, http.StatusOK, envelope{"account": account}, nil)
  if err != nil {
    app.serverErrorResponse(w, r, err)
  }
}

func (app *application) createAccount(w http.ResponseWriter, r *http.Request){
  var input struct {
    Document_Number string `json:"document_number"`
  }

  err := app.readJSON(w, r, &input)
  if err != nil {
    app.badRequestResponse(w, r, err)
    return
  }

  account := &data.Account{
    Document_Number: input.Document_Number,
  }
  err = app.models.Accounts.Insert(account)
  if err != nil {
    app.serverErrorResponse(w, r, err)
    return
  }

  err = app.writeJson(w, http.StatusCreated, envelope{"account": account}, nil)
  if err != nil {
    app.serverErrorResponse(w, r, err)
  }
}

