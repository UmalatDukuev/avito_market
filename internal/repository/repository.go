package repository

import (
	"market/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Information
}

type Authorization interface {
	CreateUser(models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Information interface {
	GetInfo(int) (models.Info, error)
}

type Transaction interface {
	SendCoin(int) (models.Info, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Information:   NewInfoPostgres(db),
	}
}
