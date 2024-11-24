package models

type TransferCertificate struct {
	EnergyCertificateID string `json:"energyCertificateId"`
	Quantity            int    `json:"quantity"`
	Availability        int    `json:"availability"`
}
