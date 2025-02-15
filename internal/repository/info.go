package repository

import (
	"market/models"

	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) GetInfo(ID int) (models.Info, error) {
	var info models.Info
	return info, nil
}
