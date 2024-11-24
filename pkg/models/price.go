package models

type Price struct {
	Price float64 `json:"price"`
}

type PriceRequest struct {
	Quantity     int `json:"quantity"`
	Availability int `json:"availability"`
}
