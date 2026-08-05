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
