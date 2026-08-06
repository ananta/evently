package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request){
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
