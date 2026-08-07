package data

import "database/sql"

type OperationType struct {
	OperationTypeID int    `json:"operation_type_id"`
	Description     string `json:"description"`
}

type OperationTypeModel struct {
	DB *sql.DB
}

