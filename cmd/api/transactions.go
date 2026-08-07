package main

import (
	"net/http"

	"github.com/ananta/evently/internal/data"
)

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Account_ID        int64   `json:"account_id"`
		Operation_Type_ID int64   `json:"operation_type_id"`
		Amount            float32 `json:"amount"`
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
		Account_ID:       input.Account_ID,
		Amount:           input.Amount,
		OperationType_ID: input.Operation_Type_ID,
	}
	err = app.models.Transactions.Insert(transaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJson(w, http.StatusCreated, envelope{"transaction": transaction}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
