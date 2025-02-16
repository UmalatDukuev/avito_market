package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) SendCoins(fromUserID, toUserID, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", amount, fromUserID)
	if err != nil {
		return fmt.Errorf("failed to withdraw coins from user %d: %w", fromUserID, err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins + $1 WHERE id = $2", amount, toUserID)
	if err != nil {
		return fmt.Errorf("failed to deposit coins to user %d: %w", toUserID, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *TransactionPostgres) GetCoinsByID(userID int) (int, error) {
	var coins int
	query := "SELECT coins FROM users WHERE id = $1"
	row := r.db.QueryRow(query, userID)
	if err := row.Scan(&coins); err != nil {
		return 0, err
	}
	return coins, nil
}
