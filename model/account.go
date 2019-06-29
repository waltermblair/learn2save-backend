package model

type AccountResponseBody struct {
	ID 				  int       `json:"id"`
	Balance           float64   `json:"balance"`
	LastTransactionAt string `json:"last_transaction_at"`
	LastUpdatedAt     string `json:"last_updated_at"`
	CreatedAt         string `json:"created_at"`
}