package data

import (
	"database/sql"
	"errors"
)

type Account struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type AccountModel struct {
	DB *sql.DB
}

func (a AccountModel) Insert(account *Account) error {
	query := `
  INSERT INTO accounts (document_number)
  VALUES ($1)
  RETURNING account_id, document_number`
	return a.DB.QueryRow(query, account.DocumentNumber).Scan(&account.AccountID, &account.DocumentNumber)
}

func (a AccountModel) Get(id int64) (*Account, error) {
	query := `
  SELECT account_id, document_number FROM accounts WHERE account_id = $1
  `
	var account Account
	err := a.DB.QueryRow(query, id).Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &account, nil
}
