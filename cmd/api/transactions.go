package main

import (
	"errors"
	"net/http"

	"github.com/ananta/evently/internal/data"
	"github.com/shopspring/decimal"
)

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID       int64           `json:"account_id"`
		OperationTypeID int64           `json:"operation_type_id"`
		Amount          decimal.Decimal `json:"amount"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := make(map[string]string)

	if input.AccountID < 1 {
		v["account_id"] = "must be a positive integer"
	}
	if !data.IsOperationType(input.OperationTypeID) {
		v["operation_type_id"] = "must be a known operation type"
	}
	if msg := validAmount(input.OperationTypeID, input.Amount); msg != "" {
		v["amount"] = msg
	}

	if len(v) > 0 {
		app.failedValidationResponse(w, r, v)
		return
	}

	// The transactions.account_id foreign key would reject this anyway, but
	// checking here turns a constraint violation into a useful message.
	_, err = app.accounts.Get(r.Context(), input.AccountID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.failedValidationResponse(w, r, map[string]string{
				"account_id": "account does not exist",
			})
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	transaction := &data.Transaction{
		AccountID:       input.AccountID,
		Amount:          input.Amount,
		OperationTypeID: input.OperationTypeID,
	}
	err = app.transactions.Insert(r.Context(), transaction)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, transaction, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
