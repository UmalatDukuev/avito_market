package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) BuyItem(userID int, itemName string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var itemID, price int
	err = tx.QueryRow("SELECT id, price FROM merch WHERE name = $1", itemName).Scan(&itemID, &price)
	if err != nil {
		return fmt.Errorf("item not found: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", price, userID)
	if err != nil {
		return fmt.Errorf("failed to deduct balance: %w", err)
	}

	var existingQuantity int
	err = tx.QueryRow("SELECT quantity FROM inventory WHERE user_id = $1 AND merch_id = $2", userID, itemID).Scan(&existingQuantity)

	if err != nil {
		_, err = tx.Exec("INSERT INTO inventory (user_id, merch_id, quantity) VALUES ($1, $2, 1)", userID, itemID)
		if err != nil {
			return fmt.Errorf("failed to insert item into inventory: %w", err)
		}
	} else {
		_, err = tx.Exec("UPDATE inventory SET quantity = quantity + 1 WHERE user_id = $1 AND merch_id = $2", userID, itemID)
		if err != nil {
			return fmt.Errorf("failed to update item quantity: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *ItemPostgres) GetCoinsByID(userID int) (int, error) {
	var coins int
	query := "SELECT coins FROM users WHERE id = $1"
	row := r.db.QueryRow(query, userID)
	if err := row.Scan(&coins); err != nil {
		return 0, err
	}
	return coins, nil
}

func (r *ItemPostgres) GetItemPrice(itemName string) (int, error) {
	var price int
	query := "SELECT price FROM merch WHERE name = $1"
	err := r.db.QueryRow(query, itemName).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("failed to get price for item %s: %w", itemName, err)
	}
	return price, nil
}
