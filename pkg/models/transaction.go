package models

import "time"

type Transaction struct {
	TransactionID       string    `json:"transactionId"`
	CertificateTokenRef string    `json:"tokenRef"`
	FromUserID          string    `json:"fromUserId"`
	ToUserID            string    `json:"toUserId"`
	TransactionDate     time.Time `json:"transactionDate"`
	Price               float64   `json:"price"`
}
