package data

import (
	"context"
	"database/sql"
	"time"

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
	EventDate       time.Time       `json:"event_date"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (t TransactionModel) Insert(ctx context.Context, transaction *Transaction) error {
	query := `
  INSERT INTO transactions (account_id, operation_type_id, amount)
  VALUES ($1, $2, $3)
  RETURNING transaction_id, account_id, operation_type_id, amount, event_date
  `
	return t.DB.QueryRowContext(ctx, query, transaction.AccountID, transaction.OperationTypeID, transaction.Amount).Scan(
		&transaction.TransactionID,
		&transaction.AccountID,
		&transaction.OperationTypeID,
		&transaction.Amount,
		&transaction.EventDate,
	)
}
