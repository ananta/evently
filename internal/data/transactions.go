package data

type Transactions struct {
  Transaction_ID int `json:"transaction_id"`
  Account_ID int `json:"account_id"`
  OperationType_ID int  `json:"operation_type_id"`
  Amount Amount  `json:"amount,omitzero"`
  EventDate string  `json:"event_date"`
}
