package models

type EnergyProducer struct {
	ID                 int     `json:"id"`
	ProductionCapacity float64 `json:"productionCapacity"`
}
