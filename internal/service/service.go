package service

import (
	"market/internal/repository"
	"market/models"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Information interface {
	GetInfo(int) (models.Info, error)
}

type Transaction interface {
	SendCoins(fromUserID, toUserID, amount int) error
}

type Item interface {
	BuyItem(userID int, itemName string) error
}

type Service struct {
	Authorization
	Information
	Transaction
	Item
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Information:   NewInfoService(repo.Information),
		Transaction:   NewTransactionService(repo.Transaction),
		Item:          NewItemService(repo.Item),
	}
}
