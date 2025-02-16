package models

type Transaction struct {
	ID        int    `json:"id" db:"id"`
	FromUser  string `json:"from_user" db:"from_user"`
	ToUser    string `json:"to_user" db:"to_user"`
	Amount    int    `json:"amount" db:"amount"`
	Item      int    `json:"item" db:"item"`
	Timestamp string `json:"timestamp" db:"timestamp"`
}
