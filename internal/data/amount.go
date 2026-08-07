package data

import (
	"fmt"
	"strconv"
)

type Amount float32

func (a Amount) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("$%d", a)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
