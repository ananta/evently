package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
  "strings"
  "github.com/ananta/evently/internal/data"

	"github.com/julienschmidt/httprouter"

)

type envelope map[string]any
// write json, we could use json.NewEncoder(), but this allows us to set conditional http response headers
// also minimal perf diff compared to json.NewEncoder()
func (app *application) writeJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
  js, err :=  json.MarshalIndent(data, "", "\t")
  if err != nil {
    return err
  }
  js = append(js, '\n')
  
  for key,values := range headers {
    for _, value := range values {
      w.Header().Add(key, value)
    }
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write(js)
  return nil
}

// decode the request body and map it to dest
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
  // limit the size of the request body to 1MB(1,048,576 bytes)
  r.Body = http.MaxBytesReader(w, r.Body, 1_048_576)
  // disallow unknown fields before decoding
  dec := json.NewDecoder(r.Body)
  dec.DisallowUnknownFields()

  err := dec.Decode(dst)
  if err != nil {
    var syntaxError *json.SyntaxError
    var unmarshalTypeError *json.UnmarshalTypeError
    var invalidUnmarshalError *json.InvalidUnmarshalError
    var maxBytesError *http.MaxBytesError

    switch {
    case errors.As(err, &syntaxError):
      return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
    case errors.Is(err, io.ErrUnexpectedEOF):
      return errors.New("body contains badly-formed JSON")
    case errors.As(err, &unmarshalTypeError):
      if unmarshalTypeError.Field != "" {
        return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
      }
      return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
    case errors.Is(err, io.EOF):
      return errors.New("body must not be empty")
    case strings.HasPrefix(err.Error(), "json: unknown field "):
      fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
      return fmt.Errorf("body contains unknown key %s", fieldName)
    case errors.As(err, &maxBytesError):
      return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

    case errors.As(err, &invalidUnmarshalError):
      panic(err)
    default:
    return err
    }
  }


  // call the decode again, if its a success, it means we have multiple json objects, reject
  err = dec.Decode(&struct{}{})
  if !errors.Is(err, io.EOF){
    return errors.New("body must only contain a single JSON value")
  }
  return nil
}


func (app *application) seedOperationTypes(){
  operation_types := []string{
    "Normal Purchase",
    "Purchase with installments",
    "Withdrawal",
    "Credit Voucher",
  }

  hasError := false
  for _, operation := range operation_types {
    operation_type := data.OperationType {
      Description: operation,
    }
    // TOOD: handle error
    err := app.models.OperationTypes.Insert(&operation_type)
    if err != nil {
      hasError = true
    }

  }
  if hasError == true {
    app.logger.Info("Failed Seeding data")
    return
  }
  app.logger.Info("Success Seeding data")

}

