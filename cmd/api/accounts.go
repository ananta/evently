package main

import (
	"errors"
	"net/http"

	"github.com/ananta/evently/internal/data"
)

func (app *application) getAccount(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	account, err := app.accounts.Get(r.Context(), int64(id))
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, account, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createAccount(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DocumentNumber string `json:"document_number"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if msg := validDocumentNumber(input.DocumentNumber); msg != "" {
		app.failedValidationResponse(w, r, map[string]string{"document_number": msg})
		return
	}

	account := &data.Account{
		DocumentNumber: input.DocumentNumber,
	}
	err = app.accounts.Insert(r.Context(), account)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateDocument):
			app.failedValidationResponse(w, r, map[string]string{
				"document_number": "an account with this document number already exists",
			})
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, account, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
