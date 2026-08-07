package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound    = errors.New("record not found")
	ErrDuplicateDocument = errors.New("duplicate document number")
)

type Models struct {
	Accounts       AccountModel
	Transactions   TransactionModel
	OperationTypes OperationTypeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Accounts: AccountModel{
			DB: db,
		},
		Transactions: TransactionModel{
			DB: db,
		},
		OperationTypes: OperationTypeModel{
			DB: db,
		},
	}
}
