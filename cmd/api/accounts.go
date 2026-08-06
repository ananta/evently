package main

import (
	"fmt"
	"net/http"

	"github.com/ananta/evently/internal/data"
)

func (app *application) createAccount(w http.ResponseWriter, r *http.Request){
  var input struct {
    Document_Number string `json:"document_number"`
  }

  err := app.readJSON(w, r, &input)
  if err != nil {
    app.badRequestResponse(w, r, err)
    return
  }

  fmt.Fprintf(w, "%+v\n", input)
}

