package models

type BlockchainMethod struct {
	Method string   `json:"method"`
	Args   []string `json:"args"`
}
