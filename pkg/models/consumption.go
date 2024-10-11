package models

type Consumption struct {
	UserID           int     `json:"userId"`
	ConsumptionYear  int     `json:"consumptionYear"`
	ConsumptionMonth int     `json:"consumptionMonth"`
	EnergyConsumed   float64 `json:"energyConsumed"`
}
