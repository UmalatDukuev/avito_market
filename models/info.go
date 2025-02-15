package models

type Info struct {
	Coins       int    `json:"coins"`
	Inventory   []Item `json:"inventory"`
	CoinHistory struct {
		Received []Transaction `json:"received"`
		Sent     []Transaction `json:"sent"`
	} `json:"coinHistory"`
}
