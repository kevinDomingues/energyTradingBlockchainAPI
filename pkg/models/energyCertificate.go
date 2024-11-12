package models

type EnergyCertificate struct {
	energyCertificateId   string `json:"energyCertificateId"`
	OwnerID               string `json:"ownerId"`
	ProducerID            string `json:"producerId"`
	UsableMonth           int    `json:"usableMonth"`
	UsableYear            int    `json:"usableYear"`
	RegulatoryAuthorityID string `json:"regulatoryAuthorityID"`
	AvailableToSell       bool   `json:"availableToSell"`
	EnergyType            int    `json:"energyType"`
}
