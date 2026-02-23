package models

type TransactionResponse struct {
	ID     uint    `json:"id"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type AccountWithTransactionsResponse struct {
	AccountID    uint                  `json:"account_id"`
	Transactions []TransactionResponse `json:"transactions"`
}
