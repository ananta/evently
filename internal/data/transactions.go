package data


type Transaction struct {
  Transaction_ID int64 `json:"transaction_id"`
  Account_ID int64 `json:"account_id"`
  OperationType_ID int64  `json:"operation_type_id"`
  Amount float32  `json:"amount,omitzero"`
  EventDate string  `json:"event_date"`
}
