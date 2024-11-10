package models

import "time"

type EnergyCertificate struct {
	TokenRef              string    `json:"tokenRef"`
	OwnerID               string    `json:"ownerId"`
	ProducerID            string    `json:"producerId"`
	EmissionDate          time.Time `json:"emissionDate"`
	UsableMonth           int       `json:"usableMonth"`
	UsableYear            int       `json:"usableYear"`
	RegulatoryAuthorityID string    `json:"regulatoryAuthorityID"`
	AvailableToSell       bool      `json:"availableToSell"`
	EnergyType            int       `json:"energyType"`
}
