package models

type Admin struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	BlockchainUser int    `json:"blockchainUser"`
	Address        string `json:"address"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
}
