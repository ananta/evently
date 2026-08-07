package main

import (
	"net/http"

	"github.com/ananta/evently/internal/data"
)

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID       int64   `json:"account_id"`
		OperationTypeID int64   `json:"operation_type_id"`
		Amount          float32 `json:"amount"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// TODO:
	// validate if operation type exists
	// validate if proper sign is used in operation type
	// validate document number

	transaction := &data.Transaction{
		AccountID:       input.AccountID,
		Amount:          input.Amount,
		OperationTypeID: input.OperationTypeID,
	}
	err = app.models.Transactions.Insert(transaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"transaction": transaction}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
