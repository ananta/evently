package data

import (
	"database/sql"
	"errors"
)

type Transaction struct {
	Transaction_ID   int64   `json:"transaction_id"`
	Account_ID       int64   `json:"account_id"`
	OperationType_ID int64   `json:"operation_type_id"`
	Amount           float32 `json:"amount,omitzero"`
	EventDate        string  `json:"event_date"`
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
	err := t.DB.QueryRow(query, transaction.Account_ID, transaction.OperationType_ID, transaction.Amount).Scan(&transaction.Transaction_ID, &transaction.Account_ID, &transaction.OperationType_ID, &transaction.Amount)
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
