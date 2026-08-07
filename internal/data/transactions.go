package data

import (
	"database/sql"
	"errors"

	"github.com/shopspring/decimal"
)

// Unmarshal JSON without the quotes as decimal package's default is quoted strings
func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

type Transaction struct {
	TransactionID   int64           `json:"transaction_id"`
	AccountID       int64           `json:"account_id"`
	OperationTypeID int64           `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
	EventDate       string          `json:"event_date"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (t TransactionModel) Insert(transaction *Transaction) error {
	query := `
  INSERT INTO transactions (account_id, operation_type_id, amount)
  VALUES ($1, $2, $3)
  RETURNING transaction_id, account_id, operation_type_id, amount
  `
	err := t.DB.QueryRow(query, transaction.AccountID, transaction.OperationTypeID, transaction.Amount).Scan(&transaction.TransactionID, &transaction.AccountID, &transaction.OperationTypeID, &transaction.Amount)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}
	return nil
}
