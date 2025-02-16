package repository

import (
	"market/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Information
	Transaction
}

type Authorization interface {
	CreateUser(models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Information interface {
	GetInfo(int) (models.Info, error)
}

type Transaction interface {
	SendCoins(fromUserID, toUserID, amount int) error
	GetCoinsByID(userID int) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Information:   NewInfoPostgres(db),
		Transaction:   NewTransactionPostgres(db),
	}
}
