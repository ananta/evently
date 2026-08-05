package data

import (
  "fmt"
  "strconv"
)

type Amount int

func (a Amount) MarshalJSON() ([]byte, error){
  jsonValue := fmt.Sprintf("$%d", a)
  quotedJSONValue := strconv.Quote(jsonValue)
  return []byte(quotedJSONValue), nil
}
