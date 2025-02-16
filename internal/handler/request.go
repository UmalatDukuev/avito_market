package handler

type SendCoinRequest struct {
	ToUser int `json:"toUser" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

type ReceivedCoinRequest struct {
	FromUser int `json:"toUser" binding:"required"`
	Amount   int `json:"amount" binding:"required"`
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
