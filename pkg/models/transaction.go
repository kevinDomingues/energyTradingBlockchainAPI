package models

import "time"

type Transaction struct {
	TokenRef        string    `json:"tokenRef"`
	BlockchainUser  string    `json:"blockchainUser"`
	TransactionDate time.Time `json:"transactionDate"`
	Price           float64   `json:"price"`
}
