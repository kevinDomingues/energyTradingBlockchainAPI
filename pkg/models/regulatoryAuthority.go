package models

type RegulatoryAuthority struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	Address string `json:"address"`
	APIURL  string `json:"apiURL"`
}
