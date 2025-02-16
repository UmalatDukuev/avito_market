package models

type Item struct {
	ID          int    `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
	Quantity    int    `json:"quantity" db:"quantity"`
	Price       int    `json:"price,omitempty" db:"price"`
}
