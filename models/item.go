package models

type Item struct {
	ID       int    `json:"id" db:"id"`
	Type     string `json:"type" db:"type"`
	Quantity int    `json:"quantity" db:"quantity"`
	OwnerID  int    `json:"owner_id" db:"owner_id"`
}
