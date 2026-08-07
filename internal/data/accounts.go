package data

import (
	"database/sql"
	"errors"
	"strings"
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
	err := a.DB.QueryRow(query, account.DocumentNumber).Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		// document_number is UNIQUE; report the clash as a domain error so the
		// handler can return 422 rather than a generic 500.
		if strings.Contains(err.Error(), "accounts_document_number_key") {
			return ErrDuplicateDocument
		}
		return err
	}
	return nil
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
