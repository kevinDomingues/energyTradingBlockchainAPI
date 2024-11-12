package models

type Consumption struct {
	ID               uint    `gorm:"primaryKey"`
	UserID           string  `json:"userId"`
	ConsumptionYear  int     `json:"consumptionYear"`
	ConsumptionMonth int     `json:"consumptionMonth"`
	EnergyTypeId     int     `json:"energyTypeId"`
	EnergyConsumed   float64 `json:"energyConsumed"`
}
