package models

import "time"

type EnergyCertificate struct {
	TokenRef              string    `json:"tokenRef"`
	OwnerID               int       `json:"ownerId"`
	ProducerID            int       `json:"producerId"`
	EmissionDate          time.Time `json:"emissionDate"`
	UsableMonth           int       `json:"usableMonth"`
	UsableYear            int       `json:"usableYear"`
	RegulatoryAuthorityID int       `json:"regulatoryAuthorityID"`
}
