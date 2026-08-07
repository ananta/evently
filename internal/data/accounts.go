package data

import (
	"database/sql"
	"errors"
	"fmt"
)

type Account struct {
  Account_ID int64 `json:"acount_id"`
  Document_Number string  `json:"document_number"`
}

type AccountModel struct {
  DB *sql.DB
}

func (a AccountModel) Insert(account *Account) error {
  query := `
  INSERT INTO accounts (document_number)
  VALUES ($1)
  RETURNING account_id, document_number`
  return a.DB.QueryRow(query, account.Document_Number).Scan(&account.Account_ID, &account.Document_Number)
}

func (a AccountModel) Get(account_id int64) (*Account, error) {
  query := `
  SELECT account_id, document_number FROM accounts WHERE account_id = $1
  `
  fmt.Println(account_id)
  var account Account
  err := a.DB.QueryRow(query, account_id).Scan(&account.Account_ID, &account.Document_Number)
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

