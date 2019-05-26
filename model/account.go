package model

import "time"

type AccountResponseBody struct {
	ID 				  int       `json:"id"`
	Balance           float64   `json:"balance"`
	LastTransactionAt time.Time `json:"last_transaction_at"`
	LastUpdatedAt     time.Time `json:"last_updated_at"`
	CreatedAt         time.Time `json:"created_at"`
}