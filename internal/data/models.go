package data

import (
	"database/sql"
)

type Models struct {
  Accounts AccountModel
  Transactions TransactionModel
  OperationTypes OperationTypeModel
}

func NewModels(db *sql.DB) Models {
  return Models {
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
