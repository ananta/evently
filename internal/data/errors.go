package data

import "errors"

var (
	ErrRecordNotFound    = errors.New("record not found")
	ErrDuplicateDocument = errors.New("duplicate document number")
)
