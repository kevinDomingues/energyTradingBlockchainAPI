package models

type User struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	BlockchainUser string `json:"blockchainUser"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
}
