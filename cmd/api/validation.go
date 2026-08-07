package main

import (
	"strings"
	"unicode"

	"github.com/ananta/evently/internal/data"
	"github.com/shopspring/decimal"
)

// validDocumentNumber returns the reason a document number is unacceptable, or
// an empty string if it is fine. The column allows 50 characters.
func validDocumentNumber(doc string) string {
	switch {
	case doc == "":
		return "must be provided"
	case len(doc) > 50:
		return "must not be more than 50 characters long"
	case strings.TrimFunc(doc, unicode.IsDigit) != "":
		return "must contain digits only"
	}
	return ""
}

// validAmount returns the reason an amount is unacceptable for the given
// operation type, or an empty string if it is fine. Purchases and withdrawals
// are recorded with negative amounts, credit vouchers with positive ones.
func validAmount(operationTypeID int64, amount decimal.Decimal) string {
	switch {
	case amount.IsZero():
		return "must not be zero"
	case !data.IsOperationType(operationTypeID):
		// The sign rule is unknowable, and the operation type is reported
		// separately, so say nothing about the amount here.
		return ""
	case data.IsCredit(operationTypeID) && amount.IsNegative():
		return "must be positive for a credit voucher"
	case !data.IsCredit(operationTypeID) && amount.IsPositive():
		return "must be negative for a purchase or withdrawal"
	}
	return ""
}
