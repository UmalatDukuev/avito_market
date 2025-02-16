package models

type Info struct {
	Coins       int `json:"coins"`
	Inventory   []Item
	CoinHistory struct {
		Received []ReceivedCoinRequest
		Sent     []SendCoinRequest
	}
}

type SendCoinRequest struct {
	ToUser int `json:"toUser" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type ReceivedCoinRequest struct {
	FromUser int `json:"fromUser" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
}
