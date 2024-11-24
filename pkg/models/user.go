package models

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	BlockchainUser string `json:"blockchainUser"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	UserType       int    `json:"userType"`
}

type UserValidation struct {
	UserId string `json:"userId"`
}

type UserValidationResponse struct {
	RegulatoryAuthority string `json:"regulatoryId"`
	Accepted            bool   `json:"accepted"`
}
